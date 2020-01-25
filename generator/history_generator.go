package generator

import (
	"math/rand"
	"time"

	"github.com/smiyaguchi/headwater/config"
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/generator/writer"
	"github.com/smiyaguchi/headwater/schema"
)

type HistoryGenerator struct{}

func (hg *HistoryGenerator) Generate(schema schema.Schema, config config.Config) {
	data := make([][]string, config.Count)

	if !schema.HasFrom || !schema.HasTo {
		panic("From and To field is required")
	}

	for i := 0; i < config.Count; i++ {
		if config.Header && i == 0 {
			data[i] = schema.Names()
			continue
		}

		d := faker.Fake(schema, config.Loss)

		historyData := generateHistory(schema, config.Count-i, d.RowValue)

		for j := 0; j < len(historyData); j++ {
			data[i] = historyData[j]
			if j != (len(historyData) - 1) {
				i++
			}
		}
	}

	if err := writer.Write(data); err != nil {
		panic(err)
	}
}

func generateHistory(schema schema.Schema, countRange int, row []string) [][]string {
	historyCount := generateHistoryCount(countRange)

	data := make([][]string, historyCount)

	indexFrom, _ := schema.InfoFrom()
	indexTo, _ := schema.InfoTo()

	typeFrom := getType(schema, indexFrom)
	typeTo := getType(schema, indexTo)

	formatFrom := generateFormat(typeFrom)
	formatTo := generateFormat(typeTo)

	lastdayTo := generateLastDay(typeTo)

	from, _ := time.Parse(formatFrom, row[indexFrom])
	var diffDays = int(time.Now().Sub(from).Seconds()) / 60 / 60 / 24

	if historyCount == 1 || diffDays == 1 || diffDays == 0 {
		value := copyRow(row)
		value[indexTo] = lastdayTo
		data[0] = value
		return data
	}

	if diffDays <= 5 {
		historyCount = 2
	}

	var span int = diffDays / historyCount

	toDate := from
	for i := 0; i < historyCount; i++ {
		value := copyRow(row)
		if i == 0 {
			toDate = toDate.Add(time.Hour * 24 * time.Duration(span))
			toDate = time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 0, 0, 0, 0, time.UTC)
			value[indexTo] = toDate.Add(-1 * time.Nanosecond).Format(formatTo)
			data[i] = value
			continue
		}
		value[indexFrom] = toDate.Format(formatFrom)
		toDate = toDate.Add(time.Hour * 24 * time.Duration(span))

		if i < (historyCount - 1) {
			value[indexTo] = toDate.Add(-1 * time.Nanosecond).Format(formatTo)
		} else {
			value[indexTo] = lastdayTo
		}
		data[i] = value
	}

	return data
}

func generateHistoryCount(countRange int) int {
	rand.Seed(time.Now().UnixNano())

	historyCount := 0
	if countRange >= 6 {
		historyCount = rand.Intn(6)
	}

	if historyCount == 0 {
		historyCount = 1
	}

	return historyCount
}

func generateFormat(typ string) string {
	format := ""

	switch typ {
	case "DATE":
		format = "2006-01-02"
	case "DATETIME":
		format = "2006-01-02 15:04:05"
	case "TIMESTAMP":
		format = "2006-01-02 15:04:05.000000 MST"
	default:
		format = ""
	}

	return format
}

func generateLastDay(typ string) string {
	day := ""

	switch typ {
	case "DATE":
		day = "9999-12-31"
	case "DATETIME":
		day = "9999-12-31 23:59:59"
	case "TIMESTAMP":
		day = "9999-12-31 23:59:59.999999 UTC"
	default:
		day = ""
	}

	return day
}

func getType(schema schema.Schema, index int) string {
	return schema.Columns[index].Type
}

func copyRow(row []string) []string {
	copyRow := make([]string, len(row))
	copy(copyRow, row)
	return copyRow
}
