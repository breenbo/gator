package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/breenbo/gator/internal/xml"
)

func HandleAgg(s *State, cmd Command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"

	xml, err := xml.FetchFeed(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(xml)

	return nil
}
