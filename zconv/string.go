package zconv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"reflect"
	"strconv"
	"time"
)

func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case []byte:
		return string(value)
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.Itoa(int(value))
	case uint8:
		return strconv.Itoa(int(value))
	case uint16:
		return strconv.Itoa(int(value))
	case uint32:
		return strconv.Itoa(int(value))
	case uint64:
		return strconv.FormatUint(value, 10)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case gtime.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *gtime.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case bool:
		return strconv.FormatBool(value)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case string:
		return value
	default:
		if value == nil {
			return ""
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		switch kind {
		case reflect.Array,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.UnsafePointer,
			reflect.Chan,
			reflect.Map:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Actors []string `json:"actors"`
}

func main() {
	//m := Movie{
	//	Title:  "Inception",
	//	Year:   2010,
	//	Actors: []string{"Leonardo DiCaprio", "Ellen Page", "Tom Hardy"},
	//}
	//
	//// 使用MarshalIndent进行格式化编码
	//jsonStr, err := json.MarshalIndent(m, "", "    ")
	//if err != nil {
	//	fmt.Println("Error marshalling JSON:", err)
	//	return
	//}

	//g.Dump(jsonStr)
	//w := bytes.NewReader([]byte{'a'})
	//json.NewDecoder(w).UseNumber()

	a := `{"name":"zhuzhu", "age":18}`
	g.DumpJson(a)
	//fmt.Println(a)
	DumpJson(a)
}

func DumpJson(jsonContent string) {
	buffer := bytes.NewBuffer(nil)
	err := json.Indent(buffer, []byte(jsonContent), "", "\t")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(buffer.String())
}
