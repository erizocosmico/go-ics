package ics

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	uts               = "1136239445"
	icsFormat         = "20060102T150405Z"
	icsFormatWholeDay = "20060102"
)

func downloadFromURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func trimFieldName(field string) string {
	parts := strings.SplitN(field, ":", 2)
	return strings.TrimSpace(parts[len(parts)-1])
}

func trimField(field, cutset string) string {
	re, _ := regexp.Compile(cutset)
	cutsetRem := re.ReplaceAllString(field, "")
	return strings.TrimRight(cutsetRem, "\r\n")
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func parseDayNameToIcsName(day string) string {
	switch day {
	case "Mon":
		return "MO"
	case "Tue":
		return "TU"
	case "Wed":
		return "WE"
	case "Thu":
		return "TH"
	case "Fri":
		return "FR"
	case "Sat":
		return "ST"
	case "Sun":
		return "SU"
	default:
		return ""
	}
}
