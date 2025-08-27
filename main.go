package main

import (
	"log"

	cli "github.com/breenbo/gator/internal/cli"
	config "github.com/breenbo/gator/internal/config"
	"github.com/breenbo/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	config := config.Read()
	dbQueries := database.InitDatabase(config.Db_url)
	state := cli.State{
		Db:  dbQueries,
		Cfg: &config,
	}

	c := cli.RegisterFn()
	cmd := cli.GetUserEntries()

	if err := c.Run(&state, cmd); err != nil {
		log.Fatal(err)
	}
}
