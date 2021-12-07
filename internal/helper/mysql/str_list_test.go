package mysql

import (
	"log"
	"testing"
)

func TestStrList_Scan(t *testing.T) {

	list := StrList{"ff"}
	_ = (&list).Scan([]byte("测试,测试"))

	log.Println(list)
}
