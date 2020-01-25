package generator

import (
	"github.com/smiyaguchi/headwater/config"
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/generator/writer"
	"github.com/smiyaguchi/headwater/schema"
)

type NormalGenerator struct{}

var keys = make(map[string]int)

func (ng *NormalGenerator) Generate(schema schema.Schema, config config.Config) error {
	data := make([][]string, config.Count)

	for i := 0; i < config.Count; i++ {
		if config.Header && i == 0 {
			data[i] = schema.Names()
			continue
		}

		d := faker.Fake(schema, config.Loss)

		if d.Key != "" {
			if _, exist := keys[d.Key]; exist {
				i--
				continue
			}
		}

		data[i] = d.RowValue
		keys[d.Key] = 0
	}

	if err := writer.Write(data); err != nil {
		return err
	}

	return nil
}
