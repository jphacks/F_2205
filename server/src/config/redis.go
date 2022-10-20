package config

import (
	"os"
)

// GetRedisUrlはredisの接続先情報を取得します、もし見つからなかった場合はアプリを終了します
func GetRedisUrl() string {
	redisUrl := os.Getenv("REDIS_URL")
	// if redisUrl == "" {
	// 	log.Fatal("error: REDIS_URL not found")
	// }
	return redisUrl
}
