package foundation

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/YWJSonic/ServerUtility/code"
	"github.com/YWJSonic/ServerUtility/messagehandle"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// DeleteArrayElement ...
func DeleteArrayElement(elementIndex interface{}, array []interface{}) []interface{} {
	count := len(array)
	for index := 0; index < count; index++ {
		if elementIndex == array[index] {
			return append(array[:index], array[index+1:]...)
		}
	}
	return array
}

// ToJSONStr Convert to json string
func ToJSONStr(data interface{}) []byte {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return jsonString
}

// JSONToString conver JsonStruct to JsonString
func JSONToString(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", " ")
	STR := string(data)
	STR = strings.ReplaceAll(STR, string(10), ``)
	return STR
}

// StringToJSON ...
func StringToJSON(jsStr string) map[string]interface{} {
	return ByteToJSON([]byte(jsStr))
}

// ByteToJSON ...
func ByteToJSON(jsByte []byte) map[string]interface{} {
	var data map[string]interface{}
	if errMsg := json.Unmarshal(jsByte, &data); errMsg != nil {
		panic(errMsg)
	}

	return data
}

// InterfaceTofloat64 ...
func InterfaceTofloat64(v interface{}) float64 {
	return v.(float64)
}

// InterfaceToInt ...
func InterfaceToInt(v interface{}) int {
	switch v.(type) {
	case float64:
		return int(InterfaceTofloat64(v))
	case int:
		return v.(int)
	case int64:
		return int(v.(int64))
	default:
		panic("Conver Error")
	}
}

// InterfaceToInt64 ...
func InterfaceToInt64(v interface{}) int64 {
	switch v.(type) {
	case float64:
		return int64(v.(float64))
	case int:
		return int64(v.(int))
	case int64:
		return v.(int64)
	default:
		messagehandle.ErrorLogPrintln("Conver", v)
		panic("Conver Error")
	}
}

// InterfaceToBool ...
func InterfaceToBool(v interface{}) bool {
	switch v.(type) {
	case int:
		return v.(bool)
	case bool:
		return v.(bool)
	default:
		panic("Conver Error")
	}
}

// InterfaceToDynamicInt ...
func InterfaceToDynamicInt(v interface{}) code.Code {
	return code.Code(InterfaceTofloat64(v))
}

// InterfaceToString ...
func InterfaceToString(v interface{}) string {
	return v.(string)
}

// MD5Code encode MD5
func MD5Code(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// RandomMutily ...
func RandomMutily(rangeInt []int, pickCount int) []int {
	var result []int
	var targetPoint int

	if pickCount >= len(rangeInt) {
		return rangeInt
	}

	for i, imax := 0, pickCount; i < imax; i++ {
		targetPoint = rand.Intn(len(rangeInt))
		result = append(result, rangeInt[targetPoint])
		if len(rangeInt) > 1 {
			rangeInt = append(rangeInt[:targetPoint], rangeInt[targetPoint+1:]...)
		}
	}
	return result
}

// RangeRandomInt64 array random index
func RangeRandomInt64(rangeInt []int64) int {
	var Sum int64

	for _, value := range rangeInt {
		Sum += value
	}

	random := rand.Int63n(Sum)

	Sum = 0
	for i, value := range rangeInt {
		Sum += value
		if Sum > random {
			return i
		}
	}
	return -1
}

// RangeRandom array random index
func RangeRandom(rangeInt []int) int {
	Sum := 0

	for _, value := range rangeInt {
		Sum += value
	}

	random := rand.Intn(Sum)

	Sum = 0
	for i, value := range rangeInt {
		Sum += value
		if Sum > random {
			return i
		}
	}
	return -1
}

// ConevrToTimeInt64 Get time point
func ConevrToTimeInt64(year int, month time.Month, day, hour, min, sec, nsec int) int64 {
	return time.Date(year, month, day, hour, min, sec, nsec, time.Local).Unix()
}

// MapToArray return map keys and values
func MapToArray(mapData map[int64]int64) (keys, values []int64) {
	for key, value := range mapData {
		keys = append(keys, key)
		values = append(values, value)
	}
	return
}

// AppendMap map append map
func AppendMap(Target map[string]interface{}, Source map[string]interface{}) map[string]interface{} {
	for Key, Value := range Source {
		Target[Key] = Value
	}
	return Target
}

// ArrayShift Array Type []map[string]interface{}
func ArrayShift(Target []map[string]interface{}) (map[string]interface{}, []map[string]interface{}) {

	var out map[string]interface{}
	out = Target[0]
	Target = Target[1:]

	return out, Target
}

// CopyArray new array memory array
func CopyArray(source []int) []int {
	result := make([]int, len(source))
	copy(result, source)
	return result
}
