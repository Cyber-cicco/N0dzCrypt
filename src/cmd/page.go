package cmd

import (
	"fmt"
	"fr/nzc/service/page"

	"github.com/spf13/cobra"
)

var createPageCmd = &cobra.Command{
	Use:   "page",
	Short: "add a page to your project",
	Long: `
        Add a working page to your project. By default, it will create a new HTML file, a new Controller, a new Root and a new Irrigator.
    `,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("args: %v\n", args)
        page.CreatePage(args)
    },
}

func init() {
}
