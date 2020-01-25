package config

type Config struct {
	Count  int
	Loss   bool
	Header bool
}

func New(count int, loss bool, header bool) Config {
	if header {
		count += 1
	}

	return Config{
		Count:  count,
		Loss:   loss,
		Header: header,
	}
}
