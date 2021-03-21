package main

import (
	"log"
	"main/client"
)

func main() {
	// setup clients
	secureCLI := client.Setup(true)
	insecureCLI := client.Setup(false)
	secureURL := "https://www.golang.org"
	insecureURL := "https://expired.badssl.com"

	// skip tls verification = false
	if err := client.GetStatusCode(secureCLI, secureURL); err != nil {
		log.Println(err)
	}

	// skip tls verification = false
	if err := client.GetStatusCode(secureCLI, insecureURL); err != nil {
		log.Println(err)
	}

	// skip tls verification = true
	if err := client.GetStatusCode(insecureCLI, insecureURL); err != nil {
		log.Println(err)
	}

	// skip tls verification = false
	c := client.Controller{Client: secureCLI}
	if err := c.Get(insecureURL); err != nil {
		log.Println(err)
	}

}

/*
HTTP Status https://www.golang.org: 200
2021/03/21 09:53:13 Get "https://expired.badssl.com": x509: certificate has expired or is not yet valid: current time 2021-03-21T09:53:13+02:00 is after 2015-04-12T23:59:59Z
HTTP Status https://expired.badssl.com: 200
2021/03/21 09:53:14 Get "https://expired.badssl.com": x509: certificate has expired or is not yet valid: current time 2021-03-21T09:53:14+02:00 is after 2015-04-12T23:59:59Z
*/
