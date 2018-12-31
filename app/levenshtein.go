package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/agnivade/levenshtein"
)

// DNSTwistResult is the struct that holds the output from the DNSTwist command.
type DNSTwistResult struct {
	DNSa       string `json:"dns-a"`
	DNSmx      string `json:"dns-mx"`
	DNSns      string `json:"dns-ns"`
	DomainName string `json:"domain-name"`
	Fuzzer     string `json:"fuzzer"`
}

// GoLevenshtein function finds the smallest Levenshtein distance from the given website and those from DNSTwist and returns the percentage and the website's URL.
func GoLevenshtein(siteName string) (float32, string) {
	outDNS, err := exec.Command("dnstwist", "-j", "-r", siteName).Output()
	fmt.Println("Resultado do comando:", string(outDNS))
	if err != nil {
		fmt.Println("Error while executing command")
		log.Fatal(err)
	}

	var d []DNSTwistResult
	err = json.Unmarshal(outDNS, &d)
	if err != nil {
		fmt.Println("Error during unmarshall")
		log.Fatal(err)
	}

	outWget, err := exec.Command("wget", "-O-", "--waitretry=0", "--tries=5", siteName).Output()
	if err != nil {
		fmt.Println("Error during main site wget command")
		log.Fatal(err)
	}
	siteCode := string(outWget)

	var smallestPercent float32
	var smallestPercentWebsite string
	var percent float32
	smallestPercent = 100
	for _, result := range d {
		if result.DomainName == "globo.com" {
			continue
		}
		outWgetInstance, err := exec.Command("wget", "-O-", "--waitretry=0", "--tries=5", result.DomainName).Output()
		if err != nil {
			fmt.Println("Error during wget from one of the dnstwist websites")
			// log.Fatal(err)
		}
		siteCodeInstance := string(outWgetInstance)
		distance := levenshtein.ComputeDistance(siteCodeInstance, siteCode)
		siteCodeSize := len(siteCode)
		percent = 100 * (float32(distance) / float32(siteCodeSize))
		fmt.Println("\n Website: ", result.DomainName)
		fmt.Println("O percentual é de: ", percent)

		if percent <= smallestPercent {
			smallestPercent = percent
			smallestPercentWebsite = result.DomainName
		}
	}

	return smallestPercent, smallestPercentWebsite
}
