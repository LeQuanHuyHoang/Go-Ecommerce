package initialize

import (
	"github.com/LeQuanHuyHoang/Go-Ecommerce/global"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
