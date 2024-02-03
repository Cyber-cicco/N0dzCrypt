package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "nsc",
  Short: "Utilitaire d'organisation et de génération de code pour les applications web basées sur Spring boot, Thymeleaf, Tailwind et HTMX",
  Long: 
`
Utilitaire d'organisation et de génération de code pour les applications web basées sur Spring boot, Thymeleaf, Tailwind et HTMX.
Offre un utilitaire permettant de générer des applications basées sur des recommandations d'organisation de code pensées par un développeur junior atteint d'une forme aigüe d'autisme.
Offre un package manager permettant de récupérer des composants basées sur le stack sus-mentionnées pour les intégrer au code de votre application.
Offre un utilitaire permettant de générer des éléments spécifiques à l'organisation de code 
`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
    rootCmd.AddCommand(cnsa)
}


