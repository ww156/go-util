package util

import (
	"strconv"
	"time"
)

var (
	// Loc 时区
	Loc *time.Location
)

func init() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	Loc = loc
}

// FormatYMD 格式化年月日
func FormatYMD(format string) string {
	return time.Now().In(Loc).Format(format)
}

// FormatYMDH 格式化年月日时
func FormatYMDH() string {
	return time.Now().In(Loc).Format("2006010215")
}

// Week 获取周
func Week() string {
	return strconv.Itoa(int(time.Now().In(Loc).Weekday()))
}

// FormateLastYMD 昨天日期
func FormateLastYMD(format string) string {
	return time.Now().Add(time.Hour * -24).In(Loc).Format(format)
}
