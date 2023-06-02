package xuuid

import "github.com/google/uuid"

func init() {
	uuid.EnableRandPool()
}

func Generate() string {
	return uuid.NewString()
}
