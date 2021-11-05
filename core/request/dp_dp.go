package request

type Dp struct {
	Title       string `json:"title" form:"title"`
	Content     string `json:"content" form:"content"`
	ProjectUUID string `json:"project_uuid" form:"project_uuid"`
}
