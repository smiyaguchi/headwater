package writer

import (
	"os"

	csv "github.com/JensRantil/go-csv"
)

func Write(data [][]string, quote string) error {
	d := csv.Dialect{
		Delimiter:      csv.DefaultDelimiter,
		Quoting:        csv.QuoteNone,
		DoubleQuote:    csv.DefaultDoubleQuote,
		EscapeChar:     csv.DefaultEscapeChar,
		QuoteChar:      csv.DefaultQuoteChar,
		LineTerminator: csv.DefaultLineTerminator,
		Comment:        csv.DefaultComment,
	}

	switch quote {
	case "ALL":
		d.Quoting = csv.QuoteAll
	case "NONNUMERIC":
		d.Quoting = csv.QuoteNonNumeric
	default:
		d.Quoting = csv.QuoteNone
	}

	w := csv.NewDialectWriter(os.Stdout, d)
    if err := w.WriteAll(data); err != nil {
        return err
    }

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
