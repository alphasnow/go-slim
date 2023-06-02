package xsnowflake

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func NewFlake(cfg *Config) *sonyflake.Sonyflake {
	st := sonyflake.Settings{}
	st.StartTime = cfg.getStartTime()
	st.MachineID = func() (uint16, error) {
		return cfg.getMachineID(), nil
	}
	sf = sonyflake.NewSonyflake(st)
	return sf
}

type SnowFlake struct {
	sf *sonyflake.Sonyflake
}

func NewSnowFlake(cfg *Config) *SnowFlake {
	sf := NewFlake(cfg)
	return &SnowFlake{sf: sf}
}

func (s *SnowFlake) Generate() (uint64, error) {
	return sf.NextID()
}
