package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	videoID := flag.String("video", "123456", "album/photo id")
	text := flag.String("text", "test from jmapi-go", "comment text")
	flag.Parse()

	username := os.Getenv("JM_USERNAME")
	password := os.Getenv("JM_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("please set JM_USERNAME and JM_PASSWORD")
	}

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	if _, err := client.Login(username, password); err != nil {
		log.Fatalf("login failed: %v", err)
	}

	resp, err := client.AlbumComment(*videoID, *text, "", "true", "")
	if err != nil {
		log.Fatalf("AlbumComment failed: %v", err)
	}

	fmt.Printf("comment resp: %+v\n", resp)
}
