package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range servers {
		fmt.Println(ip.Host)
	}
}
