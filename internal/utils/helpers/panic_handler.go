package helpers

import "fmt"

func HandlePanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
