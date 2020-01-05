package schema

type Schema struct {
	ColumnName string `json:"column"`
	Type       string `json:"type"`
	Precision  uint8  `json:"precision,string"`
	Scale      uint8  `json:"scale,string"`
	Unique     bool   `json:"unique,string"`
	Nullable   bool   `json:"nullable,string"`
}
