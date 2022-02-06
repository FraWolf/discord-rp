package main

type JSONData struct {
	DiscordClientId string       `json:"discordClientId"`
	RichPresence    RichPresence `json:"richPresence"`
}

type RichPresence struct {
	State          string   `json:"state"`
	Details        string   `json:"details"`
	BigImageName   string   `json:"bigImageName"`
	BigImageText   string   `json:"bigImageText"`
	SmallImageName string   `json:"smallImageName"`
	SmallImageText string   `json:"smallImageText"`
	Buttons        []Button `json:"buttons"`
}

type Button struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type ReturnMessage struct {
	Message string `json:"message"`
}
