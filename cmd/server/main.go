package main

import "github.com/LeQuanHuyHoang/Go-Ecommerce/internal/router"

func main() {
	r := router.NewRouter()

	r.Run(":8080")
}
