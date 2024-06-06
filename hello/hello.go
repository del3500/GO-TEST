package main

import "fmt"

func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s", name)
	}
	return "Hello, World"
}
