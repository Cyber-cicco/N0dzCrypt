package config

/*
*   Liste des différents messages d'erreur possibles
*/

var ERR_BAD_ARGS = "Incorrect usage of the command"
var ERR_FILE_CREATION = "Error in creating a file"
var ERR_DIR_CREATION = "Error in creating a directory"
var ERR_MARSHARL = "Error in deserializing an object from JSON"
var ERR_OPEN_CONFIG = "Error opening the configuration file. If it doesn't exist, please use the init command to create it.\nOtherwise, make sure to run the command in a directory containing a valid configuration file."
var ERR_UNMARSHAL = "Error in creating an object from JSON"
var ERR_BAD_CONFIG_PACKAGE = "Error: the package specified in the configuration file does not seem to point to an existing directory."
var ERR_TEMPLATE_FILE_READ = "Error reading a template file"
var ERR_CURR_DIR_OPEN = "Error openening current directory"
var ERR_COULDNT_READ_INPUT = "Could not read user input"
var ERR_COULDNT_FIND_CONFIG = "Could'nt find any configuration file for N0dzCrypt. Make sure you use create-n0dzcrypt-app or nzc init to enable the n0dzcrypt CLI in your codebase" 
var ERR_BAD_MV_ARGS = `
Unsupported file type. N0dzCrypt only allows the following types of file to be renamed: 
    - comp : components of your application, aka thymeleaf fragments used globally in your application. Cannot be referenced in your backend routes.
    - page : page of your application, aka a standalone thymeleaf fragment that is not referenced anywhere else in your front end fragments. Can be referenced in your backend routes.
    - frag : fragment of a page of your application, aka a thymeleaf fragment that is only referenced in your pageor by other fragments of your page. 
    - layout : layout of your app, usually containing the headers and components that will not change when navigating from one page to another.
    - tstyle : a css file in your templates folder that is meant to be included inline in one or more fragments of your application. 
    - tscript : a javascript file in your templates folder that is meant to be included inline in one or more fragments of your application
    - svg : an html file containing a single svg, that can be inlined in any of your fragments.
    - style : a css file in your static folder that can be referenced through links in your html.
    - script :  a javascript file in your static folder that can be referenced through the "src" attribute in a script tag.
    - img : an image in your static folder that can be referenced through the "src" attribute in an img tag.
`

