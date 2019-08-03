package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/grantae/certinfo"
)

var f = flag.String("f", "", "file")

func main() {
	flag.Parse()

	if *f == "" {
		log.Fatal("file must set!")
		return
	}

	// Read and parse the PEM certificate file
	pemData, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		log.Fatal("Certificate decoding error")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// Print the certificate
	result, err := certinfo.CertificateText(cert)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
