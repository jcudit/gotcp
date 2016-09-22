package network

import "os"

func openDevice(ifName string) (*os.File, error) {
	file, err := os.OpenFile("/dev/"+ifName, os.O_RDWR, 0)
	return file, err
}

func createInterface(file *os.File, ifName string) error {
	return nil
}
