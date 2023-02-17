package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/e"
	"github.com/afsxt/simple-vote/service/candidate_service"
	"github.com/afsxt/simple-vote/service/theme_service"
)

type AddCandidateForm struct {
	Name        string `form:"name" valid:"Required"`
	Description string `form:"description"`
	ThemeID     int    `form:"themeID" valid:"Required"`
}

// Change godoc
// @Summary 管理员创建候选人
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @param description body string false "Description"
// @param themeID body int true "ThemeID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/candidates [post]
func AddCandidates(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCandidateForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	log.Println(1111, form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	themeService := theme_service.Theme{
		ID: form.ThemeID,
	}
	exists, err := themeService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_THEME_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_THEME, nil)
		return
	}

	candidateService := candidate_service.Candidates{
		Name:        form.Name,
		Description: form.Description,
		ThemeID:     form.ThemeID,
	}

	exists, err = candidateService.ExistBy()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CANDIDATE_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_CANDIDATE, nil)
		return
	}

	if err := candidateService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CANDIDATE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
