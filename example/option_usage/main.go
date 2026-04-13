package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	albumID := flag.String("aid", "123456", "album id")
	photoID := flag.String("pid", "123456", "photo id (for save dir demo)")
	optPath := flag.String("opt", "", "option yml path")
	flag.Parse()

	var opt jmapi.Option
	var err error
	if *optPath == "" {
		opt = jmapi.DefaultOption()
	} else {
		opt, err = jmapi.LoadOption(*optPath)
		if err != nil {
			log.Fatalf("load option failed: %v", err)
		}
	}

	client := opt.NewClient()
	album, err := client.GetAlbumDetail(*albumID)
	if err != nil {
		log.Fatalf("GetAlbumDetail failed: %v", err)
	}

	photo, err := client.GetPhotoDetail(*photoID, false, false)
	if err != nil {
		log.Fatalf("GetPhotoDetail failed: %v", err)
	}

	saveDir, err := opt.DecideImageSaveDir(*album, *photo)
	if err != nil {
		log.Fatalf("DecideImageSaveDir failed: %v", err)
	}

	fmt.Println("== option usage ==")
	fmt.Printf("client type: %s\n", opt.ClientConfig.ClientType)
	fmt.Printf("dir rule: %s\n", opt.DirRule.Rule)
	fmt.Printf("base dir: %s\n", opt.DirRule.BaseDir)
	fmt.Printf("image suffix: %q\n", opt.Download.Image.Suffix)
	fmt.Printf("sample filename(index=1): %s\n", opt.DecideImageFilename(1))
	fmt.Printf("resolved save dir: %s\n", saveDir)
}
