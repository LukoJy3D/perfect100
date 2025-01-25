package utils

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

func Achievements(gameName string) {

	type GamesData struct {
		Games []struct {
			ID   string `yaml:"id"`
			Name string `yaml:"name"`
		} `yaml:"games"`
	}

	yamlData, err := os.ReadFile("games.yml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data GamesData
	if err := yaml.Unmarshal(yamlData, &data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Loop through the games and make the HTTP request for each one
	for _, game := range data.Games {

		if gameName != "all" && game.Name != gameName {
			continue // Skip if gameName is specified and doesn't match the current game
		}

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

		// Get the total number of achievements
		totalAchievements := doc.Find(".achieveRow").Length()

		// Extract the game title from the title text
		gameTitle := strings.Split(title, "::")[1]
		gameTitle = strings.TrimSpace(gameTitle)
		gameTitleRaw := gameTitle
		gameTitle = ReplaceForbiddenCharacters(gameTitle)

		// Create a folder with the game title
		if _, err := os.Stat("guides/" + gameTitle + "/achievements"); err == nil {
			// Folder already exists, no need to create
		} else {
			if os.IsNotExist(err) {
				err := os.Mkdir("guides/"+gameTitle+"/achievements", os.ModePerm)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}
		}

		// Create a achievement list file
		guideList, err := os.Create("guides/" + gameTitle + "/" + gameTitle + ".md")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		defer guideList.Close()

		// Front matter information
		frontMatter := fmt.Sprintf("---\r\nlayout: default\r\ntitle: %s\r\nhas_children: true\r\n---\r\n\r\n", gameTitleRaw)

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
			guidePath := "achievements/" + modifiedAchiName + ".md"

			var markdown string
			if i == totalAchievements-1 {
				markdown = fmt.Sprintf("## [%s](%s) <img align=\"right\" src=\"%s\" alt=\"'%s' achievement icon\" width=\"96\" height=\"96\">\r\n\r\n"+
					"Owned by **%s** of players\r\n\r\nObjective: _%s_\r\n", achieveTextH3, guidePath, imageSrc, achieveTextH3, achievePercent, achieveTextH5)
			} else {
				markdown = fmt.Sprintf("## [%s](%s) <img align=\"right\" src=\"%s\" alt=\"'%s' achievement icon\" width=\"96\" height=\"96\">\r\n\r\n"+
					"Owned by **%s** of players\r\n\r\nObjective: _%s_\r\n\r\n---\r\n\r\n", achieveTextH3, guidePath, imageSrc, achieveTextH3, achievePercent, achieveTextH5)
			}

			fmt.Printf("Populating achievement list - Game: \"%s\", Achievement: \"%s\"...\n", gameTitleRaw, achieveTextH3)

			//markdown content with front matter
			if _, err = guideList.WriteString(markdown); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			// Create a file with the achievement title as the name
			if _, err := os.Stat("guides/" + gameTitle + "/achievements/" + modifiedAchiName + ".md"); err != nil {

				f2, err := os.OpenFile("guides/"+gameTitle+"/achievements/"+modifiedAchiName+".md", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				defer f2.Close()

			} else {
				//if file exists update achieveTextH3, imageSrc, achieveTextH5, but do not overwrite answer section under ---
				fileread, err := os.ReadFile("guides/" + gameTitle + "/achievements/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				var achievements_guide string

				parts := strings.Split(string(fileread), "---")
				achievements_guide = parts[3]

				f2, err := os.Create("guides/" + gameTitle + "/achievements/" + modifiedAchiName + ".md")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				fmt.Printf("Updating guide stats - Game: \"%s\", Achievement: \"%s\"...\r\n", gameTitleRaw, achieveTextH3)

				// update the variables
				markdown3 := fmt.Sprintf("---\r\nlayout: default\r\ntitle: %s\r\nparent: %s\r\n---\r\n\r\n"+
					"## %s (%s) <img align=\"right\" src=\"%s\" alt=\"'%s' achievement icon\" width=\"96\" height=\"96\">\r\n\r\n_%s_\r\n\r\n---%s",
					achieveTextH3, gameTitleRaw, achieveTextH3, achievePercent, imageSrc, achieveTextH3, achieveTextH5, achievements_guide)

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
