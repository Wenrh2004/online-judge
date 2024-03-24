package dal

import "online-judge/biz/dal/mysql"

func InitDal() {
	mysql.InitDB()
}
