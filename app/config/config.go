package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type ServerConfig struct {
	Bind                string `mapstructure:"server_bind,omitempty"`
	Port                string `mapstructure:"server_port,omitempty"`
	Schema              string `mapstructure:"server_schema,omitempty"`
	ReadTimeout         int    `mapstructure:"server_read_timeout,omitempty"`
	WriteTimeout        int    `mapstructure:"server_write_timeout,omitempty"`
	TokenExpireDuration int    `mapstructure:"server_token_expire_duration,omitempty"`
	Salt                string `mapstructure:"server_salt,omitempty"`
	TokenKey            string `mapstructure:"server_token_key,omitempty"`
	LogLevel            string `mapstructure:"server_log_level" `
	LogShowPath         bool   `mapstructure:"server_log_show_path"`
	ExternalDomainName  string `mapstructure:"server_external_domain_name"`
	ExternalPort        string `mapstructure:"server_external_port"`
}

type LdapConfig struct {
	Enable      bool   `mapstructure:"enable,omitempty"`
	Url         string `mapstructure:"url,omitempty"`
	Port        int    `mapstructure:"port,omitempty"`
	Domain      string `mapstructure:"domain,omitempty"`
	RootDn      string `mapstructure:"root_dn,omitempty"`
	MemberOf    string `mapstructure:"member_of,omitempty"`
	TypeKeyWord string `mapstructure:"type_key_word,omitempty"`
	DefPass     string `mapstructure:"def_pass,omitempty"`
}

type DbConfig struct {
	Name              string `mapstructure:"db_name,omitempty"`
	User              string `mapstructure:"db_user,omitempty"`
	Password          string `mapstructure:"db_password,omitempty"`
	Type              string `mapstructure:"db_type,omitempty"`
	Host              string `mapstructure:"db_host,omitempty"`
	Port              string `mapstructure:"db_port,omitempty"`
	SSLMode           string `mapstructure:"db_ssl_mode"`
	MigrationsDirPath string `mapstructure:"db_migration_path"`
}

type TestDbConfig struct {
	Name              string `mapstructure:"test_db_name,omitempty"`
	User              string `mapstructure:"test_db_user,omitempty"`
	Password          string `mapstructure:"test_db_password,omitempty"`
	Type              string `mapstructure:"test_db_type,omitempty"`
	Host              string `mapstructure:"test_db_host,omitempty"`
	Port              string `mapstructure:"test_db_port,omitempty"`
	SSLMode           string `mapstructure:"test_db_ssl_mode"`
	MigrationsDirPath string `mapstructure:"test_db_migration_path"`
}

func InitConfig() {

	wddir, _ := os.Getwd()
	appDir := strings.Split(wddir, "/app")[0] + "/app"

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../app")
	viper.AddConfigPath(appDir)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("ARCHIVE")

	err := viper.ReadInConfig()

	if err != nil {
		logrus.Errorf("Error read config file: default:  %v", err)
		os.Exit(1)
	}

	viper.AutomaticEnv()
	logrus.Info("Config prepared")

	logrus.Info("Loaded config database: %v", GetDbConfig())
	logrus.Info("Loaded config database: %v", GetTestDbConfig())
	logrus.Info("Loaded config server: %v", GetServerConfig())
	logrus.Info("Loaded config ldap: %v", GetLdapConfig())
}

func GetServerConfig() *ServerConfig {
	serverConfig := &ServerConfig{}

	if err := viper.Unmarshal(&serverConfig); err != nil {
		logrus.Panicf("Can't load server config. %v", err)
		os.Exit(1)
	}

	return serverConfig
}

func GetLdapConfig() *LdapConfig {
	ldapConfig := &LdapConfig{}

	if err := viper.Unmarshal(&ldapConfig); err != nil {
		logrus.Panicf("Can't load ldap config. %v", err)
		os.Exit(1)
	}

	return ldapConfig
}

func GetDbConfig() *DbConfig {
	dbConfig := &DbConfig{}

	if err := viper.Unmarshal(&dbConfig); err != nil {
		logrus.Panicf("Can't load db_adapter config. %v", err)
		os.Exit(1)
	}

	return dbConfig
}

func GetTestDbConfig() *DbConfig {
	testdbConfig := &TestDbConfig{}

	if err := viper.Unmarshal(&testdbConfig); err != nil {
		logrus.Panicf("Can't load db_adapter config. %v", err)
		os.Exit(1)
	}

	dbConfig := DbConfig{
		Name:              testdbConfig.Name,
		User:              testdbConfig.User,
		Password:          testdbConfig.Password,
		Type:              testdbConfig.Type,
		Host:              testdbConfig.Host,
		Port:              testdbConfig.Port,
		SSLMode:           testdbConfig.SSLMode,
		MigrationsDirPath: testdbConfig.MigrationsDirPath,
	}

	return &dbConfig
}
