package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

// ResponseXls 向前端返回Excel文件
// 参数 content 为上面生成的io.ReadSeeker， fileTag 为返回前端的文件名
func ResponseXls(c *gin.Context, content io.ReadSeeker, fileTag string) {
	fileName := fmt.Sprintf("%s%s%s.xlsx", time.Now(), `-`, fileTag)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
}
