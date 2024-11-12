package parsing

import (
	"errors"
	"github.com/DoTuanAnh2k1/parsing-template/utils"
	"regexp"
	"strings"
	"time"
)

func findFormateDate(template string) string {
	return replaceFormatDate(template)
}

func replaceFormatDate(template string) string {
	// Create regex to find all format
	re := regexp.MustCompile(`/=FormatDate\([a-zA-Z 0-9-\/]*\,[a-zA-Z 0-9-]*\)`)

	// Replace all formatDecimal
	return re.ReplaceAllStringFunc(template, formatDate)
}

func formatDate(match string) string {
	match = strings.Split(match, "(")[1]
	match = strings.Split(match, ")")[0]
	parts := strings.Split(match, ",")
	res, err := convertToDateStandard(parts[0])
	utils.CheckError(err)
	return res
}

func convertToDateStandard(input string) (string, error) {
	// List format I have already known....
	dateFormats := []string{
		"2006-01-02",      // ISO 8601 (yyyy-mm-dd)
		"02/01/2006",      // dd/mm/yyyy
		"01/02/2006",      // mm/dd/yyyy
		"2006.01.02",      // yyyy.mm.dd
		"02-Jan-2006",     // dd-Mon-yyyy
		"Jan 02, 2006",    // Mon dd, yyyy
		"02 January 2006", // dd Month yyyy
		"January 2, 2006", // Month dd, yyyy
		"02-01-2006",      // dd-mm-yyyy
		"01-02-2006",      // mm-dd-yyyy
	}

	var parsedDate time.Time
	var err error

	// try format every format I know
	for _, format := range dateFormats {
		parsedDate, err = time.Parse(format, input)
		if err == nil {
			return parsedDate.Format("02/01/2006"), nil
		}
	}

	return "", errors.New("cannot format date")
}
