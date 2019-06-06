package add

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Add(args []string) {
	// If the file doesn't exist, create it, or append to the file
	path := viper.GetString("configDir")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
	f, err := os.OpenFile(viper.GetString("recordFileName"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, endpoint := range args {
		if _, err := f.Write([]byte(endpoint + "\n")); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
