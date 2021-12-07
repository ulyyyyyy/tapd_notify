package ginresp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strconv"
	"strings"
)

var mapper = map[AppCode]string{}

// ParseAndStoreCodeMsg 解析并存储 AppCode 的 `值` 与 `注释`
func ParseAndStoreCodeMsg() error {
	f, err := parser.ParseFile(token.NewFileSet(), "", goFileAppCode, parser.ParseComments)
	if err != nil {
		return err
	}

	baseNum := 0
	for _, dd := range f.Decls {
		gd := dd.(*ast.GenDecl)
		if gd.Tok != token.CONST {
			continue
		}

		counter := 0
		for _, sp := range gd.Specs {
			vs := sp.(*ast.ValueSpec)
			if len(vs.Values) != 0 {
				// iota + N
				// 获得 N 作为本组常量的起始数
				strExpr := types.ExprString(vs.Values[0])
				strNum := strings.TrimPrefix(strExpr, "iota + ")
				bi, err := strconv.Atoi(strNum)
				if err != nil {
					return err
				}
				baseNum = bi
			}
			if vs.Names[0].Name == "_" {
				counter++
				continue
			}
			mapper[AppCode(baseNum+counter)] = strings.TrimSpace(vs.Comment.Text())
			counter++
		}
	}
	return nil
}
