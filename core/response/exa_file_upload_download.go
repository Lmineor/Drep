package response

import "github.com/drep/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
