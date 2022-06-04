package config

type Config struct {
	KafkaURL string
}

func NewConfig() Config {
	return Config{
		KafkaURL: "localhost:9092",
	}
}
