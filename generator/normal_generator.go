package generator

import (
	"encoding/csv"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/schema"
)

var keys = make(map[string]int)

func Generate(schema schema.Schema, count int) {
	data := make([][]string, count)
	gofakeit.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		d := faker.Fake(schema)

		if _, exist := keys[d.Key]; exist {
			i--
			continue
		}
		data[i] = d.RowValue
		keys[d.Key] = 0
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		panic(err)
	}
}
