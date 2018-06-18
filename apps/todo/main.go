package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

const notesKey = "notes"
const form = `<form action="/" method="post">
<textarea name="note"></textarea><br />
<input type="submit" value="Submit" />
</form>`
const baseTemplate = `<html>
<head><title>TODO app</title></head>
<body>
<h1>TODO App</h1>
%s
</body>
</html>
`

func storeNote(db *redis.Client, note string) (int, string) {
	if note == "" {
		return http.StatusBadRequest, failPage("empty note")
	}

	err := db.RPush(notesKey, note).Err()
	if err != nil {
		return http.StatusInternalServerError, failPage(err.Error())
	}

	code, response := getNotes(db)
	return code, response
}

func getNotes(db *redis.Client) (int, string) {
	notes, err := db.LRange(notesKey, 0, -1).Result()
	if err != nil {
		return http.StatusInternalServerError, failPage(err.Error())
	}

	return http.StatusOK, formPage(notes)
}

func formPage(notes []string) string {
	notesInHTML := make([]string, len(notes))
	for i, note := range notes {
		notesInHTML = append(notesInHTML, fmt.Sprintf("<li>%d: %s</li>", i+1, note))
	}
	contents := []string{
		form,
		"<ul>",
		strings.Join(notesInHTML, "\n"),
		"</ul>",
	}
	return fmt.Sprintf(baseTemplate, strings.Join(contents, "\n"))
}

func failPage(message string) string {
	contents := []string{
		"<p>",
		message,
		"</p>",
		"<p><a href=\"/\">Back</a></p>",
	}
	return fmt.Sprintf(baseTemplate, strings.Join(contents, "\n"))
}

func createRootHandler(db *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var code int
		var response string
		if r.Method == http.MethodPost {
			note := r.FormValue("note")
			code, response = storeNote(db, note)
		} else {
			code, response = getNotes(db)
		}
		w.WriteHeader(code)
		fmt.Fprint(w, response)
	}
}

func redisClient() *redis.Client {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		panic("No Redis address provided!")
	}
	password := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("NOT_OK") == "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "KO")
	}
}

func main() {
	redisC := redisClient()
	http.HandleFunc("/", createRootHandler(redisC))
	http.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
