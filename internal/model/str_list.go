package model

import (
	"database/sql/driver"
	"strings"
)

type strList []string

func (p strList) Value() (driver.Value, error) {
	if p == nil || len(p) == 0 {
		return []byte(""), nil
	}
	join := strings.Join(p, ",")
	return []byte(join), nil
}

func (p *strList) Scan(value interface{}) error {
	if string(value.([]byte)) == "" {
		tmp := strList{}
		*p = tmp
		return nil
	}
	split := strings.Split(string(value.([]byte)), ",")
	// 把p值塞进去
	tmp := strList(split)
	*p = tmp
	return nil
}
