package utils

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)


type Database struct {
	host		string		`mapstructure:"dbhost"`
	port        int			`mapstructure:"dbport"`
	database    string		`mapstructure:"database"`
	user        string		`mapstructure:"dbuser"`
	password    string		`mapstructure:"dbpassword"`
}

type Server struct {
	name        string		`mapstructure:"name"`
	host        string		`mapstructure:"host"`
	port        int			`mapstructure:"port"`
}

type Config struct {
	server 		Server		`mapstructure:"server"`
	database	Database	`mapstructure:"database"`
}
func LoadToml(configPath string) Config {
	var config Config
	viper := viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigFile("config")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}


	fmt.Println(viper.Get("servername"))
	toml.DecodeFile(configPath, &config)
	return config
}

func ReadFile(filePath string) []byte {

	file, fileErr := os.Open(filePath)
	if fileErr!=nil{
		panic(fileErr)
	}
	fmt.Println(file)
	defer file.Close()

	fileContent, fileContentErr := ioutil.ReadAll(file)
	if fileContentErr!=nil{
		panic(fileContentErr)
	}
	fmt.Println(fileContent)
	return fileContent


}