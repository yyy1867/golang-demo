package config

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
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Host string
	Port int
}

type RedisConfig struct {
	Addrs    []string
	Password string
}

type Config struct {
	Zookeeper ZookeeperConfig
	Openstack OpenstackConfig
	Server    ServerConfig
	Mysql     MySqlConfig
	Redis     RedisConfig
}
