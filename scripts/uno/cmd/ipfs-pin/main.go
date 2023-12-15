package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	err := runCommand("ipfs", "files", "rm", "-r", "--force", "/Pins")
	if err != nil {
		log.Fatal(err)
	}

	err = runCommand("ipfs", "files", "mkdir", "--cid-version", "1", "-p", "/Pins")
	if err != nil {
		log.Fatal(err)
	}

	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		splits := strings.SplitN(line, " ", 2)
		if len(splits) != 2 {
			log.Fatal("must have two splits")
		}

		err := runCommand("ipfs", "pin", "add", "--progress", "--label", splits[1], splits[0])
		if err != nil {
			log.Fatal(err)
		}

		err = runCommand("ipfs", "files", "mkdir", "--cid-version", "1", "-p", "/Pins/"+path.Dir(splits[1]))
		if err != nil {
			log.Fatal(err)
		}

		err = runCommand("ipfs", "files", "cp", splits[0], "/Pins/"+splits[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	err = runCommand("ipfs", "pin", "verify")
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
