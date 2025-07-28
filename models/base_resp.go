package models

// Base Pagination
type BasePagination struct {
	Page         int `json:"page"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total"`
	TotalPage    int `json:"total_page"`
}

// BaseResp represents a standard response structure
type BaseResp[T any] struct {
	Message  string `json:"message"`
	Data     T      `json:"data,omitempty"`
	DataList []T    `json:"data_list,omitempty"`
}
