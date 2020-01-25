package generator

import (
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/generator/writer"
	"github.com/smiyaguchi/headwater/schema"
)

type NormalGenerator struct{}

var keys = make(map[string]int)

func (ng *NormalGenerator) Generate(schema schema.Schema, count int, loss bool, header bool) {
	if header {
		count += 1
	}
	data := make([][]string, count)

	for i := 0; i < count; i++ {
		if header && i == 0 {
			data[i] = schema.Names()
			continue
		}

		d := faker.Fake(schema, loss)

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
		panic(err)
	}
}
