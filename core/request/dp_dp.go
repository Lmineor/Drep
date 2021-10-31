package request

type Dp struct {
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	ProjectID uint   `json:"project_id" form:"project_id"`
}
