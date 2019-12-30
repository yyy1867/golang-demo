package common

type ZookeeperConfig struct {
	Addrs []string
}

type OpenstackConfig struct {
	Url      string
	Username string
	Password string
	Domain   string
}

type Config struct {
	Zookeeper ZookeeperConfig
	Openstack OpenstackConfig
}

var DbConfig Config
