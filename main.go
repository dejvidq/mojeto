package main

import (
	"fmt"
	"os"
	"path"

	"github.com/dejvidq/mojeto/options"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Show help")
		return 
	}
	arg := os.Args[1]
	switch arg {
	case "init":
		home, _ := os.UserHomeDir()
		conf_path := ""
		if len(os.Args) == 3 {
			conf_path = path.Join(os.Args[2], "mojeto.conf")
		} else {
			conf_path = path.Join(home, ".config", "mojeto.conf")
	    }
        options.Init(conf_path)
	case "backup":
		options.Backup()
	default:
		fmt.Println("Show help")
	}
}
