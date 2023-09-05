package main

import "fmt"

func main() {
	const name = "Mario Jorge"
	openRate := 30.8

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)
}