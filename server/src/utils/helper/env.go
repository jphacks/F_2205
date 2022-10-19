package helper

import "os"

// GetEnvOrDefaultはenvPathに指定された環境変数を取得する
// 取得できなかった場合はdefaultEnvを利用する
func GetEnvOrDefault(envPath string, defaultEnv string) string {
	env := os.Getenv(envPath)
	if env == "" {
		return defaultEnv
	}
	return env
}
