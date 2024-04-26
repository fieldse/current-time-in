// Import and run the cityLookup package for the city passed on the command line
package main

import (
	"github.com/fieldse/current-time-in/pkg/cli"
	"github.com/rs/zerolog/log"
)

func main() {
	var args []string // fixme: get from sys.argv
	err := cli.RunCLI(args)
	if err != nil {
		log.Fatal().Msg("CLI failed")
	}
}
