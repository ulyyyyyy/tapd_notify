package tapd

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/mysql"
	"github.com/ulyyyyyy/tapd_notify/internal/model/webhook_cfg"
	"strconv"
	"strings"
	"time"
)

// GetConfigByProject 根据项目组获取配置情况
// @Tags 配置相关接口
// @Summary 查询项目组下配置简单信息
// @Accept application/json
// @Produce application/json
// @Router /tapd/configs/:project [get]
func GetConfigByProject(c *gin.Context) {
	projectId := c.Param("project")
	pageNumber, err := strconv.Atoi(c.Query("pageNumber"))
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrPageNumber, err.Error())
		return
	}
	if pageNumber < 1 {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrPageSize, err.Error())
		return
	}
	total, err := webhook_cfg.CountByProjectId(projectId)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrProjectId, err.Error())
		return
	}
	cfgList, err := webhook_cfg.FindByProjectId(projectId, pageNumber, pageSize)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrProjectId, err.Error())
		return
	}
	ginresp.NewSuccess(c, newConfigsResponse(total, cfgList))
}

//
// GetConfigById 根据Id获取单条配置情况
// @Tags 配置相关接口
// @Summary 查询单条配置完整信息
// @Accept application/json
// @Produce application/json
// @Router /tapd/config/:project/:id [get]
func GetConfigById(c *gin.Context) {
	project := c.Param("project")
	id := c.Param("id")

	cfg, err := webhook_cfg.FindByProjectAndId(project, id)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrWebhookCfgId, err.Error())
		return
	}
	ginresp.NewSuccess(c, cfg)
}

//
// CreateConfig 创建一条配置
// @Tags 配置相关接口
// @Summary 创建一条新的配置
// @Accept application/json
// @Produce application/json
// @Router /tapd/config/:project [post]
func CreateConfig(c *gin.Context) {
	project := c.Param("project")

	email := c.GetHeader("X-GateWay-Email")
	_tx := mysql.DB().Begin()

	var cfg webhook_cfg.WebhookCfg
	err := c.BindJSON(&cfg)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrRequestBodyParse, err.Error())
		return
	}
	if cfg.ProjectId != project {
		ginresp.NewFailure(c, ginresp.ErrCfgInsertField, nil)
		return
	}

	cfg.CreatedAt = time.Now()
	cfg.Creator = email
	err = webhook_cfg.Insert(&cfg, _tx)
	if err != nil {
		_tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			ginresp.NewFailure(c, ginresp.ErrCfgInsertDuplicate, nil)
		} else {
			ginresp.NewFailure(c, ginresp.ErrCfgInsert, err.Error())
		}
		return
	}
	_tx.Commit()
}

// UpdateConfigById 根据Id更新单条配置
// @Tags 配置相关接口
// @Summary 更新一条新的配置
// @Accept application/json
// @Produce application/json
// @Router /tapd/config/:project/:id [put]
func UpdateConfigById(c *gin.Context) {
	id := c.Param("id")
	project := c.Param("project")
	email := c.GetHeader("X-GateWay-Email")

	exits, err := webhook_cfg.ExitsByIdAndProject(id, project)
	if !exits || err != nil {
		ginresp.NewFailure(c, ginresp.ErrConfigNotExists, "配置不存在")
	}

	// 获取Body中配置
	var cfg webhook_cfg.WebhookCfg
	err = c.BindJSON(&cfg)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrRequestBodyParse, "配置传入失败")
		return
	}
	if project != cfg.ProjectId {
		ginresp.NewFailure(c, ginresp.ErrCfgUpdateField, "ProjectId 不匹配")
		return
	}

	// 更新操作
	err = webhook_cfg.UpdateById(&cfg, email)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrCfgInsert, err.Error())
		return
	}
	ginresp.NewSuccess(c, cfg)
}

// UpdateStatusById 根据ID调整一个配置的开关
// @Tags 配置相关接口
// @Summary 仅更新一条配置的状态
// @Accept application/json
// @Produce application/json
// @Router /tapd/config/status/:project/:id [put]
func UpdateStatusById(c *gin.Context) {
	id := c.Param("id")
	project := c.Param("project")
	email := c.GetHeader("X-GateWay-Email")

	cfg, err := webhook_cfg.FindByProjectAndId(project, id)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrConfigNotExists, "配置不存在")
	}

	cfg.Status = !cfg.Status
	err = webhook_cfg.UpdateById(&cfg, email)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrCfgUpdate, err.Error())
		return
	}
	ginresp.NewSuccess(c, cfg)
}

// DeleteConfigById 根据Id删除一条配置
// @Tags 配置相关接口
// @Summary 删除一条新的配置
// @Accept application/json
// @Produce application/json
// @Router /tapd/config/:project/:id [delete]
func DeleteConfigById(c *gin.Context) {
	id := c.Param("id")
	project := c.Param("project")

	//
	rlt, err := webhook_cfg.ExitsByIdAndProject(id, project)
	if !rlt {
		ginresp.NewFailure(c, ginresp.ErrConfigNotExists, id)
		return
	}

	err = webhook_cfg.DeleteById(id)
	if err != nil {
		ginresp.NewFailure(c, ginresp.ErrCfgDelete, id)
		return
	}
	ginresp.NewSuccess(c, id)
}

// configsResponse
type configsResponse struct {
	Total   int64                    `json:"total"`
	Configs []webhook_cfg.WebhookCfg `json:"rows"`
}

func newConfigsResponse(total int64, configs []webhook_cfg.WebhookCfg) *configsResponse {
	return &configsResponse{Total: total, Configs: configs}
}
