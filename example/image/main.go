package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	albumID := flag.String("id", "123456", "album id")
	out := flag.String("out", "cover.jpg", "output file")
	flag.Parse()

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	data, err := client.DownloadAlbumCover(*albumID)
	if err != nil {
		log.Fatalf("DownloadAlbumCover failed: %v", err)
	}

	if err := os.WriteFile(*out, data, 0o644); err != nil {
		log.Fatalf("write file failed: %v", err)
	}

	fmt.Printf("saved cover to %s (%d bytes)\n", *out, len(data))
}
