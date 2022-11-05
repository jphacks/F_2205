package generate

import (
	"crypto/rand"
	"fmt"
)

// MakeRandomStrFromLettersは指定したlengthの長さの文字列を生成します
func MakeRandomStrFromLetters(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

	// 乱数を生成
	b := make([]byte, length) // ex: [123 242 ...]
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("MakeRandomStrFromLetters Read Error : %w", err)
	}

	// letters からランダムに取り出して文字列を生成
	var res string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		res += string(letters[int(v)%len(letters)])
	}
	return res, nil
}
