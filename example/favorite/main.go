package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	page := flag.Int("page", 1, "page")
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

	res, err := client.FavoriteFolder(*page, jmapi.OrderByLatest, "0", username)
	if err != nil {
		log.Fatalf("FavoriteFolder failed: %v", err)
	}

	fmt.Printf("favorite total: %d\n", res.Total)
	for i, item := range res.Items {
		if i >= 20 {
			break
		}
		fmt.Printf("%d) [%s] %s\n", i+1, item.ID, item.Name)
	}
}
