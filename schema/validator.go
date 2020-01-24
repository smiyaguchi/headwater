package schema

import (
	"errors"
	"fmt"
	"strings"
)

type Validator struct {
}

func (v *Validator) Type(columns []Column) error {
	validTypeMap := map[string]bool{
		"STRING":    true,
		"BYTES":     true,
		"INTEGER":   true,
		"FLOAT":     true,
		"NUMERIC":   true,
		"BOOLEAN":   true,
		"TIMESTAMP": true,
		"DATE":      true,
		"TIME":      true,
		"DATETIME":  true,
		"GEOGRAPHY": true,
	}

	for _, c := range columns {
		t := strings.ToUpper(c.Type)

		if _, exist := validTypeMap[t]; !exist {
			return errors.New(fmt.Sprintf("Not support type %s", c.Type))
		}
	}

	return nil
}

func (v *Validator) Mode(columns []Column) error {
	validModeMap := map[string]bool{
		"NULLABLE": true,
		"REQUIRED": true,
		"REPEATED": true,
	}

	for _, c := range columns {
		m := strings.ToUpper(c.Mode)
		if _, exist := validModeMap[m]; !exist {
			return errors.New(fmt.Sprintf("Not support mode %s", c.Mode))
		}
	}

	return nil
}

func (v *Validator) FromTo(columns []Column) error {
	for _, c := range columns {
		if c.From || c.To {
			if c.Type != "DATE" && c.Type != "DATETIME" && c.Type != "TIMESTAMP" {
				return errors.New(fmt.Sprintf("From and To field is support type DATE DATETIME TIMESTAMP"))
			}
		}
	}
	return nil
}
