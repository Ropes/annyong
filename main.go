package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/araddon/gou"
	"github.com/coreos/go-etcd/etcd"
	"github.com/ropes/annyong/pkg"
)

type cmdSlice []string

func (c *cmdSlice) String() string {
	return fmt.Sprintf("%s", *c)
}
func (c *cmdSlice) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func parseCmds(cmds cmdSlice) *map[string]string {
	cmap := make(map[string]string)
	for _, c := range cmds {
		tmp := strings.Split(c, "|")
		if len(tmp) == 2 {
			cmap[tmp[0]] = tmp[1]
		}
	}
	return &cmap
}

var (
	Info      *log.Logger
	Err       *log.Logger
	etcd_host string
	ttl       uint64
	pathStub  string
	logLevel  string
	cmds      cmdSlice
	cmdMap    map[string]string
)

func main() {
	flag.StringVar(&logLevel, "logLevel", "debug", "Log Level to run[debug,info,warn,error,fatal]")
	flag.StringVar(&etcd_host, "etcd_host", "http://127.0.0.1:4001", "Connection string to the etcd [cluster]")
	flag.StringVar(&pathStub, "pathStub", "/annyong", "Base etcd directory path to use for saving data")
	flag.Var(&cmds, "cmd", "List of cmd$command-string pairs to potentially run")
	flag.Uint64Var(&ttl, "ttl", 20, "TTL of directories created")

	flag.Parse()
	gou.SetupLogging(logLevel)

	gou.Debugf("%#v\n", cmds)
	gou.Debugf("%#v\n", parseCmds(cmds))

	//Connect to etcd
	machines := []string{etcd_host}
	ec := etcd.NewClient(machines)

	//Create base stub path
	ec.CreateDir(pathStub, 0)

	//Discover Base Host information
	ip, err := annyong.GetIP()
	if err != nil {
		gou.Warnf("Error getting IP: %#v \n", err)
	}
	h, _ := os.Hostname()
	gou.Debugf("%s: %s\n", h, ip)

	path := fmt.Sprintf("%s/%s", pathStub, h)
	go annyong.HoldDir(ec, path, ttl)

	path = fmt.Sprintf("%s/ip", path)
	go annyong.PostKey(ec, path, ip, ttl)

	//Sleep in loop to let goroutines update etcd
	for {
		time.Sleep(10 * time.Second)
	}

}
