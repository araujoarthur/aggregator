package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/araujoarthur/aggregator/internal/config"
	"github.com/araujoarthur/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	configData, err := config.Read()
	if err != nil {
		log.Fatalln(err.Error())
	}
	stateData := state{Config: &configData}
	db, err := sql.Open("postgres", stateData.Config.DBUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	stateData.DbQueries = database.New(db)

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	stateData.RegisteredCommands = &cmds

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleListUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", handleListFeeds)
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Invalid command arguments")
	}

	cmd := command{Name: os.Args[1], Args: os.Args[2:]}

	if err = cmds.run(&stateData, cmd); err != nil {
		log.Fatal(err)
	}
}
