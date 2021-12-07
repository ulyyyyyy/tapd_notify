package configs

const (
	Local       = "local.yaml"  // 本地开发
	Development = "dev.yaml"    // 测试服
	CiTest      = "citest.yaml" // 云上开发服
	Production  = "prod.yaml"   // 正式服
)

// Active 当前使用的配置文件,
// 由 条件编译 `tags` 进行赋值
var Active string
