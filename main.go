package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/VonC/conmode/version"
	"github.com/erikgeiser/coninput"

	"github.com/alecthomas/kong"
	"github.com/ryboe/q"
	"github.com/spewerspew/spew"
)

//go:embed version/*
var versionFs embed.FS

// Config stores arguments and subcommands
// From https://github.com/alecthomas/kong/issues/51
type Config struct {
	Version bool `help:"if true, print Version and exit." short:"v"`
	Debug   bool `help:"if true, print debug info" short:"d"`
	Mode    struct {
		Mode    string     `arg:"" default:"" optional:"" help:"Input console Mode code to decrypt"`
		Display DisplayCmd `cmd:"" default:"" help:"Display console modes"`
	} `arg`
}

type DisplayCmd struct{}

func fatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: error '%+v'", msg, err)
	}
}

// myproject main entry
func main() {

	version.VersionFS = versionFs
	var err error

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fatal("Unable to find current program execution directory", err)
	log.Println(dir)

	cli := &Config{}
	display := false
	for _, a := range os.Args {
		if a == "display" {
			display = true
			break
		}
	}
	if !display {
		os.Args = append(os.Args, "display")
	}
	//spew.Dump(os.Args)
	ctx := kong.Parse(cli)
	//spew.Dump(cli)
	//spew.Dump(ctx)
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
	fmt.Println(os.Args[0])
	if cli.Mode.Mode == "" || cli.Mode.Mode == "display" {
		printDefaultConsoleMode()
	} else {
		conMode, err := strconv.ParseUint(cli.Mode.Mode, 10, 32)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("conmode %d: '%s'\n", conMode, coninput.DescribeInputMode(uint32(conMode)))
	}
	return nil
}
