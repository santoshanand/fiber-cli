package utils

import (
	"fmt"
	"os"
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0777); err != nil {
			fmt.Println("Unable to create directory")
		}
		return nil
	} else {
		fmt.Println("directory already exists")
		return err
	}
}
