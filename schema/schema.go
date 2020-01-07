package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Column struct {
	ColumnName string `json:"column"`
	Type       string `json:"type"`
	Precision  uint8  `json:"precision,string"`
	Scale      uint8  `json:"scale,string"`
	Unique     bool   `json:"unique,string"`
	Nullable   bool   `json:"nullable,string"`
}

type Schema struct {
	Columns []Column
}

func ReadFile(path string) (Schema, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Schema{}, err
	}

	var columns []Column
	if err := json.Unmarshal(bytes, &columns); err != nil {
		return Schema{}, err
	}

	if err := validate(columns); err != nil {
		return Schema{}, err
	}

	var schema = Schema{Columns: columns}
	return schema, nil
}

func validate(columns []Column) error {
	if err := validateType(columns); err != nil {
		return err
	}

	return nil
}

func validateType(columns []Column) error {
	for _, v := range columns {
		t := strings.ToUpper(v.Type)
		if t == "STRING" ||
			t == "BYTES" ||
			t == "INTEGER" ||
			t == "FLOAT" ||
			t == "NUMERIC" ||
			t == "BOOLEAN" ||
			t == "TIMESTAMP" ||
			t == "DATE" ||
			t == "TIME" ||
			t == "DATETIME" ||
			t == "GEOGRAPHY" ||
			t == "RECORD" {

				return nil
		}
		return errors.New(fmt.Sprintf("Not support type %s", v.Type))	
	}
	return nil
}
