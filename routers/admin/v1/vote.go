package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/e"
	"github.com/afsxt/simple-vote/pkg/setting"
	"github.com/afsxt/simple-vote/pkg/util"
	"github.com/afsxt/simple-vote/service/vote_service"
)

// Change godoc
// @Summary 管理员获取选举结果，如有candidate参数即为该候选人的票数，否则为该主题下所有候选人得票结果
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Param theme body int false "Theme"
// @Param candidate body int false "Candidate"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/theme/{themeID} [get]
func GetThemeResult(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	voteService := vote_service.Vote{
		ThemeID:     com.StrTo(c.Param("themeID")).MustInt(),
		CandidateID: com.StrTo(c.Query("candidate")).MustInt(),
	}

	votes, err := voteService.GetVote()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_VOTE_GET_DETAIL_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = votes

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// Change godoc
// @Summary 管理员获取某个候选人的支持用户
// @Schemes
// @Description
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/vote/theme/{themeID}/candidate/{candidateID} [get]
func GetCandidateUsers(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	voteService := vote_service.Vote{
		ThemeID:     com.StrTo(c.Param("themeID")).MustInt(),
		CandidateID: com.StrTo(c.Param("candidateID")).MustInt(),
		PageNum:     util.GetPage(c),
		PageSize:    setting.AppSetting.PageSize,
	}

	votes, err := voteService.GetVoteUsers()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_GET_VOTE_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = votes

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
