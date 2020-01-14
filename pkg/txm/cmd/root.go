package cmd

import (
	"fmt"
	"os"

	"github.com/Bo0km4n/arc/pkg/txm/driver"
	"github.com/Bo0km4n/arc/pkg/txm/executor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "txm",
	Short: "txm",
	Long:  "txm",
}

func init() {
	rootCmd.AddCommand(driver.NewDriverCmd())
	rootCmd.AddCommand(executor.NewExecutorCmd())
}

// Execute is entry points
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
