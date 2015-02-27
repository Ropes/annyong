package annyong

import (
	"fmt"
	"time"

	"github.com/araddon/gou"
	"github.com/coreos/go-etcd/etcd"
)

func FormNode() (*etcd.Node, error) {
	return nil, nil
}

//Creates and holds
func HoldDir(ec *etcd.Client, directory string, ttl uint64) {
	fmt.Println("Holding Directory: ", directory)
	resp, _ := ec.CreateDir(directory, ttl)
	if resp == nil {
		gou.Info("Directory creation response is nil!")
	}
	gou.Infof("Held %#v \n", resp)

	for {

		ec.UpdateDir(directory, ttl)
		time.Sleep(10 * time.Second)
	}
}

//Post data to a key and continually update the ttl
func PostKey(ec *etcd.Client, key, value string, ttl uint64) {
	resp, _ := ec.Create(key, value, ttl)
	if resp == nil {
		gou.Warnf("Key[%s] failed to be created\n", key)
	}

}
