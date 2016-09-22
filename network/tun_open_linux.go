package network

import (
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

const (
	ifNamSiz = 15
	tunPath  = "/dev/net/tun"
)

type ifReq struct {
	Name  [ifNamSiz]byte
	Flags uint16
}

func openDevice(ifName string) (*os.File, error) {
	file, err := os.OpenFile(tunPath, os.O_RDWR, 0)
	return file, err
}

func createInterface(file *os.File, ifName string) error {
	// Configure interface request
	var req ifReq
	req.Flags = 0
	copy(req.Name[:ifNamSiz], []byte(ifName))
	req.Flags |= uint16(syscall.IFF_TUN)
	req.Flags |= uint16(syscall.IFF_NO_PI)

	// Send interface request
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL, file.Fd(),
		uintptr(syscall.TUNSETIFF),
		uintptr(unsafe.Pointer(&req)))
	if err != 0 {
		return err
	}

	// FIXME: Doesn't seem to work; shelling out for now
	// enableInterface(file, ifName)
	if err := enableInterfaceViaShell(ifName); err != nil {
		return err
	}

	return nil
}

func enableInterface(file *os.File, ifName string) error {

	// Configure interface request
	var req ifReq
	req.Flags = 0
	req.Flags |= uint16(syscall.IFF_UP)
	req.Flags |= uint16(syscall.IFF_RUNNING)
	copy(req.Name[:ifNamSiz], []byte(ifName))

	// Send interface request
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL, file.Fd(),
		uintptr(syscall.SIOCSIFFLAGS),
		uintptr(unsafe.Pointer(&req)))
	if err != 0 {
		return err
	}

	return nil
}

func enableInterfaceViaShell(ifName string) error {
	if err := exec.Command("ip", "link", "set", ifName, "up").Run(); err != nil {
		return err
	}
	if err := exec.Command("ip", "-6", "addr", "add", "fc00::1/7", "dev", ifName).Run(); err != nil {
		return err
	}

	return nil
}
