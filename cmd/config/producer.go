package config

type ProducerConf struct {
	URI          string `mapstructure:"uri"`
	Queue        string `mapstructure:"queue"`
	ExchangeName string `mapstructure:"exchange_name"`
	ExchangeType string `mapstructure:"exchange_type"`
	BindingKey   string `mapstructure:"binding_key"`
}
