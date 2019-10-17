package konwledge

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	//"gotravel/prepare"
	"log"
)

const (
	host     = "192.168.128.147"
	port     = 5432
	user     = "admin"
	password = "123456"
	dbName   = "my_db"
	schema   = "marcopolo_gotravel"
)

//func main() {
//
//	user := &UserTbl{
//		Id:       1,
//		Username: "Windows",
//		Sex:      1,
//		Info:     "操作系统",
//	}
//
//	SessionUserTest(user)
//}

func getDBEngine() *xorm.Engine {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbName)
	//格式
	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL() //菜鸟必备

	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")
	return engine
}

//table name 为user_tbl
type UserTbl struct {
	Id       int
	Username string
	Sex      int
	Info     string
}

//func SelectByCondion(sql string,pojo interface{}) interface{}{
//	var obj []prepare.DestinationReq
//	engine := getDBEngine()
//	i := &pojo
//	engine.SQL(sql).Find(&i)
//	return i
//}

//查询所有
func selectAll() {
	var user []UserTbl
	engine := getDBEngine()
	engine.SQL("select * from user_tbl").Find(&user)
	fmt.Println(user)
}

//条件查询
func selectUser(name string) {
	var user []UserTbl
	engine := getDBEngine()
	engine.Where("user_tbl.username=?", name).Find(&user)
	fmt.Println(user)
}

//可以用Get查询单个元素
func selectOne(id int) {
	var user UserTbl
	engine := getDBEngine()
	engine.Id(id).Get(&user)
	//engine.Alias("u").Where("u.id=?",id).Get(&user)
	fmt.Println(user)
}

//添加
func InsertUser(user *UserTbl) bool {
	engine := getDBEngine()
	rows, err := engine.Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//删除(根据名称删除)
func DeleteUser(name string) bool {
	user := UserTbl{
		Username: name,
	}
	engine := getDBEngine()
	rows, err := engine.Delete(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//利用sql删除
func DeleteUserBySQL(name string) bool {
	engine := getDBEngine()
	result, err := engine.Exec("delete from user_tbl where username=?", name)
	if err != nil {
		log.Println(err)
		return false
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true
	}
	return false
}

//更新
func UpdateUser(user *UserTbl) bool {
	engine := getDBEngine()
	//Update(bean interface{}, condiBeans ...interface{}) bean是需要更新的bean,condiBeans是条件
	rows, err := engine.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		log.Println(err)
		return false
	}
	if rows > 0 {
		return true
	}
	return false
}

//利用session进行增删改
//用session的好处就是可以事务处理
func SessionUserTest(user *UserTbl) {
	engine := getDBEngine()
	session := engine.NewSession()
	session.Begin()
	_, err := session.Insert(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	user.Username = "windows"
	_, err = session.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	//_, err = session.Delete(user)
	//if err != nil {
	//	session.Rollback()
	//	log.Fatal(err)
	//}

	err = session.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
