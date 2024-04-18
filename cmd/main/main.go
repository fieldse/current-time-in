// Import and run the cityLookup package for the city passed on the command line
package main

import "github.com/fieldse/current-time-in/pkg/citylookup"

func main() {
	var args []string // fixme: get from sys.argv
	citylookup.RunCLI(args)
}
