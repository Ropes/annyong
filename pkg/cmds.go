package annyong

import (
	"bytes"
	"net"
	"os/exec"
	"strings"

	"github.com/araddon/gou"
)

//Using *nix system hostname command; find the host's broadcast IP
//Returns the IP as a string
func GetIP() (string, error) {
	cmd := exec.Command("hostname", "-i")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func FindInterfaces() {
	ifaddrs, _ := net.InterfaceAddrs()

	for _, a := range ifaddrs {
		gou.Debugf("%#v %#v \n", a.Network(), a.String())
	}
}

func Cmd(cmd string) (interface{}, error) {
	return nil, nil
}

func HttpQuery(resource string) (interface{}, error) {
	return nil, nil
}
