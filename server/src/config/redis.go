package config

import "github.com/jphacks/F_2205/server/src/utils/helper"

// GetRedisUrlはredisの接続先情報を取得します
func GetRedisUrl() string {
	return helper.GetEnvOrDefault("REDIS_URL", "f_2205-redis:6379")
}
