package utils

import (
	"os/user"
)

func GetHomeDir() string {

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return currentUser.HomeDir
}
