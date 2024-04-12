# Current Time In...


## Summary

A Go utility to get the current time in a given city by name. Uses the WorldTimeAPI service.



## Usage

```
$ current-time-in New York

The current time is New York City is 11:37 AM, Saturday April 12, 2024
```


## Installation

You can clone this repo, and run the file directly from its current location with Go, or build and copy the binary to a directory in your system path. 


Running directly:
```
go main.go
```

Build & copy to system path
```
# Build
go build main.go -o dist/current-time-in 

# Copy to system path
cp dist/current-time-in /usr/local/bin/

# Run it from system path
current-time-in New York
```


## Requirements

- Go v1.21+
Should run on Mac, Windows, or Linux.


## Maintainer
Matt Fields - hello@mattfields.dev

## Project homepage
https://github.com/fieldse/current-time-in

## Credits
- WorldTimeAPI - https://worldtimeapi.org/

## License
- MIT License
