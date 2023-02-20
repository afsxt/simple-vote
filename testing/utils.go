package testing

import (
	"math/rand"
	"time"

	"github.com/afsxt/simple-vote/models"
	"github.com/afsxt/simple-vote/pkg/gredis"
	"github.com/afsxt/simple-vote/pkg/logging"
	"github.com/afsxt/simple-vote/pkg/setting"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	setup()
}

func setup() {
	setting.Setup("../conf/app.ini")
	models.Setup()
	gredis.Setup()
	logging.Setup()
}

func tearDown() {}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func resetThemeTable() {
	models.GetDB().Exec("truncate table vote_theme")
}

func resetCandidateTable() {
	models.GetDB().Exec("truncate table vote_candidates")
}

func resetUserTable() {
	models.GetDB().Exec("truncate table vote_user")
}

func resetVoteTable() {
	models.GetDB().Exec("truncate table vote_vote")
}
