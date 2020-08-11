package main

import (
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	c := exec.Command(cmd[0], cmd[1:]...)
	for k, v := range env {
		_, ok := os.LookupEnv(k)
		if ok {
			if err := os.Unsetenv(k); err != nil {
				return -1
			}
		}
		if v != "" {
			if err := os.Setenv(k, v); err != nil {
				return -1
			}
		}
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
