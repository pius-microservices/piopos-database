package cmd

import (
	"go-gin/database/migrations/product"
	"go-gin/database/migrations/user"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "DB Migration",
	Long: "Database migration service",
}

func init() {
	initCommand.AddCommand(user.UserMigrationCMD)
	initCommand.AddCommand(product.ProductMigrationCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}