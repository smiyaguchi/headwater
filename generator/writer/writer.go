package writer

import (
	"encoding/csv"
	"os"
)

func Write(data [][]string) error {
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
