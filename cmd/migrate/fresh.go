package migration

import (
	"github.com/urfave/cli/v2"
)

func Fresh() *cli.Command {

	return &cli.Command{
		Name:  "fresh",
		Usage: "Rollback all migrations and re-run them",
		Action: func(c *cli.Context) error {
			migrations, err := InitMigration()
			if err != nil {
				return err
			}

			if err := migrations.Down(c.Context); err != nil {
				return err
			}

			if err := migrations.Up(c.Context); err != nil {
				return err
			}

			return nil
		},
	}
}
