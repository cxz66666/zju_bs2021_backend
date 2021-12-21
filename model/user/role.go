package user

type Role int

const (
	NoLogin  Role = 0
	Staff    Role = 1
	Admin    Role = 2
	SysAdmin Role = 3
)
