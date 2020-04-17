package main

import "fmt"

func hello(user string) string {
	if len(user) == 0 {
		return fmt.Sprintf("Hello Dude")
	}  else {
		return fmt.Sprintf("Hello %v", user)
	}
}
