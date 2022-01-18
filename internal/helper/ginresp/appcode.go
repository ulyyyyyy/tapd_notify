package ginresp

// 由函数 `parseCodeMsg` 计算获得 `AppCode` 的值、注释

type AppCode int

const Success AppCode = iota + 200 // 操作成功

const (
	_              AppCode = iota + 201
	ErrReceiveBody         // 请求体数据有误
)

const (
	_                     AppCode = iota + 30100 // 配置相关错误 ( 配置增/删/改/查/.. )
	ErrConfigIdError                             // 配置ID错误
	ErrConfigNotExists                           // 配置不存在
	ErrRequestBodyParse                          // webhook数据Json解析失败
	ErrProjectId                                 // projectId错误
	ErrPageSize                                  // PageNumber 错误
	ErrPageNumber                                // PageNumber 错误
	ErrWebhookCfgId                              // webhook配置id错误
	ErrCfgFindByProjectId                        // 根据projectId查找webhook配置失败
	ErrCfgInsert                                 // webhook配置新增失败
	ErrCfgInsertField                            // webhook配置字段错误
	ErrCfgInsertDuplicate                        // 配置名重复，项目组下有相同配置
	ErrCfgUpdate                                 // webhook配置更新失败
	ErrCfgUpdateField                            // webhook配置字段错误
	ErrCfgDelete                                 // webhook配置删除失败

	ErrMessagePush // 消息推送失败
)

const (
	_                AppCode = iota + 220
	ErrProxyNotExits         // 转发服务器配置不存在
)
