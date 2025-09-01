package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/Darkpowercross/gator/internal/config"
	"github.com/Darkpowercross/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &config.State{DB: dbQueries, Config: &cfg}

	cmds := &config.Commands{}
	cmds.Register("login", config.HandlerLogin)
	cmds.Register("register", config.HandlerRegister)
	cmds.Register("reset", config.HandlerReset)
	cmds.Register("users", config.HandlerUsers)
	cmds.Register("agg", config.HandlerAgg)
	cmds.Register("addfeed", config.MiddlewareLoggedIn(config.HandlerAddFeed))
	cmds.Register("feeds", config.Handlerfeeds)
	cmds.Register("follow", config.MiddlewareLoggedIn(config.HandlerFollow))
	cmds.Register("following", config.MiddlewareLoggedIn(config.HandlerFollowing))
	cmds.Register("unfollow", config.MiddlewareLoggedIn(config.HandlerUnfollow))
	cmds.Register("browse", config.MiddlewareLoggedIn(config.HandlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}

	name := os.Args[1]
	args := os.Args[2:]
	cmd := config.Command{Name: name, Args: args}

	if err := cmds.Run(s, cmd); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
