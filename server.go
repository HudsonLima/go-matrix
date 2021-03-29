package main

import (
	"net/http"
	router "github.com.com/HudsonLima/GoMatrix/http"
	"github.com.com/HudsonLima/GoMatrix/service"
	"github.com.com/HudsonLima/GoMatrix/controller"
	"fmt"
)

var (
	matrixService    service.MatrixService       = service.NewMatrixService()
	matrixController controller.MatrixController = controller.NewMatrixController(matrixService)
	httpRouter       router.Router               = router.NewChiRouter()
)

func main() {
	handleRequests(httpRouter)	
}

func handleRequests(httpRouter router.Router) {
	const port string = ":8000"	

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
		fmt.Fprintf(response, "\nThe avaialable api endpoints are")
		fmt.Fprintf(response, "\n/api/echo")
		fmt.Fprintf(response, "\n/api/invert")
		fmt.Fprintf(response, "\n/api/flatten")
		fmt.Fprintf(response, "\n/api/sum")
		fmt.Fprintf(response, "\n/api/multiply")
	})

	httpRouter.POST("/api/echo", matrixController.MatrixEcho)
	httpRouter.POST("/api/invert", matrixController.MatrixInvert)
	httpRouter.POST("/api/flatten", matrixController.MatrixFlatten)
	httpRouter.POST("/api/sum", matrixController.MatrixSum)
	httpRouter.POST("/api/multiply", matrixController.MatrixMultiply)

	httpRouter.SERVE(port)
}