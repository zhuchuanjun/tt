package zconv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"io"
	"os"
)

func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.Marshal(v)
	if err != nil {
		return nil, gerror.Wrap(err, "json marshal failed")
	}
	return
}

func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	marshaledBytes, err = json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return nil, gerror.Wrap(err, "json marshal indent failed")
	}
	return
}

func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, &v)
	if err != nil {
		return gerror.Wrap(err, "json unmarshal failed")
	}
	return
}

func UnmarshalUseNumber(data []byte, v interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err = decoder.Decode(v)
	if err != nil {
		return gerror.Wrap(err, "json.UnmarshalUseNumber failed")
	}
	return
}

func NewEncoder(i io.Writer) *json.Encoder {
	return json.NewEncoder(i)
}

func NewDecoder(i io.Reader) *json.Decoder {
	return json.NewDecoder(i)
}

func Valid(data []byte) bool {
	return json.Valid(data)
}

func Print() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}
