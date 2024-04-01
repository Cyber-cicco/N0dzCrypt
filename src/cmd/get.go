package cmd

import (
	"github.com/spf13/cobra"
)

var get = &cobra.Command{
  Use:   "get",
  Short: "Allows you to get a part of another nodzcrypt application using it's github adress",
  Long: 
`
Allows you to use the github adress of another n0dzcrypt app git repository. It will clone the repository, get the content you asked for, and copy it in your own app.
You can copy any 
`,
}

func init() {
}
