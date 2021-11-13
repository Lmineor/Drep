package response

import "github.com/drep/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
