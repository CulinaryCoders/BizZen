package controllers

import (
	"server/config"
)

func GetSigningKey() []byte {
	config.InitEnvConfigs()
	var mySigningKey = []byte(config.ConfigVars.JWT_Key)
	return mySigningKey
}
