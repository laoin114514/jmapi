package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	albumID := flag.String("id", "123456", "album id")
	flag.Parse()

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	album, err := client.GetAlbumDetail(*albumID)
	if err != nil {
		log.Fatalf("GetAlbumDetail failed: %v", err)
	}

	fmt.Printf("album id: %s\n", album.ID)
	fmt.Printf("name: %s\n", album.Name)
	fmt.Printf("author: %v\n", album.Author)
	fmt.Printf("tags: %v\n", album.Tags)
	fmt.Printf("likes/views: %s/%s\n", album.Likes, album.Views)
}
