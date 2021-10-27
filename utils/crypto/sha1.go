package crypto

import (
	"annotation/utils/setting"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// SHA1 return the sha1(string), warning: it will transform string to upper
func SHA1(content string) string {
	h:=sha1.New()
	h.Write([]byte(content))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// Password2Secret convert the password(plaintext) to salted, use double SHA1 and saltA saltB, database store its output
func Password2Secret(password string) string {
	saltA:=setting.SecretSetting.SaltA
	saltB:=setting.SecretSetting.SaltB
	return SHA1(saltB+SHA1(saltA+password))
}
