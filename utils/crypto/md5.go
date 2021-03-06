package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

//EncodeMD5 accept a string value, and return md5(value) by hex string
func EncodeMD5(value string) string {
	m:=md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

