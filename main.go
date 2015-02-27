package main

import (
	"flag"
	"fmt"
	"log"
	"net"
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

func findIp() {
	ifaddrs, _ := net.InterfaceAddrs()

	for _, a := range ifaddrs {
		fmt.Printf("%#v %#v \n", a.Network(), a.String())
	}
}

func main() {
	flag.StringVar(&logLevel, "logLevel", "debug", "Log Level to run[debug,info,warn,error,fatal]")
	flag.StringVar(&etcd_host, "etcd_host", "http://127.0.0.1:4001", "Connection string to the etcd [cluster]")
	flag.Uint64Var(&ttl, "TTL", 20, "TTL of directories created")
	pathStub = "/annyong"

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

	path = fmt.Sprintf("%s/ip", path, ip)
	gou.Info(path)
	go annyong.PostKey(ec, path, ip, 0)

	/*
		resp, _ := ec.Get("hihi", false, false)
		if resp == nil {
			Info.Print("response is nil!")
		}
		Info.Printf("get: %#v \n", resp)

			  path := fmt.Sprintf("/annyong/%s", h)
				resp, _ = ec.Create("hostname", h, path, ttl)
				if resp == nil {
					Info.Print("response is nil!")
				}
				Info.Printf("%#v \n", resp)
	*/
	for {
		time.Sleep(10 * time.Second)
	}

}
