// Import and run the cityLookup package for the city passed on the command line
package main

import (
	"github.com/fieldse/current-time-in/pkg/citylookup"
)

var Logger = citylookup.Logger

func main() {
	var args []string // fixme: get from sys.argv
	err := citylookup.RunCLI(args)
	if err != nil {
		Logger.Fatal().Err(err).Msgf("CLI failed")
	}
}
