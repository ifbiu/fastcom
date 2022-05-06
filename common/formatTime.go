package common

import "time"

func FormatTime(tempTime time.Time) (string) {
	now := time.Now()
	if now.Year() == tempTime.Year() &&
		now.Month() == tempTime.Month() &&
		now.Day() == tempTime.Day() {
		return tempTime.Format("15:04")
	}else if now.Year() == tempTime.Year() &&
		now.Month() == tempTime.Month() &&
		now.Day()-1 == tempTime.Day() {
		return tempTime.Format("昨天 15:04")
	}else if now.Year() == tempTime.Year(){
		return tempTime.Format("01月02日")
	}else{
		return tempTime.Format("2006年01月02日")
	}
}