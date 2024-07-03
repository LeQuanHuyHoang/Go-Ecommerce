package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	r := InitRouter()

	r.Run(":8080")
}
