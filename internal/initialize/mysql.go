package initialize

import (
	"fmt"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/global"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/pkg/setting"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func checkErrorPanic(err error, errMsg string) {
	if err == nil {
		global.Logger.Error(errMsg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%d)/%d?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	checkErrorPanic(err, "InitMySQL failed")

	global.Mdb = db

	// set Pool
	SetPool(m)
}

func SetPool(config setting.MySQLSetting) {
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("mysql error: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(config.MaxIdleConns))
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
}

func migrateTables() {

}
