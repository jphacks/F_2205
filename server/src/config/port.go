package config

import "github.com/jphacks/F_2205/server/src/utils/helper"

// Portはサーバーのport番号を返します
func Port() string {
	return helper.GetEnvOrDefault("PORT", "8080")
}
