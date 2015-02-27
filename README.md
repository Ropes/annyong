Annyong!
-------

Annyong!  Minimal service discovery service which broadcasts determines information about itself and broadcasts to etcd.
(Designed with Google Compute Engine, Saltstack, and etcd in mind).

Prototyped at Lytics hackathon... so sleep depravity will apply.

##Design
* Binary is loaded onto machine configured with minimal setup information to connect to etcd.
  * Configuration holds ip:port to etcd cluster
  * Binary will gather information via static system lookups or by using resource lookups defined optionally in the config
  * Writes the data to etcd
    * Constantly refreshes it with a short TTL so that if nodes fail and disappear the data removes itself from etcd
  * Logs information on lookups to stdout/err


##CLI
* MVP
`annyong -etcd_host localhost:4001`
  * Post ip and hostname to etcd


###Would be cool
* Expected values  
`annyong -pinghttp localhost:9200/endpoint$expected_value`
`annyong -pingexec "some bash command and"$expectedvalue`

* etcd connection  
`annyong -etcd 10.240.10.10:4001`

* Specify values discovered from environment  
`annyong -tag vmname$cmdtag -cmd "cmdtag$curl localhost"`
  * specify desired values linked to a comamnd tag, tag links to a command which is run
  * tag

##TODO
* etcd client/connections
  * TTL system
* Configuration
  * CLI flags
  * Resource paths 
* Gather system information
  * Check config for resource path
  * Default to pull metrics
* Listen to system information
  * Query http endpoints(elasticsearch)
  * Specify expected values

