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

	ip, err := annyong.GetIP()
	if err != nil {
		gou.Infof("Error getting IP: %#v \n", err)
	}

	machines := []string{etcd_host}
	ec := etcd.NewClient(machines)
	fmt.Printf("%#v\n", ec)

	h, _ := os.Hostname()
	gou.Info(h)

	path := fmt.Sprintf("%s/%s", pathStub, h)
	gou.Info(path)

	ec.CreateDir(pathStub, 0)
	go annyong.HoldDir(ec, path, ttl)

	path = fmt.Sprintf("%s/ip", path)
	gou.Debug(path)
	gou.Debug(ip)
	go annyong.PostKey(ec, path, ip, 0)

	//Sleep in loop to let goroutines update etcd
	for {
		time.Sleep(10 * time.Second)
	}

}
