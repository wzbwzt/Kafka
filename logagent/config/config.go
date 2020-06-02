package config

//整个配置文件
type AppConf struct {
	KafkaConf `ini:"kafka"`
	TailConf `ini:"taillog"`
}


//Kafka 配置文件结构体
type KafkaConf struct {
	Address string `ini:"addrs"`
	Topic string`ini:"topic"`
}

//Tail 配置文件结构体
type TailConf struct {
	Path string `ini:"path"`
}

