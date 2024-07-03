package global

import (
	"github.com/LeQuanHuyHoang/Go-Ecommerce/pkg/logger"
	"github.com/LeQuanHuyHoang/Go-Ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
