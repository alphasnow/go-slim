package xsnowflake

import "time"

type Config struct {
	StartTime string `yaml:"startTime"`
	MachineID int    `yaml:"machineID"`
}

func (c *Config) getStartTime() time.Time {
	tt, err := time.Parse("2006-01-02 15:04:05", c.StartTime)
	if err != nil {
		panic(err)
	}
	return tt
}

func (c *Config) getMachineID() uint16 {
	return uint16(c.MachineID)
}

func NewConfig() *Config {
	return &Config{
		StartTime: "2020-01-01 00:00:00",
		MachineID: 2020, // uint16,最大值 65535
	}
}
