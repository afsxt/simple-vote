package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/afsxt/simple-vote/pkg/app"
	"github.com/afsxt/simple-vote/pkg/e"
	"github.com/afsxt/simple-vote/service/user_service"
	"github.com/afsxt/simple-vote/service/vote_service"
)

type VerifyUserForm struct {
	Email  string `form:"email" valid:"Required"`
	IDCard string `form:"idCard" valid:"Required"`
}

// Change godoc
// @Summary 用户较验
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @param idCard body string true "IDCard"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/vote/verify [post]
func VerifyUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form VerifyUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Email:  form.Email,
		IDCard: form.IDCard,
	}

	valid, err := userService.CheckValid()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_USER_VALID_FAIL, nil)
		return
	}
	if !valid {
		appG.Response(http.StatusOK, e.ERROR_INVALID_USER, nil)
		return
	}

	userService.Verify = 1
	if err := userService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// Change godoc
// @Summary 用户获取选举状态
// @Schemes
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param themeID body string true "ThemeID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/vote/theme/{themeID} [get]
func GetVoteDetails(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	voteService := vote_service.Vote{
		ThemeID: com.StrTo(c.Param("themeID")).MustInt(),
	}

	votes, err := voteService.GetVote()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VOTE_DETAIL_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = votes

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
