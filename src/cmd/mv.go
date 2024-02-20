package cmd

import (
	"fr/nzc/service/mv"

	"github.com/spf13/cobra"
)

var fileType string

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "move or rename a file",
	Long: `
Move or rename an element of your frontend n0dzcrypt app, updating any reference in your backend.
By default, it takes two arguments : the relative path of the file you want to rename, and the relative path of the new name of the file.
It will rename the file, and search for any reference in your front end fragments and the Routes.java file in your backend and update it accordingly.
You can also pass an extra argument with the --type or -t flag. This will enforce some behaviours, and will now take paths relative to the folder you have defined in n0dzcrypt.json for the type of file you're moving.
It has to be one of the following :
    - comp : components of your application, aka thymeleaf fragments used globally in your application. Cannot be referenced in your backend routes, so it won't update the Routes.java file
    - page : page of your application, aka a standalone thymeleaf fragment that is not referenced anywhere else in your front end fragments. Can be referenced in your backend routes. That way, it won't update any reference in your front-end file, but will update the Routes.java
    - frag : fragment of a page of your application, aka a thymeleaf fragment that is only referenced in your pageor by other fragments of your page. Will update every reference in your page folder and Routes.java 
    - layout : layout of your app, usually containing the headers and components that will not change when navigating from one page to another. Will update the Routes.java
    - tstyle : a css file in your templates folder that is meant to be included inline in one or more fragments of your application. Will update any file in your templates folder
    - tscript : a javascript file in your templates folder that is meant to be included inline in one or more fragments of your application. Will update any file in your templates folder
    - svg : an html file containing a single svg, that can be inlined in any of your fragments. Will update any file in your templates folder
    - style : a css file in your static folder that can be referenced through links in your html. Will update any file in your templates folder
    - script :  a javascript file in your static folder that can be referenced through the "src" attribute in a script tag. Will update any file in your templates folder
    - img : an image in your static folder that can be referenced through the "src" attribute in an img tag. Will update any file in your templates folder

    `,
    Run: func(cmd *cobra.Command, args []string) {
        mv.MovePage(args, fileType)
    },
}

func init() {
    mvCmd.Flags().StringVarP(&fileType, "type", "t", "", "Type de la classe que vous souhaitez créer")
}
