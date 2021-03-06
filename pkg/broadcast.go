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

//Creates and holds specified directory and refreshes the ttl every 5 seconds
func HoldDir(ec *etcd.Client, directory string, ttl uint64) {
	fmt.Println("Holding Directory: ", directory)
	resp, _ := ec.CreateDir(directory, ttl)
	if resp == nil {
		gou.Info("Directory creation response is nil!")
	}
	gou.Infof("Held %#v \n", resp)

	for {

		ec.UpdateDir(directory, ttl)
		time.Sleep(5 * time.Second)
	}
}

//Post data to a key and refresh set ttl if specified
func PostKey(ec *etcd.Client, key, value string, ttl uint64) {
	//TTL is 0, simply create permanent key
	if ttl == 0 {
		resp, _ := ec.Create(key, value, ttl)
		if resp == nil {
			gou.Warnf("Key[%s] failed to be created\n", key)
		}
		//TTL specified, refresh the ttl every 5 seconds
	} else {
		resp, _ := ec.Create(key, value, ttl)
		if resp == nil {
			gou.Warnf("Key[%s] failed to be created\n", key)
		}
		for {
			ec.Update(key, value, ttl)
			time.Sleep(5 * time.Second)
		}
	}
}
