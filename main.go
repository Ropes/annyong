package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/coreos/go-etcd/etcd"
	"github.com/ropes/annyong/Annyong"
)

var (
	Info      *log.Logger
	Err       *log.Logger
	etcd_host string
	ttl       uint64
)

func findIp() {
	ifaddrs, _ := net.InterfaceAddrs()

	for _, a := range ifaddrs {
		fmt.Printf("%#v %#v \n", a.Network(), a.String())
	}
}

func main() {
	flag.StringVar(&etcd_host, "etcd_host", "http://127.0.0.1:4001", "Connection string to the etcd [cluster]")
	flag.Uint64Var(&ttl, "TTL", 30, "TTL of directories created")

	flag.Parse()
	Info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info.Print("annyong!\n")

	ip, err := annyong.GetIP()
	if err != nil {
		Info.Printf("Error getting IP: %#v \n", err)
	}
	fmt.Printf("%s \n ", ip)

	machines := []string{etcd_host}
	ec := etcd.NewClient(machines)
	fmt.Printf("%#v\n", ec)

	h, _ := os.Hostname()
	Info.Print(h)
	path := fmt.Sprintf("/annyong/%s", h)
	Info.Print(path)

	resp, _ := ec.Get("hihi", false, false)
	if resp == nil {
		Info.Print("response is nil!")
	}
	Info.Printf("get: %#v \n", resp)

	resp, _ = ec.CreateDir(path, ttl)
	if resp == nil {
		Info.Print("response is nil!")
	}
	Info.Printf("%#v \n", resp)

	/*
		  path := fmt.Sprintf("/annyong/%s", h)
			resp, _ = ec.Create("hostname", h, path, ttl)
			if resp == nil {
				Info.Print("response is nil!")
			}
			Info.Printf("%#v \n", resp)
	*/

}
