package main

import (
	"github.com/dimasbayuseno/cisdi-go-test/cmd/api"
	migration "github.com/dimasbayuseno/cisdi-go-test/cmd/migrate"
	"github.com/dimasbayuseno/cisdi-go-test/config"
	"github.com/urfave/cli/v2"
	"os"
)

// @title github.com/dimasbayuseno/cisdi-go-test
// @version 1.0
// @description This is a sample server cell for github.com/dimasbayuseno/cisdi-go-test.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.synapsis.id
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath
// Run initializes whole application.

func main() {
	appCli := cli.NewApp()
	appCli.Name = config.Get().Service.Name
	appCli.Commands = []*cli.Command{
		migration.Root(),
		api.Serve(),
	}

	if err := appCli.Run(os.Args); err != nil {
		panic(err)
	}
}
