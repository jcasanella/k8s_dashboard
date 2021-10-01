package main

import (
	"github.com/jcasanella/k8s_dashboard/router"
)

func main() {
	router := router.Router()
	router.Run("localhost:8080")
}
