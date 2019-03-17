package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ENV_DEV     = "dev"
	ENV_PROD    = "prod"
	ENV_TEST    = "test"
	ENV_CI      = "ci"
	ENV_STAGING = "staging"
)

type Config interface {
	IsDev() bool
	IsProd() bool
	IsTest() bool
	IsCI() bool
	IsStaging() bool

	GetName() string
	GetHttpPort() string
	GetTlsverify() bool
	GetCertFile() string
	GetKeyFile() string

	GetRpcPort() string

	SetConfigFile(path string)
}

// The base configuration file can be embedded by app's specified configuration
type BaseConfig struct {
	ConfigFile  string
	Env         string
	ServiceName string
	HttpPort    string
	Tlsverify   bool
	CertFile    string
	KeyFile     string
	RpcPort     string
}

func (this *BaseConfig) IsDev() bool     { return this.Env == ENV_DEV }
func (this *BaseConfig) IsProd() bool    { return this.Env == ENV_PROD }
func (this *BaseConfig) IsTest() bool    { return this.Env == ENV_TEST }
func (this *BaseConfig) IsCI() bool      { return this.Env == ENV_CI }
func (this *BaseConfig) IsStaging() bool { return this.Env == ENV_STAGING }

func (this *BaseConfig) GetHttpPort() string       { return this.HttpPort }
func (this *BaseConfig) GetRpcPort() string        { return this.RpcPort }
func (this *BaseConfig) GetTlsverify() bool        { return this.Tlsverify }
func (this *BaseConfig) GetCertFile() string       { return this.CertFile }
func (this *BaseConfig) GetKeyFile() string        { return this.KeyFile }
func (this *BaseConfig) SetConfigFile(path string) { this.ConfigFile = path }

func LoadConfigAndSetupEnv(conf Config, env string, configPath string) {
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey(fmt.Sprintf("%s.%s", conf.GetName(), env), conf)
	if err != nil {
		panic(err)
	}

	conf.SetConfigFile(configPath)

	fmt.Printf("\nThe configuration is: \n%+v\n\n", conf)
}
