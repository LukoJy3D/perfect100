package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Struct to represent the JSON data
	type GamesData struct {
		Games []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"games"`
	}

	// Read the file into a byte slice
	jsonData, err := os.ReadFile("games.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a struct
	var data GamesData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Loop through the games and make the HTTP request for each one
	for _, game := range data.Games {
		url := fmt.Sprintf("https://steamcommunity.com/stats/%s/achievements", game.ID)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Process the response
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Get the text from the title tag
		title := doc.Find("title").Text()

		// Extract the game title from the title text
		gameTitle := strings.Split(title, "::")[1]
		gameTitle = strings.TrimSpace(gameTitle)
		gameTitleRaw := gameTitle
		gameTitle = ReplaceForbiddenCharacters(gameTitle)

		// Create a folder with the game title
		if _, err := os.Stat(gameTitle + "/guides"); err == nil {
			// Folder already exists, no need to create
		} else {
			if os.IsNotExist(err) {
				err := os.Mkdir(gameTitle+"/guides", os.ModePerm)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}
		}

		// Create a achievement list file
		guideList, err := os.Create(gameTitle + "/" + gameTitle + ".md")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		defer guideList.Close()

		// Front matter information
		frontMatter := fmt.Sprintf("---\nlayout: default\ntitle: %s\nhas_children: true\n---\n\n", gameTitleRaw)

		if _, err = guideList.WriteString(frontMatter); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		doc.Find(".achieveRow").Each(func(i int, selection *goquery.Selection) {
			imageSrcBase := path.Base(selection.Find(".achieveImgHolder img").AttrOr("src", ""))
			imageSrc := "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/" + game.ID + "/" + imageSrcBase
			achievePercent := selection.Find(".achievePercent").Text()
			achieveTextH3 := selection.Find(".achieveTxt h3").Text()
			achieveTextH5 := selection.Find(".achieveTxt h5").Text()
			modifiedAchiName := ReplaceForbiddenCharacters(achieveTextH3)
			guidePath := gameTitle + "/guides/" + modifiedAchiName + ".md"

			markdown := fmt.Sprintf("# %s ([guide](%s)) <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\n\nOwned by **%s** of players\n\n_%s_\n\n---\n\n",
				achieveTextH3, guidePath, imageSrc, achievePercent, achieveTextH5)

			fmt.Printf("Populating achievement list - Game: \"%s\", Achievement: \"%s\"...\n", gameTitle, achieveTextH3)

			//markdown content with front matter
			if _, err = guideList.WriteString(markdown); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			// Create a file with the achievement title as the name
			if _, err := os.Stat(gameTitle + "/guides/" + modifiedAchiName + ".md"); err != nil {

				f2, err := os.OpenFile(gameTitle+"/guides/"+modifiedAchiName+".md", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				//run ai setup only before guide generation
				filename := "gpt4all-lora-quantized.bin"
				url := "https://the-eye.eu/public/AI/models/nomic-ai/gpt4all/gpt4all-lora-quantized.bin"

				if _, err := os.Stat(filename); os.IsNotExist(err) {
					fmt.Println("Downloading gpt4all binary...")
					cmd := exec.Command("curl", "-L", "-o", filename, url)
					if err := cmd.Run(); err != nil {
						fmt.Println("Error downloading gpt4all binary file:", err)
						return
					}
				} else {
					fmt.Println("gpt4all binary file already exists")
				}

				fmt.Printf("Generating guide using AI - Game: \"%s\", Achievement: \"%s\"...\r\n", gameTitle, achieveTextH3)

				// Command to run the gpt4all-lora-quantized binary
				var cmd *exec.Cmd

				if runtime.GOOS == "windows" {
					cmd = exec.Command("./gpt4all/chat/gpt4all-lora-quantized-win64.exe", "-p", fmt.Sprintf(
						"Write a step-by-step guide that would help unlock '%s' achievement in a game called '%s'. It requires to %s.",
						achieveTextH3, gameTitle, achieveTextH5))
				} else if runtime.GOOS == "linux" {
					cmd = exec.Command("./gpt4all/chat/gpt4all-lora-quantized-linux", "-p", fmt.Sprintf(
						"Write a step-by-step guide that would help unlock '%s' achievement in a game called '%s'. It requires to %s.",
						achieveTextH3, gameTitle, achieveTextH5))
				} else {
					fmt.Println("Unsupported operating system, add additional if statement for your OS")
					return
				}

				// Capture the output of the command
				out, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}

				// Convert the output to a string
				guide := string(out)

				markdown2 := fmt.Sprintf("---\nlayout: default\ntitle: %s\nparent: %s\n---\n\n# %s <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\r\n\r\n_%s_\r\n\r\n***\r\n\r\n%s\r\n\r\n"+
					":heavy_check_mark: Guide was generated by AI\r\n\r\n:x: Guide edited by a human", achieveTextH3, gameTitleRaw, achieveTextH3, imageSrc, achieveTextH5, guide)
				f2.WriteString(markdown2)

				defer f2.Close()

			} else {
				//if file exists update achieveTextH3, imageSrc, achieveTextH5, but do not overwrite answer section under ***
				fileread, err := os.ReadFile(gameTitle + "/guides/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				var part2 string

				parts := strings.Split(string(fileread), "***")
				part2 = parts[1]

				f2, err := os.Create(gameTitle + "/guides/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				fmt.Printf("Updating guide stats - Game: \"%s\", Achievement: \"%s\"...\r\n", gameTitle, achieveTextH3)

				// update the variables
				markdown3 := fmt.Sprintf("---\nlayout: default\ntitle: %s\nparent: %s\n---\n\n# %s (%s) <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\r\n\r\n_%s_\r\n\r\n***%s",
					achieveTextH3, gameTitleRaw, achieveTextH3, achievePercent, imageSrc, achieveTextH5, part2)

				// write the updated content to the file
				if _, err = f2.WriteString(markdown3); err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				defer f2.Close()
			}
		})
	}
}

func ReplaceForbiddenCharacters(s string) string {
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, ":", "_", -1)
	s = strings.Replace(s, "?", "_", -1)
	s = strings.Replace(s, "*", "_", -1)
	s = strings.Replace(s, "!", "_", -1)
	s = strings.Replace(s, "[", "_", -1)
	s = strings.Replace(s, "]", "_", -1)
	s = strings.Replace(s, "\"", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "\\", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, ",", "_", -1)
	s = strings.Replace(s, "'", "_", -1)
	s = strings.Replace(s, "@", "_", -1)
	s = strings.Replace(s, "#", "_", -1)
	s = strings.Replace(s, "$", "_", -1)
	s = strings.Replace(s, "%", "_", -1)
	s = strings.Replace(s, "^", "_", -1)
	s = strings.Replace(s, "&", "_", -1)
	s = strings.Replace(s, "(", "_", -1)
	s = strings.Replace(s, ")", "_", -1)
	s = strings.Replace(s, "+", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, "=", "_", -1)
	s = strings.Replace(s, "<", "_", -1)
	s = strings.Replace(s, ">", "_", -1)
	return s
}
