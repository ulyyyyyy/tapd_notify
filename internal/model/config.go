// Package model
package model

import (
	"github.com/ulyyyyyy/tapd_notify/internal/helper/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

const (
	_tableName       = "tapd_conf"
	_mapperTableName = "field_mapper"
)

//
// 注意，这里由于存在两个关键key(project_id & platform_key)，所以这里需要了解逻辑。
// 部分数据是需要向前端传可读性高的数据，部分数据是需要向后端自用id类数据（例如字段的name 与id）
// 所以表现为，前后端传的数据不同

var (
	_ schema.Tabler = (*WebhookCfg)(nil)
)

type WebhookCfg struct {
	gorm.Model

	Name    string `json:"webhook_name" gorm:"size:10;column:webhook_name;not null"` // 配置名
	Creator string `json:"creator" gorm:"type:varchar(30);column:creator;not null"`  // 配置创建人
	Updater string `json:"updater" gorm:"type:varchar(30);column:updater"`           // 配置更新人
	Comment string `json:"comment" gorm:"size:255;column:comment"`                   // 配置备注
	Status  bool   `json:"status" gorm:"size:1;column:status"`                       // 配置状态

	ProjectId      string         `json:"project_id" gorm:"size:30;column:project_id;not null"`                // 推送项目
	MatchCondition JsonListParser `json:"match_condition" gorm:"type:varchar;size:600;column:match_condition"` // 推送匹配条件
	SummaryField   StrList        `json:"summary_field" gorm:"type:varchar;size:30;column:summary_field"`      // 概要字段
	PushField      StrList        `json:"push_field" gorm:"type:varchar;size:30;column:push_field"`            // 推送字段
	PushList       JsonListParser `json:"push_list" gorm:"type:varchar;size:1000;column:push_list"`            // 推送人列表
}

func (cfg *WebhookCfg) TableName() string {
	return _tableName
}

// Insert 插入一条配置
func Insert(wbCfg *WebhookCfg, tx *gorm.DB) error {
	err := mysql.Insert(wbCfg, tx)
	return err
}

// FindById 根据Id搜索具体配置
func FindById(id string) (wbCfg WebhookCfg, err error) {
	err = mysql.FindByID(id, &wbCfg)
	return
}

func FindByProjectAndId(project, id string) (wbCfg WebhookCfg, err error) {
	filter := map[string]interface{}{
		"id":         id,
		"project_id": project,
	}
	err = mysql.FindOneBy(filter, &wbCfg, mysql.DB())
	return
}

// FindByProjectId 根据项目组Id获取到所有基本信息
func FindByProjectId(projectId string) (wbCfgList []WebhookCfg, err error) {
	err = mysql.DB().Table(_tableName).Where("project_id", projectId).Order("id ASC").Find(&wbCfgList).Error
	return
}

// FindPageByProjectId 根据项目组Id获取到所有基本信息
func FindPageByProjectId(projectId string, pageNumber, pageSize int) (wbCfgList []WebhookCfg, err error) {
	err = mysql.DB().Table(_tableName).Where("project_id", projectId).Order("id ASC").Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&wbCfgList).Error
	return
}

// CountByProjectId 统计项目组下所有配置数量，用以前端分页
func CountByProjectId(projectId string) (total int64, err error) {
	err = mysql.DB().Model(&WebhookCfg{}).Where("project_id", projectId).Count(&total).Error
	return total, err
}

// FindDetailById 根据id获取全部信息
func FindDetailById(id int) (wbCfg WebhookCfg, err error) {
	// 所有
	err = mysql.DB().Table(_tableName).Where("id", id).Order("id ASC").Find(&wbCfg).Error
	summaryField := wbCfg.SummaryField
	pushField := wbCfg.PushField

	// 处理
	var summaryFieldName, pushFieldName []string
	db := mysql.DB()
	err = db.Table(_mapperTableName).Select("field_name").Where("field_id IN ?", summaryField).Scan(&summaryFieldName).Error
	if err != nil {
		return WebhookCfg{}, err
	}
	err = db.Table(_mapperTableName).Select("field_name").Where("field_id IN ?", pushField).Scan(&pushFieldName).Error
	if err != nil {
		return WebhookCfg{}, err
	}
	err = wbCfg.SummaryField.Scan(summaryFieldName)
	if err != nil {
		return WebhookCfg{}, err
	}
	return
}

// UpdateStatus 更新配置开启状态
func (cfg *WebhookCfg) UpdateStatus() error {
	_tx := mysql.DB()
	return _tx.Model(cfg).Update("status", !cfg.Status).Error
}

// UpdateById 根据配置id更新配置内容
func UpdateById(cfg *WebhookCfg, updater string) error {
	tx := mysql.DB().Begin()

	(*cfg).Updater = updater
	(*cfg).UpdatedAt = time.Now()
	err := mysql.Save(*cfg, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	return err
}

// Delete 根据配置id删除配置内容
func (cfg *WebhookCfg) Delete(tx *gorm.DB) error {
	return mysql.Delete(cfg, tx)
}

// DeleteById 根据配置id删除配置内容
func DeleteById(id string) (err error) {
	tx := mysql.DB().Begin()
	err = mysql.DeleteById(id, &WebhookCfg{}, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// ExitsByIdAndProject 检查项目下配置id是否存在
func ExitsByIdAndProject(id, project string) (exits bool, err error) {
	filter := map[string]interface{}{
		"id":         id,
		"project_id": project,
	}
	exits, err = mysql.ExistBy(filter, &WebhookCfg{}, mysql.DB())
	return
}
