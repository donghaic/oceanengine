package project

import "github.com/bububa/oceanengine/marketing-api/model"

// UpdateResponse 更新计划API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// ProjectID 广告项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// PromotionIDs 广告项目ID集合
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的广告项目
type UpdateError struct {
	// ProjectID 广告项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ObjectType 错误对象类型
	// 枚举值： BASIC 广告基本设置、MATERIAL 广告素材组合、BUDGET 广告预算
	ObjectType string `json:"object_type,omitempty"`
	// ErrorCode 错误信息
	ErrorCode int `json:"error_code,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message"`
}