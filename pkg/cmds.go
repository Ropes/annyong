package annyong

import (
	"bytes"
	"os/exec"
)

//Using *nix system hostname command; find the host's broadcast IP
func GetIP() (string, error) {
	cmd := exec.Command("hostname", "-i")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func Cmd(cmd string) (interface{}, error) {
	return nil, nil
}

func HttpQuery(resource string) (interface{}, error) {
	return nil, nil
}
