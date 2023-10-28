package router

import (
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	// Roda o servidor em uma goroutine para que o código seguinte não seja bloqueado
	go func() {
		router.Run(":8080")
	}()

	// Dá uma pausa curta para garantir que o servidor esteja rodando
	time.Sleep(2 * time.Second)

	// Abre o navegador
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "http://localhost:8080/swagger/index.html")
	case "darwin":
		cmd = exec.Command("open", "http://localhost:8080/swagger/index.html")
	default: // linux
		cmd = exec.Command("xdg-open", "http://localhost:8080/swagger/index.html")
	}
	err := cmd.Start()
	if err != nil {
		log.Println("erro ao abrir o navegador:", err)
	}

	select {} // bloqueia indefinidamente, para que a main goroutine não termine
}
