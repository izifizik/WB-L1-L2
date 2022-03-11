package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		if scanner.Scan() {
			line := scanner.Text()

			if line == `\q` {
				break
			}

			commands := strings.Split(line, " | ")

			err := CommandsExec(commands)
			if err != nil {
				log.Println("error to run command: ", err.Error())
			}
		}
	}

	fmt.Println("Bye!")

}

func CommandsExec(commands []string) error {
	for _, command := range commands {
		out, err := exec.Command(command).Output()
		if err != nil {
			return fmt.Errorf("%s, %s", command, err.Error())
		}

		output := string(out[:])
		fmt.Println(output)
	}
	return nil
}
