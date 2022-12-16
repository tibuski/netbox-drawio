package main

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	NETBOX_URL      = "https://demo.netbox.dev/api"
	NETBOX_USER     = "admin"
	NETBOX_PASSWORD = "admin"
)

func main() {

	// Create a token
	token := CreateToken()

	// Get devices
	GetDevices(token)

}

func CreateToken() string {

	var jsonData = []byte(`{
		"username": "admin",
		"password": "admin"
	}`)

	// Create a new token in netbox
	// Create a new request to create a token
	request, err := http.NewRequest("POST", NETBOX_URL+"/users/tokens/provision/", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	// get key from response body and put it in token variable

	return token

}

func GetDevices(token string) {

	// New GET request with Authorization header and token
	req, err := http.NewRequest("GET", NETBOX_URL+"/dcim/devices/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Token "+token)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	println(resp.Status)

}
