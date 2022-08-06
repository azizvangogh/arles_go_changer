package main

import (
	"fmt"

	"github.com/millefalcon/go-subprocess/subprocess"
)

func main() {
	// İnput Girişleri

	fmt.Print("Enter your interface: ")
	var Iface string
	fmt.Scanln(&Iface)
	fmt.Print("Enter your mac: ")
	var mac string
	fmt.Scanln(&mac)

	subprocess.Popen("ifconfig", Iface, "down")
	subprocess.Popen("ifconfig", Iface, "hw", "ether", mac)
	subprocess.Popen("ifconfig", Iface, "up")
	fmt.Println("Mac Adress Changed")

	if mac == "" && Iface == "" {
		fmt.Println("Usage: macchanger -i <interface> -m <mac>")
	}
}
