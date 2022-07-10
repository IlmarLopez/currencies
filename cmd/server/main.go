package main

import "flag"

// Version is the version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the configuration file")

func main() {
	flag.Parse()

}
