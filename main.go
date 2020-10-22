package main

import(
	"log"

	"github.com/miraikeitai2020/backend-spot/pkg/config"
	"github.com/miraikeitai2020/backend-spot/pkg/server"
	"github.com/miraikeitai2020/backend-spot/pkg/database"
	"github.com/miraikeitai2020/backend-spot/pkg/server/controller"
)

func main() {
	addr := config.GetRouterAddr()

	// TODO: database connection function exec here.
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	ctrl := controller.Init(db)
	if err := server.Router(ctrl).Run(addr); err != nil {
		panic(err)
	}
}
