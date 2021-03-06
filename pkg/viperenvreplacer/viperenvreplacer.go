package viperenvreplacer

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func ViperReplaceEnvs() {
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
}
