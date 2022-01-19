package model

import (
	"database/sql/driver"
	"strings"
)

type StrList []string

func (p StrList) Value() (driver.Value, error) {
	if p == nil || len(p) == 0 {
		return []byte(""), nil
	}
	join := strings.Join(p, ",")
	return []byte(join), nil
}

func (p *StrList) Scan(value interface{}) error {
	if string(value.([]byte)) == "" {
		tmp := StrList{}
		*p = tmp
		return nil
	}
	split := strings.Split(string(value.([]byte)), ",")
	// 把p值塞进去
	tmp := StrList(split)
	*p = tmp
	return nil
}
