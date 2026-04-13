package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	photoID := flag.String("id", "654321", "photo/chapter id")
	flag.Parse()

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	photo, err := client.GetPhotoDetail(*photoID, true, true)
	if err != nil {
		log.Fatalf("GetPhotoDetail failed: %v", err)
	}

	fmt.Printf("photo id: %s\n", photo.ID)
	fmt.Printf("album id: %s\n", photo.AlbumID)
	fmt.Printf("name: %s\n", photo.Name)
	fmt.Printf("scramble id: %s\n", photo.ScrambleID)
	fmt.Printf("image count: %d\n", len(photo.PageArr))
}
