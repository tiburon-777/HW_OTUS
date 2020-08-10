package main

import "os"

func main() {
	// парсим аргументы
	args := os.Args
	env, err := ReadDir(args[1])
	if err != nil {
		os.Exit(111)
	}
	os.Exit(RunCmd(args[2:], env))
}
