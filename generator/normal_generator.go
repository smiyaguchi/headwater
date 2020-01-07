package generator

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/smiyaguchi/headwater/schema"
	"github.com/brianvoe/gofakeit"
)

func Generate(schema schema.Schema, count int) {
	data := make([][]string, count)	
	gofakeit.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		d := make([]string, len(schema.Columns))
		for j, v := range schema.Columns {
			t := strings.ToUpper(v.Type)
			if t == "STRING" {
				d[j] = gofakeit.Sentence(1)
			} else if t == "INTEGER" {
				d[j] = strconv.FormatUint(gofakeit.Uint64(), 10)
			} else if t == "NUMERIC" {
				d[j] = strconv.FormatInt(gofakeit.Int64(), 10)
			} else if t == "FLOAT" {
				d[j] = strconv.FormatFloat(gofakeit.Float64(), 'e', 9, 64)
			} else if t == "NUMERIC" {
				d[j] = strconv.FormatFloat(gofakeit.Float64(), 'e', 9, 64)
			} else if t == "BOOLEAN" {
				d[j] = strconv.FormatBool(gofakeit.Bool())
			} else if t == "TIMESTAMP" {
				d[j] = gofakeit.Date().String()	
			} else if t == "DATE" {
				d[j] = gofakeit.Date().Format("2006-01-02")
			} else if t == "TIME" {
				d[j] = gofakeit.Date().Format("15:04:05")
			} else if t == "DATETIME" {
				d[j] = gofakeit.Date().Format("2006-01-02 15:04:05")
			}
		}
		data[i] = d
	}
	
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(data)
	
	if err := w.Error(); err != nil {
		panic(err)
	}
}
