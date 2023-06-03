package main

import (
	"github.com/kerimcetinbas/golang/router"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

func main() {

	router.StartRouter(":8080")
}
