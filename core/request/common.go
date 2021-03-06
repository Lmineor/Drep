package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

// GetByUUID Find by uuid structure
type GetByUUID struct {
	UUID string `json:"uuid" form:"uuid"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type UuidsReq struct {
	Uuids []string `json:"uuids" form:"uuids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string // 角色ID
}

type Empty struct{}
