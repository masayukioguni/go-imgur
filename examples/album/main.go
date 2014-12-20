package main

import (
	"fmt"
	"github.com/masayukioguni/go-imgur/imgur"
	"os"
)

var applicationName = "go-imgr"
var version string

func main() {
	id := os.Getenv("IMGUR_API_CLIENT_ID")
	client, _ := imgur.NewClient(&imgur.Option{ClientID: id})

	data, err := client.GalleryService.GetAlbum()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	fmt.Printf("%v\n", data)

}
