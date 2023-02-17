package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/e"
)

type ChangeVoteForm struct {
	State int `form:"state" valid:"Range(0,1)"`
}

// Change godoc
// @Summary 管理员控制选择开始和结束
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Param state body int true "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote [post]
func Change(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ChangeVoteForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
