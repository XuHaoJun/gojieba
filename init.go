package gojieba

import (
	"github.com/xuhaojun/gojieba/deps/cppjieba"
	"github.com/xuhaojun/gojieba/deps/limonp"
	"github.com/xuhaojun/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
