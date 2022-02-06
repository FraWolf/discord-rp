package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getlantern/systray"
)

var (
	defaultPath    string = appdataDir()
	defaultDirName string = "frawolf-discord-rp"

	server = &http.Server{Addr: "localhost:1709", Handler: nil}
)

func main() {

	getJSONFile()
	startWebServer()
	systray.Run(onReady, onExit)

}

func onReady() {
	systray.SetIcon(getIcon("assets/clock.ico"))
	systray.SetTitle(name)
	systray.SetTooltip(name)
	changeData := systray.AddMenuItem("Change data", "Change rich presence's data")
	systray.AddMenuItem("frawolf.dev", "frawolf.dev").Disable()
	systray.AddSeparator()
	quitBtn := systray.AddMenuItem("Close", "Quit the whole app")

	go func() {
		for {
			select {
			case <-changeData.ClickedCh:
				openBrowserTab(changeDataWebsite)
			case <-quitBtn.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	data := readJson()

	loadPresence(data)

}

func onExit() {
	systray.Quit()
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
