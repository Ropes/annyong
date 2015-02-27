package annyong

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/araddon/gou"
)

type CmdSlice []string

func (c *CmdSlice) String() string {
	return fmt.Sprintf("%s", *c)
}
func (c *CmdSlice) Set(value string) error {
	*c = append(*c, value)
	return nil
}

//Parse single key|command string pairs out into their key and respective command.
//Returns a map of keys to their command
func ParseCmds(cmds CmdSlice) *map[string]string {
	cmap := make(map[string]string)
	for _, c := range cmds {
		tmp := strings.Split(c, "|")
		if len(tmp) == 2 {
			cmap[tmp[0]] = tmp[1]
		}
	}
	return &cmap
}

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
