//go:build !dev && !prod && !citest
// +build !dev,!prod,!citest

package configs

import _ "embed"

func init() {
	Active = Local
}

//go:embed local.yaml
var CfgFile []byte
