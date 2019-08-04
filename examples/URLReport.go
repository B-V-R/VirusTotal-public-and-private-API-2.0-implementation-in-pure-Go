package main

import (
	"fmt"
	"vighnesh.org/virustotal"
)

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")
	response, e := virustotalapi.URLReport("vighnesh.org")
	if e != nil {
		fmt.Println(e)
	}

	for engine, report := range response.Scans {
		fmt.Println("Scan Engine", engine)
		if report.Detected {
			fmt.Println("Version", report.Version)
			fmt.Println("Updated", report.Update)
			fmt.Println("Malware", report.Malware)
		} else {
			fmt.Println("No Malware Detected")
		}

	}
}
