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
		RowValue: make([]string, len(schema.Columns)),
		Key: "",
		From: "",	
		To: "",
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
		if t == "STRING" {
			d = gofakeit.LastName() + "_" + gofakeit.FirstName()
		} else if t == "INTEGER" {
			d = strconv.FormatUint(gofakeit.Uint64(), 10)
		} else if t == "NUMERIC" {
			d = strconv.FormatInt(gofakeit.Int64(), 10)
		} else if t == "FLOAT" {
			d = strconv.FormatFloat(gofakeit.Float64(), 'e', 9, 64)
		} else if t == "BOOLEAN" {
			d = strconv.FormatBool(gofakeit.Bool())
		} else if t == "TIMESTAMP" {
			d = gofakeit.Date().String()
		} else if t == "DATE" {
			d = gofakeit.Date().Format("2006-01-02")
		} else if t == "TIME" {
			d = gofakeit.Date().Format("15:04:05")
		} else if t == "DATETIME" {
			d = gofakeit.Date().Format("2006-01-02 15:04:05")
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
