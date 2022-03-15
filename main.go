package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/VonC/conmode/version"

	"github.com/alecthomas/kong"
	"github.com/ryboe/q"
	"github.com/spewerspew/spew"
)

// Config stores arguments and subcommands
type Config struct {
	Version bool       `help:"if true, print Version and exit." short:"v"`
	Debug   bool       `help:"if true, print debug info" short:"d"`
	Display DisplayCmd `cmd:"" default:"" help:"Display console modes"`
}

type DisplayCmd struct{}

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

	cli := &Config{}
	ctx := kong.Parse(cli)
	// Call the Run() method of the selected parsed command.
	err = ctx.Run(cli)
	ctx.FatalIfErrorf(err)

}

func (dc *DisplayCmd) Run(cli *Config) error {
	if cli.Version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	if cli.Debug {
		spew.Dump(cli)
		q.Q(cli)
	}
	printDefaultConsoleMode()
	return nil
}
