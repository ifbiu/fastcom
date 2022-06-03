package common

import "time"

func FormatTime(tempTime time.Time) (string) {
	now := time.Now()
	timeStr := now.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	if now.Year() == tempTime.Year() &&
		now.Month() == tempTime.Month() &&
		now.Day() == tempTime.Day() {
		return tempTime.Format("15:04")
	}else if tempTime.Unix() > t.Unix() + 1 - 86400- 86400 &&
		tempTime.Unix() < t.Unix() - 86400{
		return tempTime.Format("昨天 15:04")
	}else if now.Year() == tempTime.Year(){
		return tempTime.Format("01月02日")
	}else{
		return tempTime.Format("2006年01月02日")
	}
}