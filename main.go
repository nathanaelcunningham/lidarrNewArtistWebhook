package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	config Config
}

func main() {
	SetupLogger()
	defaultPort := 8080

	cfg := LoadConfig()
	app := &App{config: cfg}

	r := http.NewServeMux()

	r.HandleFunc("POST /webhook", app.NewArtist)

	log.Printf("Listening on port %d\n", defaultPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), r); err != nil {
		log.Fatal(err)
	}
}

type ArtistWebhookPayload struct {
	Artist    ArtistData `json:"artist"`
	EventType string     `json:"eventType"`
}

type ArtistData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	MBID string `json:"mbId"`
}

func (a *App) NewArtist(w http.ResponseWriter, r *http.Request) {
	var payload ArtistWebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error decoding payload: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dirs, err := os.ReadDir(a.config.LibraryPath)
	if err != nil {
		log.Println("Error reading library directory", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistExists := false
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		if dir.Name() == payload.Artist.Name {
			artistExists = true
		}
	}

	if !artistExists {
		err := os.Mkdir(fmt.Sprintf("%s/%s", a.config.LibraryPath, payload.Artist.Name), 0777)
		if err != nil {
			log.Println("Error creating artist directory", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Printf("Artist %s Folder created", payload.Artist.Name)
}
