package generator

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/smiyaguchi/headwater/schema"
	"github.com/brianvoe/gofakeit"
)

var keys = make(map[string]int)

func Generate(schema schema.Schema, count int) {
	data := make([][]string, count)	
	gofakeit.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		d := make([]string, len(schema.Columns))
		key := ""
		for j, v := range schema.Columns {
			if !v.Unique && v.Mode == "NULLABLE" {
				rand.Seed(time.Now().UnixNano())
				if rand.Intn(10) > 8 {
					continue
				}				
			}

			t := strings.ToUpper(v.Type)
			if t == "STRING" {
				d[j] = gofakeit.Name() + "_" + gofakeit.City()
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
			
			if v.Unique {
				key += d[j]
			}
		}
		if _, exist := keys[key]; exist {
			i--
			continue
		}
		data[i] = d
		keys[key] = 0
	}
	
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(data)
	
	if err := w.Error(); err != nil {
		panic(err)
	}
}
