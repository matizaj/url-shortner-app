package url

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ShortenUrl(originalUrl string) string {
	h := sha256.New()
	h.Write([]byte(originalUrl))
	fmt.Println(h.Sum(nil))
	hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(hash)
	shortURL := hash[:8]
	return shortURL
}
