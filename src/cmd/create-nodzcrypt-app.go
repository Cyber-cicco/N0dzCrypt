package cmd

import (
	"fr/nzc/service/cna"

	"github.com/spf13/cobra"
)

var cnsa = &cobra.Command{
	Use:   "create-nodzcrypt-app",
	Short: "Create a N0dzCrypt application",
	Long: `
    Allows you to create a new N0dzCrypt application with the necessary dependencies to be able to start coding your app
    Installs Tailwind, Spring boot, Spring boot starter security, Spring boot starter web, Thymeleaf, Lombok, io.jsonwebtoken.
    It also creates a base template file with the needed headers, like the link to the HTMX library, and a link to your output css.
    It then create a basic web page with a tutorial on how to use N0dzCrypt.
  `,

	Run: func(cmd *cobra.Command, args []string) {
        cna.CreateNodzCryptApp()
    },
}

func init() {
}
