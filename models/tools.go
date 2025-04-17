package models

import "time"

func Unix2DateTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("yyyy-MM-dd hh:mm:ss")
}

func DateTime2Unix(str string) int64 {
	template := "yyyy-MM-dd hh:mm:ss"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func Substr(rawString string, startIndex, endIndex int) string {
	r := []rune(rawString)
	if startIndex < 0 || endIndex > len(r) || startIndex > endIndex {
		return ""
	}
	return string(r[startIndex:endIndex])
}
