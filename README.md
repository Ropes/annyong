Annyong!
-------

Annyong!  Minimal service discovery service which introspects information and broadcasts to etcd. 

(Designed with Google Compute Engine, Saltstack, and etcd in mind).

Prototyped at Lytics hackathon... so sleep depravity will apply.

Design
------
* Binary is loaded onto machine configured with minimal setup information to connect to etcd.
  * Configuration holds ip:port to etcd cluster
  * Binary can gather information via commands defined optional cli flags
  * Writes the data to etcd
    * Constantly refreshes it with a short TTL so that if nodes fail and disappear the data removes itself from etcd
* Logs information on lookups to stdout/err

* TODO: Output from commands write to PubSub channel which Hooks listen to.
  * Hooks match the key and the output value, if they match then their action triggers.


##CLI
* MVP(done)
`annyong -etcd_host http://localhost:4001 -logLevel [debug...fatal] -pathStub /annyong -ttl 30`
  * Post hostname and ip value to etcd
    * `/annyong/hostname/ip="...."`

##Examples
* Most basic  
```
annyong -etcd_host http://localhost:4001
```

####Specifying commands and hooks
`./annyong -etcd_host http://localhost:4001 -ttl 7 -cmd "hihi[[curl es:9200/_cluster/health ]]nil"  -hook "hihi[[curl miyagi]]red"
`

![annyong](http://38.media.tumblr.com/tumblr_lgf6pdu8gw1qfrk53o1_500.gif)  

###TODO

* Commands and hooks
`annyong ... -cmd hihi[[localhost:9200/endpoint]]expected_value`
`annyong ... -hook key[[some bash command and"]]expectedvalue`

* Specify values discovered from environment [grains for now]  
`annyong ... -grain key[[cmd-to-run]]keyPath`  
  * specify desired values linked to a comamnd tag, tag links to a command which is run

* Grains Webhooks
  * Tricky part is configuring it with the cli
  * Match




##TODO
* Configuration
  * CLI flags
  * etcd resource paths 
* Gather system information
  * Check config for resource path
  * Default to pull metrics
* Listen to system information
  * Query http endpoints(elasticsearch)
  * Specify expected values


