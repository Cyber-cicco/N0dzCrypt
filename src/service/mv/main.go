package mv

import (
	"errors"
	"fr/nzc/config"
	"fr/nzc/thymeleaf"
	"fr/nzc/utils"
)

func MovePage(args []string, flag string) {
    thymeleaf.TestParser()
    switch flag {
    case "" : {

    }
    case "page" : {

    }
    case "comp" : {

    }
    case "frag" : {

    }
    case "layout" : {

    }
    case "tstyle" : {

    }
    case "tscript" : {

    }
    case "svg" : {

    }
    case "style" : {

    }
    case "script" : {

    }
    case "img" : {

    }
    default : {
        utils.HandleUsageError(errors.New("Unsupported file type"), config.ERR_BAD_MV_ARGS)
    }
    }
}
