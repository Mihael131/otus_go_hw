package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	name := cmd[0]
	command := exec.Command(name, cmd[1:]...)
	command.Env = getEnv(env)
	command.Stdout = os.Stdout
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}
	returnCode = command.ProcessState.ExitCode()
	return
}

func getEnv(env Environment) []string {
	if env == nil {
		return os.Environ()
	}
	for n, v := range env {
		if v.NeedRemove {
			os.Unsetenv(n)
		} else {
			os.Setenv(n, v.Value)
		}
	}
	return os.Environ()
}
