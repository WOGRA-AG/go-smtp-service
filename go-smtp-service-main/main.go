package main

import (
	"wogra.com/smtpwebservice"

	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("smtp-service-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	log.Println("Starting smtp-service")

	router := gin.Default()
	log.Println("Adding routes")
	router.POST("/sendTestMail/", smtpwebservice.SendTestMail)
	router.POST("/sendMail/", smtpwebservice.SendMail)

	log.Println("Run router")
	router.Run("localhost:8080")
}
