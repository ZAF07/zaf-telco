package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string      `json:"port" mapstructure:"port"`
	HTTPServerConfig *HTTPConfig `json:"http_config" mapstructure:"http_config"`
	Datastore        *Datastore  `json:"datastore" mapstructure:"datastore"`
}

type HTTPConfig struct {
	Port         string        `json:"port" mapsturcture:"port"`
	ReadTimeout  time.Duration `json:"read_timeout" mapsturcture:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
}

type RPCConfig struct {
	Port         string        `json:"port" mapsturcture:"port"`
	ReadTimeout  time.Duration `json:"read_timeout" mapsturcture:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
}

type Datastore struct {
	// Primary *PrimaryDatastore
	PrimaryDatastore `json:"primary_database" mapstructure:"primary_database"`
}

type PrimaryDatastore struct {
	PrimaryDBType      string `json:"type" mapstructure:"type"`
	PrimaryDBName      string `json:"name" mapstructure:"name"`
	PrimaryDBHost      string `json:"host" mapstructure:"host"`
	PrimaryDBPort      string `json:"port" mapstructure:"port"`
	PrimaryDBPasswpord string `json:"password" mapstructure:"password"`
	PrimaryDBUser      string `json:"user" mapsturcture:"user"`
	PrimaryDBSSLMode   string `json:"ssl" mapstructure:"ssl"`
}

func (p *PrimaryDatastore) GetPrimaryConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", p.PrimaryDBHost, p.PrimaryDBPort, p.PrimaryDBUser, p.PrimaryDBName, p.PrimaryDBPasswpord, p.PrimaryDBSSLMode)
}

func InitConfig(filePath string) *Config {
	return loadConfig(filePath)
	// return &Config{
	// 	Port: ":8000",
	// 	HTTPServerConfig: &HTTPConfig{
	// 		Port:         ":8000",
	// 		ReadTimeout:  25 * time.Second,
	// 		WriteTimeout: 25 * time.Second,
	// 	},
	// 	Datastore: &Datastore{
	// 		PrimaryDatastore: PrimaryDatastore{
	// 			PrimaryDBType: "test",
	// 		},
	// 	},
	// }
}

func loadConfig(filePath string) *Config {
	AppConfig := &Config{}

	v := viper.New()
	v.SetConfigFile(filePath)
	v.ReadInConfig()
	unmarshalConfig(AppConfig, v)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("~~~ Config file '%+v' has been modified ~~~. Operation: %+v", e.Name, e.Op)
		unmarshalConfig(AppConfig, v)
	})

	return AppConfig
}

func unmarshalConfig(config *Config, v *viper.Viper) {
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("[CONFIG] Error unmarshaling app config : %+v\n", err)
	}
}
