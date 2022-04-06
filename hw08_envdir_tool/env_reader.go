package main

import (
	"io/ioutil"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		filename := file.Name()
		if strings.ContainsRune(filename, '=') {
			// the name S must not contain =
			continue
		}
		str, _ := ioutil.ReadFile(dir + "/" + filename)
		env[filename] = getEnvValue(str)
	}
	return env, nil
}

func getEnvValue(b []byte) EnvValue {
	split := strings.Split(string(b), "\n")
	str := strings.TrimRight(split[0], "\t ")
	str = strings.ReplaceAll(str, "\x00", "\n")
	return EnvValue{
		str,
		len(str) == 0,
	}
}
