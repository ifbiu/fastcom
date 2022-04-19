package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNum(num int) string {
	if num==6 {
		return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	}else if num > 6 && num <= 12 {
		var resAll string
		res1 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		res2 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
        resAll = res1+res2[:num-6]
		return resAll
	}else{
		return "最少支持6位，最多支持12位随机数"
	}
}