package main

import (
	"fmt"
	"vighnesh.org/virustotal"
)

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")
	response, e := virustotalapi.Comment("Resource", "Good File")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("Response Code", response.ResponseCode)
	fmt.Println("Message", response.Message)
}
