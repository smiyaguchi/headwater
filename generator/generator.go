package generator

import (
	"strings"

	"github.com/smiyaguchi/headwater/config"
	"github.com/smiyaguchi/headwater/schema"
)

type Generator interface {
	Generate(schema schema.Schema, config config.Config) error
}

func New(mode string) Generator {
	var g Generator

	switch strings.ToUpper(mode) {
	case "HISTORY":
		g = &HistoryGenerator{}
	default:
		g = &NormalGenerator{}
	}
	return g
}
