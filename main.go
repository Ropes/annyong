package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/araddon/gou"
	"github.com/coreos/go-etcd/etcd"
	"github.com/ropes/annyong/pkg"
)

var (
	Info      *log.Logger
	Err       *log.Logger
	etcd_host string
	ttl       uint64
	pathStub  string
	logLevel  string
)

func main() {
	flag.StringVar(&logLevel, "logLevel", "debug", "Log Level to run[debug,info,warn,error,fatal]")
	flag.StringVar(&etcd_host, "etcd_host", "http://127.0.0.1:4001", "Connection string to the etcd [cluster]")
	flag.StringVar(&pathStub, "pathStub", "/annyong", "Base etcd directory path to use for saving data")
	flag.Uint64Var(&ttl, "ttl", 20, "TTL of directories created")

	flag.Parse()
	gou.SetupLogging(logLevel)

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
	go annyong.PostKey(ec, path, ip, 0)

	//Sleep in loop to let goroutines update etcd
	for {
		time.Sleep(10 * time.Second)
	}

}
