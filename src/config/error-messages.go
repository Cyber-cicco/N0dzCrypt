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

