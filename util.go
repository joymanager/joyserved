package main

import (
	"fmt"
	"os"
	"runtime"
)

/** GetConfigFolder retrieves the folder where joyserve configuration data is stored. */
func GetConfigFolder() string {
	var config string
	if runtime.GOOS == "windows" {
		config = os.Getenv("APPDATA")
	}

	return fmt.Sprintf("%s%c%s", config, os.PathSeparator, "joyserve")
}

func GetLibraryFolders() []string {

}
