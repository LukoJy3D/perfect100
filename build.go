package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/solywsh/chatgpt"
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
		gameTitle = ReplaceForbiddenCharacters(gameTitle)

		// Create a achievement list file with the game title as the name
		achievementList, err := os.Create("achievements/" + gameTitle + ".md")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		defer achievementList.Close()

		// Create a folder with the game title
		if _, err := os.Stat("guides/" + gameTitle); err == nil {
			// Folder already exists, no need to create
		} else {
			if os.IsNotExist(err) {
				err := os.Mkdir("guides/"+gameTitle, os.ModePerm)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}
		}

		doc.Find(".achieveRow").Each(func(i int, selection *goquery.Selection) {
			imageSrc := selection.Find(".achieveImgHolder img").AttrOr("src", "")
			achievePercent := selection.Find(".achievePercent").Text()
			achieveTextH3 := selection.Find(".achieveTxt h3").Text()
			achieveTextH5 := selection.Find(".achieveTxt h5").Text()
			modifiedAchiName := ReplaceForbiddenCharacters(achieveTextH3)

			guidePath := "/guides/" + gameTitle + "/" + modifiedAchiName + ".md"

			markdown := fmt.Sprintf("# %s ([guide](%s)) <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\n\nOwned by **%s** of players\n\n_%s_\n\n---\n\n",
				achieveTextH3, guidePath, imageSrc, achievePercent, achieveTextH5)

			if _, err = achievementList.WriteString(markdown); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			api_key := os.Getenv("api_key")

			// Create a file with the achievement title as the name
			if _, err := os.Stat("guides/" + gameTitle + "/" + modifiedAchiName + ".md"); err != nil {
				f2, err := os.OpenFile("guides/"+gameTitle+"/"+modifiedAchiName+".md", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				// Create a new instance of the chatgpt client
				chat := chatgpt.New(api_key, "", 30*time.Second)
				defer chat.Close()

				question := achieveTextH3 + " is a steam achievement in " + gameTitle +
					". Can you please provide a guide on how to get it? Also add some related emojis readable in markdown!"

				answer, err := chat.Chat(question)
				if err != nil {
					fmt.Println(err)
				}

				markdown2 := fmt.Sprintf("# %s <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\n\n_%s_\n\n---\n\n%s",
					achieveTextH3, imageSrc, achieveTextH5, answer)
				f2.WriteString(markdown2)
				defer f2.Close()
			} else {
				//if file exists update achieveTextH3, imageSrc, achieveTextH5, but do not overwrite answer section under ---
				fileread, err := os.ReadFile("guides/" + gameTitle + "/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				var part2 string

				parts := strings.Split(string(fileread), "---")
				part2 = parts[1]

				f2, err := os.Create("guides/" + gameTitle + "/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				// update the variables
				markdown3 := fmt.Sprintf("# %s (%s) <img style=\"float: right;\" src=\"%s\" width=\"96\" height=\"96\">\n\n_%s_\n\n---%s",
					achieveTextH3, achievePercent, imageSrc, achieveTextH5, part2)

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
