package main

import (
	"github.com/gin-gonic/gin"
	"wogra.com/configReader"
	"wogra.com/smtpwebservice"

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

	restConfig := configReader.ReadRestConfiguration()

	connection := restConfig.Host + ":" + restConfig.Port

	router.Run(connection)
}
