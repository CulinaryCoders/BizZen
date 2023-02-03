package controllers

import (
	"server/config"
)

// TODO: Add comment documentation (func GetSigningKey)
func GetSigningKey() []byte {
	// TODO: Config var or specific "JWT_KEY" value should probably be passed to this function instead of re-initializing entire config.
	config.InitEnvConfigs()
	var mySigningKey = []byte(config.ConfigVars.JWT_Key)
	return mySigningKey
}
