package generator

import (
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/generator/writer"
	"github.com/smiyaguchi/headwater/schema"
)

var keys = make(map[string]int)

func Generate(schema schema.Schema, count int, loss bool) {
	data := make([][]string, count)

	for i := 0; i < count; i++ {
		d := faker.Fake(schema, loss)

		if _, exist := keys[d.Key]; exist {
			i--
			continue
		}
		data[i] = d.RowValue
		keys[d.Key] = 0
	}

	if err := writer.Write(data); err != nil {
		panic(err)
	}
}
