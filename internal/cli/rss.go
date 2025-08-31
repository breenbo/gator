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
	if len(cmd.Arguments) == 0 {
		log.Fatal("time duration needed")
	}

	timeDuration := cmd.Arguments[0]
	duration, err := time.ParseDuration(timeDuration)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		fmt.Print("\n\nget feed\n")
		if err := scrapeFeeds(s); err != nil {
			return err
		}
	}
}

func HandleAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 2 {
		log.Fatal("need feed name and url")
	}

	fmt.Print("get the feed\n")

	ctx := context.Background()

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

	_, feedErr := createFollowFeed(ctx, s, url)
	if feedErr != nil {
		return feedErr
	}

	fmt.Printf("%v\n", feed)

	return nil
}

func HandleListFeed(s *State, cmd Command) error {
	ctx := context.Background()

	feeds, err := s.Db.ListFeed(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s - %s from %s\n", feed.Name, feed.Url, feed.Username)
	}

	return nil
}

func HandleFollowFeed(s *State, cmd Command) error {
	ctx := context.Background()

	if len(cmd.Arguments) == 0 {
		log.Fatal("need url")
	}
	url := cmd.Arguments[0]

	value, err := createFollowFeed(ctx, s, url)
	if err != nil {
		return err
	}

	fmt.Print(value)

	return nil
}

func HandleFollowingFeed(s *State, cmd Command, user database.User) error {
	ctx := context.Background()

	userID := user.ID
	followings, err := s.Db.GetFeedFollowsForUser(ctx, userID)
	if err != nil {
		return err
	}

	for _, following := range followings {
		fmt.Printf("%s", following.FeedName)
	}

	return nil
}

// helpers
func createFollowFeed(ctx context.Context, s *State, url string) (database.CreateFeedFollowRow, error) {
	now := time.Now()
	val := database.CreateFeedFollowRow{}

	feedID, err := s.Db.GetFeedIDFromURL(ctx, url)
	if err != nil {
		return val, err
	}
	userID, err := s.Db.GetUser(ctx, s.Cfg.Current_user_name)
	if err != nil {
		return val, err
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New().String(),
		CreatedAt: now,
		UpdatedAt: now,
		FeedID:    feedID,
		UserID:    userID.ID,
	}
	value, err := s.Db.CreateFeedFollow(ctx, newFollow)
	if err != nil {
		return val, err
	}

	return value, nil
}

func HandleUnfollow(s *State, cmd Command, user database.User) error {
	ctx := context.Background()
	if len(cmd.Arguments) == 0 {
		log.Fatal("url needed")
	}

	del := database.DeleteFollowParams{
		Url:    cmd.Arguments[0],
		UserID: user.ID,
	}

	if err := s.Db.DeleteFollow(ctx, del); err != nil {
		return err
	}

	return nil
}

func scrapeFeeds(s *State) error {
	ctx := context.Background()
	// get the next feed to fetch
	next_feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	// mark the next feed as fetched
	value := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        next_feed.ID,
	}
	if err := s.Db.MarkFeedFetched(ctx, value); err != nil {
		return err
	}

	news, err := xml.FetchFeed(ctx, next_feed.Url)
	if err != nil {
		return err
	}

	for _, new := range news.Channel.Item {
		fmt.Printf("%s", new.Title)
	}

	return nil
}
