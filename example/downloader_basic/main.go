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

	opt := jmapi.DefaultOption()
	d := jmapi.NewDownloader(opt)

	album, err := d.DownloadAlbum(*albumID)
	if err != nil {
		log.Fatalf("DownloadAlbum failed: %v", err)
	}

	if err := d.RaiseIfHasFailures(); err != nil {
		log.Printf("download finished with partial failures: %v", err)
	}

	fmt.Println("== downloader summary ==")
	fmt.Printf("album: [%s] %s\n", album.ID, album.Name)
	fmt.Printf("success photo count: %d\n", len(d.SuccessImages))
	fmt.Printf("failed photo count: %d\n", len(d.FailedPhotos))
	fmt.Printf("failed image count: %d\n", len(d.FailedImages))
}
