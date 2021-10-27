package user_service

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCreateUser(t *testing.T) {
	if m, err := regexp.MatchString(PasswordRegex, "cxz666"); !m {
		fmt.Println(m,err)
	}
}