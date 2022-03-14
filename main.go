package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/VonC/conmode/version"

	"github.com/jpillora/opts"
	"github.com/ryboe/q"
	"github.com/spewerspew/spew"
)

// Config stores arguments and subcommands
type Config struct {
	Arg     string `help:"a string argument"`
	Version bool   `help:"if true, print Version and exit."`
	Debug   bool   `help:"if true, print debug info"`
}

var c = &Config{}

func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: error '%+v'", msg, err)
	}
}

// myproject main entry
func main() {

	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fatal("Unable to find current program execution directory", err)
	log.Println(dir)
	opts.New(c).
		ConfigPath(filepath.Join(dir, "conf.json")).
		Parse()
	if c.Version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	if c.Debug {
		spew.Dump(c)
		q.Q(c)
	}
	fmt.Println(os.Args[0])
	printDefaultConsoleMode()
}
