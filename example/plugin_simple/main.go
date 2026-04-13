package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

type LoggerPlugin struct{ jmapi.PluginAdapter }

func (p LoggerPlugin) Key() string { return "logger-plugin" }

func (p LoggerPlugin) AfterInit(ctx jmapi.PluginContext) error {
	fmt.Println("[plugin] after init")
	return nil
}

func (p LoggerPlugin) BeforeAlbum(ctx jmapi.PluginContext, album *jmapi.AlbumDetail) error {
	fmt.Printf("[plugin] before album: %s\n", album.ID)
	return nil
}

func (p LoggerPlugin) AfterAlbum(ctx jmapi.PluginContext, album *jmapi.AlbumDetail) error {
	fmt.Printf("[plugin] after album: %s\n", album.ID)
	return nil
}

func (p LoggerPlugin) BeforePhoto(ctx jmapi.PluginContext, photo *jmapi.PhotoDetail) error {
	fmt.Printf("[plugin] before photo: %s\n", photo.ID)
	return nil
}

func (p LoggerPlugin) AfterPhoto(ctx jmapi.PluginContext, photo *jmapi.PhotoDetail) error {
	fmt.Printf("[plugin] after photo: %s\n", photo.ID)
	return nil
}

func main() {
	albumID := flag.String("id", "123456", "album id")
	flag.Parse()

	opt := jmapi.DefaultOption()
	d := jmapi.NewDownloader(opt)
	d.RegisterPlugin(LoggerPlugin{})

	album, err := d.DownloadAlbum(*albumID)
	if err != nil {
		log.Fatalf("DownloadAlbum failed: %v", err)
	}

	fmt.Printf("download completed for album: %s\n", album.ID)
	if err := d.RaiseIfHasFailures(); err != nil {
		fmt.Printf("partial failures: %v\n", err)
	}
}
