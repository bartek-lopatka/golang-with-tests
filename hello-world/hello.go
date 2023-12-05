package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return "Hello, world!"
	}
	return englishPrefix + name + "!"
}
func main() {
	fmt.Println(Hello("Chris"))
}
