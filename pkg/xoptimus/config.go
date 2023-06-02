package xoptimus

type Config struct {
	Prime      uint64
	ModInverse uint64
	Random     uint64
}

func NewConfig() *Config {
	// https://github.com/pjebs/optimus-go
	//Prime: 350708767
	//Inverse: 158275551
	//Random: 126144979
	//Bit length: 31
	return &Config{Prime: 350708767, ModInverse: 158275551, Random: 126144979}
}
