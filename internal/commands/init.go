package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Init() error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Stat(wd + "/.wierze")
	if os.IsNotExist(err) {
		fmt.Println("wierze directory not present")
	} else {

		if !dir.IsDir() {
			return errors.New("directory present but it's a file")
		}

		return errors.New("wierze directory already present")
	}

	fmt.Println("Creating new repository base")

	err = os.Mkdir(wd+"/.wierze", 0o755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(wd+"/.wierze/objets", 0o755)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(wd + "/.wierze/HEAD")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return nil
}
