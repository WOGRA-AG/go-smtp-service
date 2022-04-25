package configReader

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type SmtpConfig struct {
	User          string
	Password      string
	Smtpserver    string
	Smtpport      string
	SenderAddress string
}

type AccessTokenList struct {
	Accesstokens []AccessToken `yaml:"Accesstokens"`
}

type AccessToken struct {
	Token    string `json:"accesstoken" yaml:"token"`
	User     string `json:"user"  yaml:"user"`
	Password string `json:"password"  yaml:"password"`
}

type RestConfig struct {
	Host string
	Port string
}

func ReadSmtpConfiguration() SmtpConfig {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	var smtpConf SmtpConfig

	err := viper.Unmarshal(&smtpConf)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	return smtpConf
}

func ReadRestConfiguration() RestConfig {

	// Set the file name of the configurations file
	viper.SetConfigName("rest")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	var restConf RestConfig

	err := viper.Unmarshal(&restConf)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	return restConf
}

func ReadAccessTokensConfiguration() []AccessToken {

	// Set the file name of the configurations file
	viper.SetConfigName("accesstoken")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	viper.Get("name")

	var access AccessTokenList

	err := viper.Unmarshal(&access)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode into struct, %v", err)
		log.Print(msg)
	}

	log.Print(fmt.Sprintf("Tokens read: %v", access))
	return access.Accesstokens
}

func CheckAccessToken(token *AccessToken) bool {

	if token == nil {
		return false
	}
	// at the moment only the token id will be checked
	// user and password will be ignored currently
	return HasAccessToken(token.Token)
}

func HasAccessToken(token string) bool {

	if token == "" {
		log.Println("Token is empty.")
		return false
	}

	tokens := ReadAccessTokensConfiguration()
	return contains(tokens, token)
}

func contains(s []AccessToken, e string) bool {
	for _, a := range s {
		if a.Token == e {
			log.Println("Given AccesToken found.")
			return true
		}
	}

	log.Println("Given AccesToken not found.")
	return false
}
