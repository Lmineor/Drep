package response

import "github.com/drep/core/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
