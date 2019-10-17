package konwledge

import (
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
	//sql := "select * from user_tbl"
	//var user []UserTbl
	//condion := SelectByCondion(sql, user)
	//logger.Debug(condion)
}
