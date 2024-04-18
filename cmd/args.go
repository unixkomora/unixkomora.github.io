package cmd

import (
	"flag"
	"fmt"
)

func Args() {
	webFlag := flag.Bool("web", false, "start a web server")
	flag.Parse()

	if *webFlag {
		Server()
		return
	} else {
		fmt.Printf("Usage: unixkomora [OPTION]\n\nhelp\tshow a list of commands\nweb\tstart a web server\n")
		return
	}
}
