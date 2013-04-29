package google

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func Code(secret string) (string, error) {
	secret = strings.ToUpper(secret)
	secret = strings.Replace(secret, " ", "", -1)
	key := make([]byte, base32.StdEncoding.DecodedLen(len(secret)))
	_, err := base32.StdEncoding.Decode(key, []byte(secret))
	if err != nil {
		return "", err
	}
	message := make([]byte, 8)
	binary.BigEndian.PutUint64(message, uint64(time.Now().Unix()/30))
	hmacsha1 := hmac.New(sha1.New, key)
	hmacsha1.Write(message)
	hash := hmacsha1.Sum([]byte{})
	offset := hash[len(hash)-1] & 0xF
	truncatedHash := hash[offset : offset+4]
	return fmt.Sprintf("%06d", (binary.BigEndian.Uint32(truncatedHash)&0x7FFFFFFF)%1000000), nil
}
