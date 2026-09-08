package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("GoEmOS v0.1.0")
	fmt.Println("Operating system emulator")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("user@goemos:~\n❯ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Printf("command: %s", command)
	}
}
