package main

import (
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	c := exec.Command(cmd[0], cmd[1:]...)
	for k, v := range env {
		c.Env = append(os.Environ(), k+"="+v)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		if code, ok := err.(*exec.ExitError); ok {
			return code.ExitCode()
		}

		return -1
	}

	return 0
}
