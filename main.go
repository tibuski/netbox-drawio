package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const netboxAPI = "http://localhost:9000/api" // Replace with the URL of your Netbox instance

// Device represents a device in Netbox
type Device struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	DeviceRole struct {
		ID int `json:"id"`
	} `json:"device_role"`
}

type response struct {
	Results []Device `json:"results"`
}

func main() {
	// Make a GET request to the /dcim/devices/ endpoint of the Netbox API
	resp, err := http.Get(netboxAPI + "/dcim/devices/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Decode the response into a response object
	var r response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		fmt.Println(err)
		return
	}

	// Group the devices by device role
	deviceRoles := make(map[int][]Device) // Map from device role ID to slice of devices
	for _, device := range r.Results {
		deviceRoles[device.DeviceRole.ID] = append(deviceRoles[device.DeviceRole.ID], device)
	}

	// Print the devices grouped by device role
	fmt.Println("Devices grouped by device role:")
	for roleID, devices := range deviceRoles {
		fmt.Printf("- Device role %d:\n", roleID)
		for _, device := range devices {
			fmt.Printf("  - %d: %s\n", device.ID, device.Name)
		}
	}
}
