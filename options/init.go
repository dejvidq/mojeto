package options

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
)

func Init(conf_path string) {
	fmt.Printf("Mojeto config initialized in: %s\n", conf_path)
	fmt.Println(fmt.Println(filepath.Abs(conf_path)))
	_, err := os.OpenFile(conf_path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
