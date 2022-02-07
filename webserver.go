package main

import (
	"fmt"
	"net/http"
)

func startWebServer() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		enableCors(&res, req)

		res.Header().Set("Content-Type", "application/json")
		fmt.Fprint(res, jsonToText(ReturnMessage{Message: "It works"}))
	})

	http.HandleFunc("/data", func(res http.ResponseWriter, req *http.Request) {
		enableCors(&res, req)

		data := readJson()

		res.Header().Set("Content-Type", "application/json")
		fmt.Fprint(res, jsonToText(data))
	})

	http.HandleFunc("/update", func(res http.ResponseWriter, req *http.Request) {
		enableCors(&res, req)

		if req.Method == "POST" {

			body := readBody(req)
			newJsonData := textToJsonData(body)

			updateActivity(newJsonData)
			fillJSONData(
				newJsonData.DiscordClientId,
				newJsonData.RichPresence.State,
				newJsonData.RichPresence.Details,
				newJsonData.RichPresence.BigImageName,
				newJsonData.RichPresence.BigImageText,
				newJsonData.RichPresence.SmallImageName,
				newJsonData.RichPresence.SmallImageText,
				newJsonData.RichPresence.Buttons)

			res.Header().Set("Content-Type", "application/json")
			fmt.Fprint(res, jsonToText(newJsonData))
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(res, jsonToText(ReturnMessage{Message: "Method not allowed"}))
		}

	})

	go func() {
		fmt.Println("[WEBSERVER] Listening on http://localhost:1709")
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Print("http server closed")
				return
			}
			fmt.Printf("Error with http server: %s", err)
		}
	}()

}

func enableCors(res *http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	allowedOrigins := []string{"https://rp.frawolf.dev", "http://localhost:3000", "http://192.168.178.48:3000"}
	for _, o := range allowedOrigins {

		if o == origin {
			(*res).Header().Set("Access-Control-Allow-Origin", origin)
			(*res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			(*res).Header().Set("Access-Control-Allow-Headers", "*")
			break
		}
	}
}
