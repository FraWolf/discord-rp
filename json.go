package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getJSONFile() {

	if _, err := os.Stat(makePath(defaultDirName)); os.IsNotExist(err) {
		err = os.Mkdir(makePath(defaultDirName), os.ModePerm)
		if err != nil {
			log.Printf("Couldn't create logs directory: %s", err)
			return
		}
		fillJSONData("DISCORD-CLIENT-ID", "Doing anything", "Sample details", "largeImageName", "largeImageText", "smallImageName", "smallImageText", []Button{
			{
				Label: "Button #1",
				Url:   mainWebsite,
			},
			{
				Label: "Button #2",
				Url:   mainWebsite,
			},
		})
	}

}

func fillJSONData(discordClientId string, state string, details string, bigImageName string, bigImageText string, smallImageName string, smallImageText string, buttons []Button) {

	data := JSONData{
		DiscordClientId: discordClientId,
		RichPresence: RichPresence{
			State:          state,
			Details:        details,
			BigImageName:   bigImageName,
			BigImageText:   bigImageText,
			SmallImageName: smallImageName,
			SmallImageText: smallImageText,
			Buttons:        buttons,
		},
	}

	file, err := os.Create(makePath(fmt.Sprintf(defaultDirName + "/data.json")))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte(jsonToText(data)))

}

func jsonToText(data interface{}) string {

	f, err := json.MarshalIndent(&data, " ", "  ")
	if err != nil {
		log.Println("Error in getting json")
	}

	return string(f)

}

func textToJsonData(data []byte) JSONData {
	var jsonFile JSONData

	json.Unmarshal([]byte(data), &jsonFile)

	return jsonFile

}

func readJson() JSONData {

	file, err := os.ReadFile(defaultPath + "/" + defaultDirName + "/data.json")
	if err != nil {
		panic(err)
	}

	return textToJsonData(file)
}

func readBody(req *http.Request) []byte {

	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	return bodyBytes

}
