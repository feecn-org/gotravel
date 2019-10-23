package konwledge

import (
	"github.com/wonderivan/logger"
	"testing"
)

func TestInsterBody(t *testing.T) {
	user := &UserTbl{
		Id:       1,
		Username: "Windows",
		Sex:      1,
		Info:     "操作系统",
	}

	SessionUserTest(user)

	selectAll()
}

func TestOxm(t *testing.T) {
	mark := &Mark{
		Id:   1,
		Body: "11",
	}
	engine := getDBEngine()
	session := engine.NewSession()
	session.Begin()
	_, err := session.Insert(mark)
	logger.Debug(err)
	err = session.Commit()
	if err != nil {
	}
	//InsertMark(mark)
}
