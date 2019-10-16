package konwledge

import "testing"

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
