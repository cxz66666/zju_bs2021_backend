package crypto

import "encoding/base64"

// Base64Encode the string content, and return the base64(content)
func Base64Encode(content string)(string)  {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

// Base64Decode the string content, and return the original string and error (if exists)
func Base64Decode(content string) (string,error) {
	s,err:= base64.StdEncoding.DecodeString(content)
	return string(s),err
}




