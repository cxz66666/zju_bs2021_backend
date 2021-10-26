// package uuid
package uuid

import "github.com/google/uuid"

func GenUUID() uuid.UUID{
	return uuid.New()
}

func GenUUIDString()string{
	return uuid.NewString()
}