// Package commonwealth is all of the common stuff we use
package commonwealth

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

// some well known data defines
const (
	identificationPath = "/var/cache/afm/identifier.id"
	defaultAppID       = "id_failure"
)

func determineCurrentIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("unable to find ip address")
}

func determineCurrentNetworkHardwareInterface(currentIP string) (string, error) {
	// get all the system's or local machine's network interfaces
	interfaces, interfaceerr := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				// only interested in the name with current IP address
				if strings.Contains(addr.String(), currentIP) {
					return interf.Name, nil
				}
			}
		}
	}
	return "", interfaceerr
}

func determineDeviceMACAddress() (string, error) {

	currentIP, err := determineCurrentIP()
	if err != nil {
		return "", err
	}

	hardwareInterfaceName, err := determineCurrentNetworkHardwareInterface(currentIP)

	if err != nil {
		return "", err
	}

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(hardwareInterfaceName)

	if err != nil {
		return "", err
	}

	macAddress := netInterface.HardwareAddr

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		return "", err
	}

	return hwAddr.String(), nil
}

func determineDeviceSerialNumber() (string, error) {
	out, err := exec.Command("/usr/sbin/dmidecode", "-t", "system").Output()
	if err == nil {
		return "", err
	}
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "Serial Number") {
			s := strings.Split(l, ":")
			return s[len(s)-1], nil
		}
	}
	return "", fmt.Errorf("unable to find serial number")
}

func DetermineDeviceClientID() string {
	// Try reading well known file location
	fileData, err := os.ReadFile(identificationPath)
	if err == nil {
		identifier := string(fileData)
		cleansedIdentifier := strings.TrimRight(identifier, "\n")
		return cleansedIdentifier
	}

	selectedID := defaultAppID
	// Try reading hardware assigned serial number
	serialNumber, err := determineDeviceSerialNumber()
	if err == nil {
		selectedID = serialNumber
	}

	// Default back to mac address
	macAddress, err := determineDeviceMACAddress()
	if err == nil {
		selectedID = macAddress
	}

	if selectedID != defaultAppID {
		file, err := os.Create(identificationPath)
		if err == nil {
			file.WriteString(selectedID)
			file.Close()
		}
	}

	return selectedID
}
