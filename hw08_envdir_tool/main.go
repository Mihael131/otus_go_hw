package main

import (
	"log"
	"os"
)

func main() {
	var dir string
	var cmd []string

	cnt := len(os.Args)
	switch {
	case cnt == 1:
		log.Fatal("Dir not set")
	case cnt == 2:
		log.Fatal("Command not set")
	case cnt > 2:
		dir = os.Args[1]
		cmd = os.Args[2:]
	}

	env, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	RunCmd(cmd, env)
}
