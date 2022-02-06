package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func openBrowserTab(url string) {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		log.Printf("Error executing browser open command: %s", err)
	}
}

func makePath(e string) string {
	return filepath.Join(defaultPath, e)
}

func appdataDir() string {
	t, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func openExplorer(file string) {
	cmd := exec.Command(`explorer`, `/select,`, file)
	cmd.Run()

}

func openMyFolder() {
	openExplorer(defaultPath + "\\" + defaultDirName + "\\data.json")
}
