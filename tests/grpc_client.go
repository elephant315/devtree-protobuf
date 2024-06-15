package main

import (
	"context"
	"fmt"
	"log"
	"github.com/elephant315/devtree-protobuf/proto" // Replace with your actual module name

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewDeviceServiceClient(conn)

	req := &proto.GetDevicesRequest{}
	res, err := c.GetDevices(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get data: %v", err)
	}

	for _, device := range res.Devices {
		fmt.Printf("Device Type: %s, Device Path: %s, Vendor ID: %s, Product ID: %s\n",
			device.DeviceType, device.DevicePath, device.VendorId, device.ProductId)
	}
}
