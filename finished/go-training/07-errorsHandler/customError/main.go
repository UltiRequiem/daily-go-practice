package main

import (
	"fmt"
)

func checkUserNameExist(username string) (bool, error) {
	if username == "UltiRequiem" {
		return true, ErrUserNameExist{UserName: username}
	}

	return false, nil
}

func main() {
	if _, err := checkUserNameExist("foo"); err == nil {
		fmt.Println("foo not exist")
	}

	if _, err := checkUserNameExist("UltiRequiem"); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := checkUserNameExist("UltiRequiem"); err != nil {
		if ok := IsErrUserNameExist(err); ok {
			fmt.Println("User UltiRequiem already exist.")
		}
	}
}
