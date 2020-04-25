package generator

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/smiyaguchi/headwater/config"
	"github.com/smiyaguchi/headwater/generator/faker"
	"github.com/smiyaguchi/headwater/generator/writer"
	"github.com/smiyaguchi/headwater/schema"
)

type HistoryGenerator struct{}

func (hg *HistoryGenerator) Generate(schema schema.Schema, config config.Config) error {
	if !schema.HasHistory && !schema.HasFrom && !schema.HasTo {
		return errors.New("history or from and to field is required")
	}

	if !schema.HasHistory && (!schema.HasFrom || !schema.HasTo) {
		return errors.New("If history field is nothing from and to field is requried")
	}

	data := make([][]string, config.Count)

	for i := 0; i < config.Count; i++ {
		if config.Header && i == 0 {
			data[i] = schema.Names()
			continue
		}

		d := faker.Fake(schema, config.Loss)

        var historyData [][]string
		if schema.HasHistory {
			historyData = hg.generateHistory(schema, config.Count-i, d.RowValue)
		} else {
			historyData = hg.generateHistoryFromTo(schema, config.Count-i, d.RowValue)
		}

		for j := 0; j < len(historyData); j++ {
			data[i] = historyData[j]
			if j != (len(historyData) - 1) {
				i++
			}
		}
	}

	if err := writer.Write(data, config.Quote); err != nil {
		return err
	}

	return nil
}

func (hg *HistoryGenerator) generateHistory(schema schema.Schema, countRange int, row []string) [][]string {
	h := newDateInfo(schema, schema.IndexHistory(), row)
	historyCount := hg.generateHistoryCount(countRange)
	data := make([][]string, historyCount)

	var diffDays int = int(time.Since(h.Time).Seconds()) / 60 / 60 / 24

	if historyCount == 1 || diffDays <= 1 {
		data[0] = row
		return data
	}

	if diffDays <= 5 {
		historyCount = 2
	}

	var span int = diffDays / historyCount

	for i := 0; i < historyCount; i++ {
		cr := hg.copyRow(row)
		if i == 0 {
			data[0] = cr
			continue
		}
		h.Time = h.addTimeSpan(span)
		cr[h.Index] = h.Time.Format(h.Format)
		data[i] = cr
	}

	return data
}

func (hg *HistoryGenerator) generateHistoryFromTo(schema schema.Schema, countRange int, row []string) [][]string {
	from := newDateInfo(schema, schema.IndexFrom(), row)
	to := newDateInfo(schema, schema.IndexTo(), row)

	historyCount := hg.generateHistoryCount(countRange)
	data := make([][]string, historyCount)

	var diffDays = int(time.Since(from.Time).Seconds()) / 60 / 60 / 24

	if historyCount == 1 || diffDays <= 1 {
		cr := hg.copyRow(row)
		cr[to.Index] = to.LastTime
		data[0] = cr
		return data
	}

	if diffDays <= 5 {
		historyCount = 2
	}

	var span int = diffDays / historyCount

	to.Time = from.Time
	for i := 0; i < historyCount; i++ {
		cr := hg.copyRow(row)
		if i == 0 {
			to.Time = to.addTimeSpan(span)
			to.Time = to.roundTimeZero(to.Time)
			cr[to.Index] = to.Time.Add(-1 * time.Nanosecond).Format(to.Format)
			data[i] = cr
			continue
		}
		cr[from.Index] = to.Time.Format(from.Format)
		to.Time = to.addTimeSpan(span)

		if i < (historyCount - 1) {
			cr[to.Index] = to.Time.Add(-1 * time.Nanosecond).Format(to.Format)
		} else {
			cr[to.Index] = to.LastTime
		}
		data[i] = cr
	}

	return data
}

type dateInfo struct {
	Index    int
	Type     string
	Format   string
	LastTime string
	Time     time.Time
}

func newDateInfo(schema schema.Schema, index int, row []string) dateInfo {
	d := dateInfo{
		Index: index,
	}

	d.Type = d.getType(schema, d.Index)
	d.Format = d.generateFormat(d.Type)
	d.LastTime = d.generateLastTime(d.Type)

	t, _ := time.Parse(d.Format, row[d.Index])
	d.Time = t

	return d
}

func (di *dateInfo) generateFormat(typ string) string {
	format := ""

	switch strings.ToUpper(typ) {
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

func (di *dateInfo) generateLastTime(typ string) string {
	t := ""

	switch strings.ToUpper(typ) {
	case "DATE":
		t = "9999-12-31"
	case "DATETIME":
		t = "9999-12-31 23:59:59"
	case "TIMESTAMP":
		t = "9999-12-31 23:59:59.999999 UTC"
	default:
		t = ""
	}

	return t
}

func (di *dateInfo) getType(schema schema.Schema, index int) string {
	return schema.Columns[index].Type
}

func (di *dateInfo) addTimeSpan(span int) time.Time {
	return di.Time.Add(time.Hour * 24 * time.Duration(span))
}

func (di *dateInfo) roundTimeZero(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func (hg *HistoryGenerator) generateHistoryCount(countRange int) int {
	rand.Seed(time.Now().UnixNano())

	historyCount := rand.Intn(6)
	if countRange < 6 || historyCount == 0 {
		historyCount = 1
	}

	return historyCount
}

func (hg *HistoryGenerator) copyRow(row []string) []string {
	copyRow := make([]string, len(row))
	copy(copyRow, row)
	return copyRow
}
