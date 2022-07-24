package main

import (
	"fmt"
)

func main() {
	fmt.Println(PrintMessage("Juan"))
}

func PrintMessage(name string) string {
	return fmt.Sprintf("Bootcamper %s", name)
}
