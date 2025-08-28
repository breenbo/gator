package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/breenbo/gator/internal/database"
	"github.com/breenbo/gator/internal/xml"
	"github.com/google/uuid"
)

func HandleAggregator(s *State, cmd Command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"

	xml, err := xml.FetchFeed(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(xml)

	return nil
}

func HandleAddFeed(s *State, cmd Command) error {
	if len(cmd.Arguments) < 2 {
		log.Fatal("need feed name and url")
	}

	fmt.Print("get the feed\n")

	ctx := context.Background()
	userName := s.Cfg.Current_user_name
	user, err := s.Db.GetUser(ctx, userName)
	if err != nil {
		return err
	}

	name := cmd.Arguments[0]
	url := cmd.Arguments[1]
	now := time.Now()
	newFeed := database.CreateFeedParams{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.Db.CreateFeed(ctx, newFeed)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	return nil
}
