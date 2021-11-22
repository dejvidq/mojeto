package options

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"errors"

	"github.com/manifoldco/promptui"
)

func yesNo() bool {
    prompt := promptui.Select{
        Label: "Select[Yes/No]",
        Items: []string{"Yes", "No"},
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result == "Yes"
}

func createConfigFile(conf_path string) {
	_, err := os.OpenFile(conf_path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	abs_path, _ := filepath.Abs(conf_path)
	// TODO create base config content and write to file
	fmt.Printf("Mojeto config initialized in: %s\n", abs_path)
}

func Init(conf_path string) {
	if _, err := os.Stat(conf_path); errors.Is(err, os.ErrNotExist) {
		createConfigFile(conf_path)
    } else {
		fmt.Printf("Config in path %s already exist. Do you want to override it with new config?\n", conf_path)
		var resp bool
		resp = yesNo()
		if resp == true {
			fmt.Println("Here override this file")
		}
    }
}
