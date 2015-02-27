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

//Action functions as either a command or response to Actions taken
//  Key is what matches Commands to their responsive webooks
//  Cmd is the execution which takes place
//  Aux functions as the extended data which is checked against or submitted
type Action struct {
	Key string
	Cmd string
	Aux string
}

//Parse single key|command string pairs out into their key and respective command.
//Returns a map of keys to their command
func ParseCmds(cmds CmdSlice) *map[string]*Action {
	cmap := make(map[string]*Action)
	for _, c := range cmds {
		tmp := strings.Split(c, "[[")
		if len(tmp) == 2 {
			tmp2 := strings.Split(tmp[1], "]]")
			if len(tmp2) == 2 {
				cmap[tmp[0]] = &Action{Key: tmp[0], Cmd: tmp2[0], Aux: tmp2[1]}
			} else {
				cmap[tmp[0]] = &Action{Key: tmp[0], Cmd: tmp[1], Aux: ""}
			}
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
