package common

type ZookeeperConfig struct {
	Addrs        []string
	MaxConnCount int
}

type OpenstackConfig struct {
	Url      string
	Username string
	Password string
	Domain   string
}

type MySqlConfig struct {
	Url      string
	Username string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Zookeeper ZookeeperConfig
	Openstack OpenstackConfig
	Server    ServerConfig
	Mysql     MySqlConfig
}

var DbConfig Config
