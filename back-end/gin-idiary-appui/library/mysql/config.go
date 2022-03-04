/*
 * @Author: liziwei01
 * @Date: 2022-03-04 15:42:58
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-04 16:02:20
 * @Description: file content
 */
package mysql

// Config 配置
type Config struct {
	Username string // 账号名
	Password string // 密码
	DbName   string // 数据库名称
	DbDriver string // 驱动名称
	Host     string
	Port     string
}

func DefaultDbConf() *Config {
	return &Config{
		Username: "root",
		Password: "123456",
		DbName:   "db_",
		DbDriver: "mysql",
		Host:     "localhost",
		Port:     "3306",
	}
}
