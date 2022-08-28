package util

import (
	"math/rand"
	"time"
)

// Convert Snake("CreateUser") style srtings to like this "Create-User"
func SnakeWithChar(s string, chr ...string) string {
	ret := ""
	if len(chr) == 0 {
		chr = make([]string, 1)
		chr[0] = "-"
	}
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 65 && s[i] <= 90 {
			ret = ret + chr[0] + string(s[i])
		} else {
			ret = ret + string(s[i])
		}
	}
	return ret
}

func RandomKey() string {
	rand.Seed(time.Now().Unix())
	key := [256]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	return string(key[:])
}
