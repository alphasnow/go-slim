package xoptimus

import "github.com/pjebs/optimus-go"

func NewOptimus(c *Config) *optimus.Optimus {
	op := optimus.New(c.Prime, c.ModInverse, c.Random)
	return &op
}
