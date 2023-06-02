package constants

type Gender int8

const (
	GenderUnknown Gender = iota
	GenderMan
	GenderWoman
)

type JWTModel string

const (
	JWTUser  JWTModel = "user"
	JWTAdmin JWTModel = "admin"
)
