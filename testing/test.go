package testing

import (
	"github.com/afsxt/simple-vote/models"
	"github.com/afsxt/simple-vote/pkg/gredis"
	"github.com/afsxt/simple-vote/pkg/logging"
	"github.com/afsxt/simple-vote/pkg/setting"
)

func setup() {
	setting.Setup("../conf/app.ini")
	models.Setup()
	gredis.Setup()
	logging.Setup()
}

func cleanTable() {}
