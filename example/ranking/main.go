package main

import (
	"flag"
	"fmt"
	"log"

	jmapi "github.com/laoin114514/jmapi"
)

func main() {
	kind := flag.String("type", "month", "month|week|day")
	page := flag.Int("page", 1, "page")
	flag.Parse()

	client := jmapi.NewClient(jmapi.Config{
		ClientType:        jmapi.ClientTypeAPI,
		AutoUpdateHost:    true,
		AutoEnsureCookies: true,
	})

	var (
		res *jmapi.SearchResult
		err error
	)

	switch *kind {
	case "month":
		res, err = client.MonthRanking(*page, jmapi.CategoryAll)
	case "week":
		res, err = client.WeekRanking(*page, jmapi.CategoryAll)
	case "day":
		res, err = client.DayRanking(*page, jmapi.CategoryAll)
	default:
		log.Fatalf("unknown type: %s", *kind)
	}
	if err != nil {
		log.Fatalf("ranking failed: %v", err)
	}

	for i, item := range res.Items {
		if i >= 10 {
			break
		}
		fmt.Printf("%d) [%s] %s\n", i+1, item.ID, item.Name)
	}
}
