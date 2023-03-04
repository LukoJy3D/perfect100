package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

			markdown := fmt.Sprintf("<img style=\"float: right;\" src=\"%s\" width=\"128\" height=\"128\"> \n\n## %s ([guide](%s))\n\nOwned by **%s** of players\n\n_%s_\n\n---\n\n",
				imageSrc, achieveTextH3, guidePath, achievePercent, achieveTextH5)

			if _, err = achievementList.WriteString(markdown); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			// Replace any spaces or other common characters
			achieveTextH3 = ReplaceForbiddenCharacters(achieveTextH3)

			// Create a file with the achievement title as the name
			f2, err := os.OpenFile("guides/"+gameTitle+"/"+achieveTextH3+".md", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			defer f2.Close()
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
