package main

import(
	"github.com/miraikeitai2020/backend-spot/pkg/config"
	"github.com/miraikeitai2020/backend-spot/pkg/server"
	"github.com/miraikeitai2020/backend-spot/pkg/server/controller"
)

func main() {
	addr := config.GetRouterAddr()

	// TODO: database connection function exec here.
	// v0.2.0

	ctrl := controller.Init(nil)
	if err := server.Router(ctrl).Run(addr); err != nil {
		panic(err)
	}
}
