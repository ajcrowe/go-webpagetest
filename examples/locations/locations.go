package main

import (
	"fmt"
	"github.com/ajcrowe/go-webpagetest"
)

func main() {
	config := wpt.NewDefaultConfig()

	client, _ := wpt.NewClient(config)

	locations, err := client.Locations()
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range locations {
		fmt.Printf("location:\t%s\nbrowser:\t%s\nlabel:\t\t%s\n", location.Location, location.Browser, location.Label)
		fmt.Println("-------------------------------------------------------------------")
	}

	locDefault := locations.GetDefault()

	fmt.Println("\n\n\t\tDefault Location")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("location:\t%s\nbrowser:\t%s\nlabel:\t\t%s\n", locDefault.Location, locDefault.Browser, locDefault.Label)
	fmt.Println("-------------------------------------------------------------------")
}
