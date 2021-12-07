package model

import (
	"database/sql/driver"
	"encoding/json"
	_ "encoding/json"
	"fmt"
)

type JsonListParser []JsonParser

type JsonParser struct {
	Field string `json:"name"`
	Value string `json:"value"`
}

// Value 保存数据，序列化
func (args JsonListParser) Value() (driver.Value, error) {
	if args == nil {
		return nil, nil
	}
	marshal, err := json.Marshal(args)
	return string(marshal), err
}

// Scan 数据库数据转model
func (args *JsonListParser) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	err := json.Unmarshal(b, &args)
	return err
}
