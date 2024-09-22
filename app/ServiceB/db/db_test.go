package db

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/lib/pq"
	"testing"
)

func Test_a(t *testing.T) {
	res := GetById(1)
	g.Dump(res)
}
