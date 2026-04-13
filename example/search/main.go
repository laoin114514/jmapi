package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	query := flag.String("q", "laoin", "search query")
	page := flag.Int("page", 1, "page")
	flag.Parse()

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	res, err := client.SearchSite(*query, *page, jmapi.OrderByLatest, jmapi.TimeAll, jmapi.CategoryAll, "")
	if err != nil {
		log.Fatalf("SearchSite failed: %v", err)
	}

	fmt.Printf("total: %d\n", res.Total)
	for i, item := range res.Items {
		if i >= 10 {
			break
		}
		fmt.Printf("%d) [%s] %s\n", i+1, item.ID, item.Name)
	}
}
