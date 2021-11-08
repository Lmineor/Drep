package request

type InviteCode struct {
	ID          float64 `json:"id"`
	InviteCode  string  `json:"invite_code"`
	AuthorityId string  `json:"authority_id"`
	Description string  `json:"description"`
}
