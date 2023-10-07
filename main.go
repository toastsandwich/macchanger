package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func exe(cmd string, args []string) {
	cmdexe := exec.Command(cmd, args...)
	cmdexe.Stdout = os.Stdout
	cmdexe.Stderr = os.Stderr
	err := cmdexe.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	iface := flag.String("iface", "", "Wireless interface to change MAC address")
	mac := flag.String("mac", "", "Set new MAC address")
	flag.Parse()

	if *iface == "" || *mac == "" {
		println("Usage: ./macchanger -iface <WirelessInterfaceName> -mac <NewMACAddress>")
		return
	}

	// Bring down the wireless interface
	exe("sudo", []string{"ip", "link", "set", *iface, "down"})

	// Set the new MAC address
	exe("sudo", []string{"ifconfig", *iface, "hw", "ether", *mac})

	// Bring the wireless interface up again
	exe("sudo", []string{"ip", "link", "set", *iface, "up"})

	log.Println("MAC address for wireless interface", *iface, "changed to", *mac)
}
