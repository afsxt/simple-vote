package v1

import (
	"net/http"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/e"
	"github.com/afsxt/simple-vote/service/theme_service"
	"github.com/afsxt/simple-vote/service/vote_service"
	"github.com/gin-gonic/gin"
)

type VoteForm struct {
	ThemeID     int `form:"themeID" valid:"Required"`
	UserID      int `form:"userID" valid:"Required"`
	CandidateID int `form:"candidateID" valid:"Required"`
}

// Change godoc
// @Summary 用户投票
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param themeID body string true "ThemeID"
// @param userID body string true "UserID"
// @param candidateID body string true "CandidateID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/vote [post]
func Vote(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form VoteForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	themeService := theme_service.Theme{
		ID: form.ThemeID,
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
	// todo: 需较验其他id正确性，先忽略

	voteService := vote_service.Vote{
		ThemeID:     form.ThemeID,
		UserID:      form.UserID,
		CandidateID: form.CandidateID,
	}

	if exists, _ := voteService.Check(); exists {
		appG.Response(http.StatusOK, e.ERROR_VOTE_AGAIN_FAILE, nil)
		return
	}

	if err := voteService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_VOTE_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
