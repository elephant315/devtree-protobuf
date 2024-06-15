package main

import (
	"fmt"
	pb "github.com/elephant315/devtree-protobuf/proto"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func main() {
	url := "http://localhost:8080/devices"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching devices:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		var errorMessage pb.Error
		if err := proto.Unmarshal(body, &errorMessage); err != nil {
			fmt.Println("Error unmarshalling error protobuf:", err)
		} else {
			fmt.Println("Error response from server:", errorMessage.Message)
		}
		return
	}

	var deviceList pb.DeviceList
	if err := proto.Unmarshal(body, &deviceList); err != nil {
		fmt.Println("Error unmarshalling protobuf:", err)
		return
	}

	for _, device := range deviceList.Devices {
		fmt.Printf("Device Type: %s, Device Path: %s, Vendor ID: %s, Product ID: %s\n",
			device.DeviceType, device.DevicePath, device.VendorId, device.ProductId)
	}
}
