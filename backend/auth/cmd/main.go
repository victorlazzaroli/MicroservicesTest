package main

import (
	"fmt"

	"github.com/victorlazzaroli/microservicesTest/auth/config"
	"github.com/victorlazzaroli/microservicesTest/auth/controllers"
	"github.com/victorlazzaroli/microservicesTest/auth/startup"
)

func main() {
	env := config.NewEnv(".env")
	db := startup.ConnectDatabase()
	startup.SyncDB(db)
	server := startup.NewServer(env)

	controllers.NewUsersController(server, db, env)

	server.Run(fmt.Sprint(":", env.ServerPort))
}
