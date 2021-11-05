package request

// User register structure
type InviteCode struct {
	Code        string `json:"code"`
	AuthorityId string `json:"authority_id"`
}
