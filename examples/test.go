package main

import (
	"fmt"
	"github.com/masayukioguni/go-imgur/imgur"
	"os"
)

var applicationName = "go-imgr"
var version string

func Account(client *imgur.Client) {

	data, _, err := client.AccountService.Account("masayukioguni")

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	fmt.Printf("%v\n", data)

}

func Album(client *imgur.Client) {

	data, _, err := client.AlbumService.GetAlbum("lDRB2")

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	fmt.Printf("%v\n", data)

}

func main() {
	id := os.Getenv("IMGUR_API_CLIENT_ID")
	client, _ := imgur.NewClient(&imgur.Option{ClientID: id})

	data, _, err := client.GalleryService.GetAlbum()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	fmt.Printf("%v\n", data)

	Account(client)
	Album(client)

}
