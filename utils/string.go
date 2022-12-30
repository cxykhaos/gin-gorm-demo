package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func Float64toA(a float64) string {
	return strconv.FormatFloat(a, 'f', 2, 64)
}
func AtoFloat64(a string) float64 {
	t, _ := strconv.ParseFloat(a, 64)
	return float64(t)
}

func AtoUin32(a string) uint32 {
	t, _ := strconv.ParseUint(a, 10, 32)
	return uint32(t)
}

func Atoint64(a string) int64 {
	t, _ := strconv.ParseInt(a, 10, 64)
	return int64(t)
}

func Atoint(a string) int {
	t, _ := strconv.ParseInt(a, 10, 64)
	return int(t)
}

func Uint32ToA(u uint32) string {
	return strconv.FormatUint(uint64(u), 10)
}
func Int64ToA(u int64) string {
	return strconv.FormatInt(int64(u), 10)
}

func StructToJson[T any](a T) string {
	res, _ := json.Marshal(a)
	return string(res)
}

func JsonToStruct[T any](as string) (T, error) {
	//var a map[string]interface{}
	a := new(T)
	err := json.Unmarshal([]byte(as), a)
	return *a, err
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func RandaomString(n int) string {
	var letters = []byte("eKEzxcBvWCUfF9ilOopL51kXYG7gSrtyhNj4mVRnb2suJD3AdQvIvTcbZ8aM0PqH")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RandomName() string {
	name := make([]byte, 20)
	i := 0
	for i < 20 {
		//利用rand.Intn(x)来伪随机生成一个[0,x)的数
		a := rand.Intn(26)
		b := rand.Intn(2)
		if b == 1 {
			name[i] = byte('A' + a)
		} else {
			name[i] = byte('a' + a)
		}
		i++
	}
	return string(name)
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5Encrypt(start string) string {
	h := md5.New()
	h.Write([]byte(start))
	return hex.EncodeToString(h.Sum(nil))
}

func Generate_uuid() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}
