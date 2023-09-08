package cmd_utils

import "fmt"

func AskYesNo(question string) bool {
	var response string
	fmt.Print(fmt.Printf("%s press (Y/n)", question))
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	switch response {
	case "Y":
		return true
	case "n":
		return false
	case "N":
		return false
	default:
		return false
	}
}
