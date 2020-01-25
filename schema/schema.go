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
	From      bool   `json:"from,string"`
	To        bool   `json:"to,string"`
}

type Schema struct {
	Columns   []Column
	HasUnique bool
	HasFrom   bool
	HasTo     bool
}

func (s *Schema) IndexFrom() int {
	for i, c := range s.Columns {
		if c.From {
			return i
		}
	}
	return -1
}

func (s *Schema) IndexTo() int {
	for i, c := range s.Columns {
		if c.To {
			return i
		}
	}
	return -1
}

func (s *Schema) Names() []string {
	names := make([]string, len(s.Columns))
	for i, c := range s.Columns {
		names[i] = c.Name
	}
	return names
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

	schema := Schema{
		Columns:   columns,
		HasUnique: hasUnique(columns),
		HasFrom:   hasFrom(columns),
		HasTo:     hasTo(columns),
	}

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

	if err := v.FromTo(columns); err != nil {
		return err
	}

	return nil
}

func hasUnique(columns []Column) bool {
	for _, c := range columns {
		if c.Unique {
			return true
		}
	}
	return false
}

func hasFrom(columns []Column) bool {
	for _, c := range columns {
		if c.From {
			return true
		}
	}
	return false
}

func hasTo(columns []Column) bool {
	for _, c := range columns {
		if c.To {
			return true
		}
	}
	return false
}
