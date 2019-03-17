package globals

import (
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"unsafe"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

/**
 * 随机字符串
 */
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

/**
 * 生成字session
 * 随机字符串+时间
 */
func GenerateSession() string {
	h := md5.New()
	randString := Krand(16, KC_RAND_KIND_ALL)
	timeString := time.Now().String()

	h.Write(randString)
	h.Write(Str2bytes(timeString))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func MakeMd5(strByte []byte) string {
	h := md5.New()
	h.Write(strByte)
	return hex.EncodeToString(h.Sum(nil))
}

func MakeMd5FromString(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}



