package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/constants"
	"github.com/afsxt/simple-vote/pkg/e"
	"github.com/afsxt/simple-vote/service/candidate_service"
	"github.com/afsxt/simple-vote/service/theme_service"
)

type AddThemeForm struct {
	Name        string `form:"name" valid:"Required"`
	Description string `form:"description"`
}

// Change godoc
// @Summary 管理员创建选举主题
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @param description body string false "Description"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/theme [post]
func AddTheme(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddThemeForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	themeService := theme_service.Theme{
		Name:        form.Name,
		Description: form.Description,
	}

	exists, err := themeService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_THEME_EXIST_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_THEME_EXIST, nil)
		return
	}

	if err := themeService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_THEME_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type ChangeThemeStateForm struct {
	State int `form:"state" valid:"Required"`
}

// Change godoc
// @Summary 管理员控制选举主题开始或结束
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Param state body int true "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/theme/{id}/state [post]
func ChangeThemeState(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ChangeThemeStateForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	themeService := theme_service.Theme{
		ID:    com.StrTo(c.Param("id")).MustInt(),
		State: int(form.State),
	}

	exists, err := themeService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_THEME_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_THEME_NOT_EXIST, nil)
		return
	}

	// 打开选举时必须满足候选人有２个以上
	if form.State == int(constants.ThemeStarted) {
		c := candidate_service.Candidates{
			ThemeID: themeService.ID,
		}
		count, err := c.GetCount()
		if err != nil {
			appG.Response(http.StatusOK, e.ERROR_THEME_GET_CANDIDATE_COUNT_FAIL, nil)
			return
		}
		if count < 2 {
			appG.Response(http.StatusOK, e.ERROR_THEME_COUNT_NOT_ENOUGH, nil)
			return
		}
	}
	if err := themeService.ChangeState(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_THEME_ADD_FAIL, nil)
		return
	}

	// 当关闭选举时发送异步任务发送邮箱
	if form.State == int(constants.ThemeFinished) {
		// todo:
		// utils.SendMail()
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// Change godoc
// @Summary 管理员更新选举主题
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/theme [put]
func UpdateTheme(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
