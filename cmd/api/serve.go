package api

import (
	migration "github.com/dimasbayuseno/cisdi-go-test/cmd/migrate"
	"github.com/dimasbayuseno/cisdi-go-test/config"
	"github.com/dimasbayuseno/cisdi-go-test/internal/server"
	dbpostgres "github.com/dimasbayuseno/cisdi-go-test/pkg/db/postgres"
	"github.com/urfave/cli/v2"
)

func Serve() *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Run the API server",
		Action: func(c *cli.Context) error {
			_, err := config.LoadConfig()
			if err != nil {
				return err
			}

			_, err = config.ParseConfig(config.GetViper())
			if err != nil {
				return err
			}

			db, err := dbpostgres.NewPgx()
			if err != nil {
				return err
			}

			migrations, err := migration.InitMigration()
			if err != nil {
				return err
			}

			if err := migrations.Up(c.Context); err != nil {
				return err
			}

			server := server.New(
				db,
			)
			return server.Run()
		},
	}

}
