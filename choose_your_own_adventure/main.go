package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web-application")
	filename := flag.String("file", "gopher.json", "File in JSON with a story")
	flag.Parse()
	fmt.Printf("Using the story in %s \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
