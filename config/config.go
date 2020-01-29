package config

import (
	"strings"
)

type Config struct {
	Count  int
	Loss   bool
	Header bool
	Quote  string
}

func New(count int, loss bool, header bool, quote string) Config {
	if header {
		count += 1
	}

	return Config{
		Count:  count,
		Loss:   loss,
		Header: header,
		Quote:  strings.ToUpper(quote),
	}
}
