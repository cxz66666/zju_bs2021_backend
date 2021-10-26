package uuid

import (
	"fmt"
	"testing"
)

func TestGenUUID(t *testing.T) {
	fmt.Println(GenUUID())
}

func TestGenUUIDString(t *testing.T) {
	fmt.Println(GenUUIDString())
}
