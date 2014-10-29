package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/homburg/mdns"
	"log"
)

func main() {
	usage := `Publish mDNS hostname.

	Create an mdns server, listening on UDP multicast, port 5353 for hostname queries.

	Usage:
	  publish-mdns-hostname <hostname> <ip>


	`
	arguments, err := docopt.Parse(usage, nil, true, "publish-mdns-hostname 1.0", false)

	if nil != err {
		log.Fatal(err)
	}

	record := fmt.Sprintf("%s IN A %s", arguments["<hostname>"], arguments["<ip>"])
	fmt.Printf("Publishing record: \"%s\"\n", record)

	mdns.Publish(record)

	fmt.Println("Serving...")
	select {}
}
