package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nlopes/slack"
)

// Customize the slashCommand to anything
const slashCommand = "/hello"

func slashCommandHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !s.ValidateToken(os.Getenv("SLACK_VERIFICATION_TOKEN")) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch s.Command {
		case slashCommand:
			log.Printf("%s slack command is issued\n", slashCommand)
			// Add code here!
			response := "Hello World!"
			// Create a JSON response
			w.Header().Set("Content-Type", "application/json")
			jsonResponse := make(map[string]string)
			jsonResponse["response_type"] = "in_channel"
			jsonResponse["text"] = response
			jsonResponseMarshal, err := json.Marshal(jsonResponse)
			if err != nil {
				log.Fatalf("error while JSON marshalling response: %s", err)
			}

			w.Write([]byte(jsonResponseMarshal))

		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {

	http.HandleFunc("/receive", slashCommandHandler())

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT environment variable not found defaulting to port %s", port)
	}

	log.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
