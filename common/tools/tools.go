package tools

import (
	"hash/fnv"
	"newdeal/common"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

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
		return common.EmptyString
	}
	return string(r[startIndex:endIndex])
}

func GetDetailKeyword(keyword string) string {
	if IsEmptyStr(keyword) {
		return common.HankakuPercentMark
	}
	var builder strings.Builder
	builder.WriteString(common.HankakuPercentMark)
	for _, ch := range keyword {
		builder.WriteString(string(ch))
		builder.WriteString(common.HankakuPercentMark)
	}
	return builder.String()
}

func PtString2String(str *string) string {
	if str == nil {
		return common.EmptyString
	}
	return *str
}

func UUIDToInt64(u uuid.UUID) int64 {
	h := fnv.New64a()
	write, err := h.Write(u[:])
	if err != nil {
		return 0
	} // 全体を使う
	return int64(write)
}

func GenerateHashPassword(password string) (string, error) {
	// 第2引数はコスト（デフォルトは bcrypt.DefaultCost ≒ 10）
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckHashPassword(hashPass, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(rawPassword))
	return err == nil
}

func IsEmptyStr(s string) bool {
	return strings.TrimSpace(s) == common.EmptyString
}
