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

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Zookeeper ZookeeperConfig
	Openstack OpenstackConfig
	Server    ServerConfig
}

var DbConfig Config
