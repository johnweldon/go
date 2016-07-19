package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"launchpad.net/goyaml"

	"github.com/johnweldon/go-misc/util/juju"
)

var defaultVar = "user"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		defaultVar = args[0]
	}

	curEnv := getCurrentEnvironment()
	settings, err := getEnvironmentSettings(curEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read environment %q\n", err)
		return
	}
	switch defaultVar {
	case "user":
		fmt.Fprintf(os.Stdout, "%s\n", settings.User)
	case "pass":
		fmt.Fprintf(os.Stdout, "%s\n", settings.Password)
	case "state":
		fmt.Fprintf(os.Stdout, "%v\n", settings.StateServers)
	default:
		fmt.Fprintf(os.Stderr, "unknown key\n")
	}
}

func getCurrentEnvironment() string {
	home := os.Getenv("HOME")
	file := path.Join(home, ".juju", "current-environment")
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return "local"
	}
	return string(raw)
}

func getEnvironmentSettings(name string) (settings juju.EnvironInfoData, err error) {
	raw, err := readEnvironment(name)
	if err != nil {
		return
	}
	err = goyaml.Unmarshal(raw, &settings)
	return
}

func readEnvironment(name string) ([]byte, error) {
	home := os.Getenv("HOME")
	file := path.Join(home, ".juju", "environments", name+".jenv")
	return ioutil.ReadFile(file)
}
