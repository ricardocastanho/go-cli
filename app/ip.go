package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func findIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
