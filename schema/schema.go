package schema

import (
	"encoding/json"
	"io/ioutil"
)

type Column struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Precision uint8  `json:"precision,string"`
	Scale     uint8  `json:"scale,string"`
	Unique    bool   `json:"unique,string"`
	Mode      string `json:"mode"`
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
	v := new(Validator)

	if err := v.Type(columns); err != nil {
		return err
	}

	if err := v.Mode(columns); err != nil {
		return err
	}

	return nil
}
