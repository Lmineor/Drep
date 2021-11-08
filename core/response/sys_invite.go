package response

import "github.com/drep/model"

// InviteCodeResponse struct
type InviteCodeResponse struct {
	InviteCode *model.SysInvite `json:"invite_code"`
}
