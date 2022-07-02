package config

import (
	"github.com/spf13/viper"

	"github.com/avoropaev/otus-go-banner-rotator/pkg/viperenvreplacer"
)

type BannerRotatorConfig struct {
	Logger   LoggerConf   `mapstructure:"logger"`
	DB       PSQLConf     `mapstructure:"psql"`
	GRPC     GRPCConf     `mapstructure:"grpc"`
	Producer ProducerConf `mapstructure:"producer"`
}

type GRPCConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func NewBannerRotatorConfig() *BannerRotatorConfig {
	return &BannerRotatorConfig{
		Logger: LoggerConf{
			Level: "info",
		},
	}
}

func ParseBannerRotatorConfig(filePath string) (*BannerRotatorConfig, error) {
	c := NewBannerRotatorConfig()

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viperenvreplacer.ViperReplaceEnvs()

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
