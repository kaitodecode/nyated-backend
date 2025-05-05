package cmd

import (
	"time"

	"github.com/joho/godotenv"
	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/config"
	"github.com/kaitodecode/nyated-backend/database/seeder"
	"github.com/spf13/cobra"
)

var commandSeeder = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		config.Init()
		AppError.Init()
		db, err := config.InitDatabase()

		if err != nil {
			panic(err)
		}

		loc, err := time.LoadLocation("Asia/Jakarta")

		if err != nil {
			panic(err)
		}

		time.Local = loc


		seeder.NewSeederRegistry(db).Run()
	},
}

func RunSeeder(){
	commandSeeder.Execute()
}