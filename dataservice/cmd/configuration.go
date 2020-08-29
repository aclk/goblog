package cmd

type Config struct {
	Environment        string `arg:"env:ENVIRONMENT"`
	CockroachdbConnUrl string `args:"env:COCKROACHDB_CONN_URL"`
	ZipkinServerUrl    string `arg:"env:ZIPKIN_SERVER_URL"`
	ConfigEventBus     string `arg:"env:CONFIG_EVENT_BUS"`
	ServerConfig
	AmqpConfig
}

type ServerConfig struct {
	Port string `arg:"env:SERVER_PORT"`
	Name string `arg:"env:SERVICE_NAME"`
}

type AmqpConfig struct {
	ServerUrl string `arg:"env:AMQP_SERVER_URL"`
}

func DefaultConfiguration() *Config {
	return &Config{
		Environment: "dev",
		CockroachdbConnUrl: "postgresql://cockroachdb1:26257/account?sslmode=disable",
		//CockroachdbConnUrl: "host=cockroachdb1 port=26257 user=root dbname=account sslmode=disable",
		//CockroachdbConnUrl: "host=cockroachdb1 port=26257 user=account dbname=account password=password sslmode=disable",
		//CockroachdbConnUrl: "host=192.168.99.130 port=26257 user=root dbname=account sslmode=disable",
		ZipkinServerUrl:    "http://zipkin:9411",
		ServerConfig: ServerConfig{
			Name: "dataservice",
			Port: "7070",
		},
		AmqpConfig: AmqpConfig{
			ServerUrl: "amqp://guest:guest@rabbitmq:5672/",
		},
	}
}
