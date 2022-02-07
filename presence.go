package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
	"tawesoft.co.uk/go/dialog"
)

func loadPresence(data JSONData) {

	if data.DiscordClientId == "DISCORD-CLIENT-ID" {
		dialog.Alert("Please, change Discord Application ID inside data.json file")
		openBrowserTab(changeDataWebsite)

	} else {

		clientIPC := client.Login(data.DiscordClientId)

		updateActivity(data)

		if clientIPC != nil {
			dialog.Alert("Please, change Discord Application ID inside data.json file")
		}

		time.Sleep(time.Second * 1000)
	}
}

func updateActivity(data JSONData) {

	buttons := []*client.Button{}

	for i := 0; i < len(data.RichPresence.Buttons); i++ {
		var singleButton = (*client.Button)(&data.RichPresence.Buttons[i])
		buttons = append(buttons, singleButton)
	}

	client.SetActivity(client.Activity{
		State:      data.RichPresence.State,
		Details:    data.RichPresence.Details,
		LargeImage: data.RichPresence.BigImageName,
		LargeText:  data.RichPresence.BigImageText,
		SmallImage: data.RichPresence.SmallImageName,
		SmallText:  data.RichPresence.SmallImageText,

		Buttons: buttons,
	})

}
