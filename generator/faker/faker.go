package faker

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/smiyaguchi/headwater/schema"
)

var keys = make(map[string]int)

type Data struct {
	ColumnValue map[string]string
	RowValue    []string
	Key         string
	From        string
	To          string
}

func Fake(schema schema.Schema, loss bool) Data {
	var data = Data{
		ColumnValue: make(map[string]string),
		RowValue:    make([]string, len(schema.Columns)),
		Key:         "",
		From:        "",
		To:          "",
	}

	gofakeit.Seed(time.Now().UnixNano())

	for i, c := range schema.Columns {
		if loss && !c.Unique && c.Mode == "NULLABLE" {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(10) > 8 {
				continue
			}
		}

		d := ""
		t := strings.ToUpper(c.Type)

		switch t {
		case "STRING":
			d = gofakeit.LastName() + "_" + gofakeit.FirstName()
		case "INTEGER", "NUMERIC":
			d = strconv.Itoa(rand.Intn(1000))
		case "FLOAT":
			d = strconv.FormatFloat(gofakeit.Float64(), 'e', 9, 64)
		case "BOOLEAN":
			d = strconv.FormatBool(gofakeit.Bool())
		case "TIMESTAMP":
			d = gofakeit.DateRange(time.Unix(0, 484633944473634951), time.Now()).Format("2006-01-02 15:04:05.000000 MST")
		case "DATE":
			d = gofakeit.DateRange(time.Unix(0, 484633944473634951), time.Now()).Format("2006-01-02")
		case "TIME":
			d = gofakeit.DateRange(time.Unix(0, 484633944473634951), time.Now()).Format("15:04:05")
		case "DATETIME":
			d = gofakeit.DateRange(time.Unix(0, 484633944473634951), time.Now()).Format("2006-01-02 15:04:05")
		default:
			panic("Not support type")
		}

		data.ColumnValue[c.Name] = d
		data.RowValue[i] = d

		if c.Unique {
			data.Key += d
		}

		if c.From {
			data.From = d
		}

		if c.To {
			data.To = d
		}
	}

	return data
}
