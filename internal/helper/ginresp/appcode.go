package ginresp

// 由函数 `parseCodeMsg` 计算获得 `AppCode` 的值、注释

type AppCode int

const Success AppCode = iota + 200 // 操作成功

const (
	_ AppCode = iota + 210 // 配置相关错误
)

const (
	_ AppCode = iota + 220
)
