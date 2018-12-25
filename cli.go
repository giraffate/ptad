package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	EnvPaperTrailAPIToken = "PAPERTRAIL_API_TOKEN"
	EnvDebug              = "PTAD_DEBUG"

	TimeFormat = "2006-01-02-15"
)

// CLI is the command line object.
type CLI struct {
	outStream, errStream io.Writer
}

// Debugf prints debug output.
func Debugf(format string, args ...interface{}) {
	if env := os.Getenv(EnvDebug); len(env) > 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// Run invokes the CLI with given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		num          int
		debug, local bool
	)

	flags := flag.NewFlagSet("ptad", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.IntVar(&num, "num", 1, "")
	flags.IntVar(&num, "n", 1, "")

	flags.BoolVar(&debug, "debug", false, "")
	flags.BoolVar(&debug, "d", false, "")

	flags.BoolVar(&local, "local", false, "")
	flags.BoolVar(&local, "l", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		log.Printf("[ERROR] %v", err)
		return 1
	}

	if debug {
		os.Setenv(EnvDebug, "1")
		Debugf("%s", "Run as debug mode")
		Debugf("num: %d", num)
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) < 1 {
		log.Println("[ERROR] No arguments found")
		return 1
	}

	t, err := time.ParseInLocation(TimeFormat, parsedArgs[0], time.Local)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return 1
	}
	if local {
		t = t.UTC()
	}

	client := NewPaperTrailClient(os.Getenv(EnvPaperTrailAPIToken))
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			formattedTime := t.Add(time.Duration(i) * time.Hour).Format("2006-01-02-15")
			if err := client.DownloadArchive(formattedTime); err != nil {
				log.Printf("[ERROR] %v", err)
				return
			}
			Debugf("Completed: %s", formattedTime)
		}(i)
	}
	wg.Wait()

	return 0
}
