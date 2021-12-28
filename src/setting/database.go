// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 20:40
// @Software: GoLand

package setting

import "api.openfileplatform.com/src/globals/database"

func InitDatabase() (err error) {
	err = database.InitMysqlClient()
	if err != nil {
		return
	}
	err = database.InitRedisClient()
	return
}
