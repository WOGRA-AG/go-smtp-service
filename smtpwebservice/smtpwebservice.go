package smtpwebservice

import (
	"net/http"

	"wogra.com/configReader"
	"wogra.com/smtpsender"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Mail struct {
	Sender   string                   `json:"sender"`
	Receiver []string                 `json:"receiver"`
	Subject  string                   `json:"subject"`
	Message  string                   `json:"message"`
	Token    configReader.AccessToken `json:"accesstoken"`
}

func SendTestMail(c *gin.Context) {

	log.Println("SendTestMail called")

	var token configReader.AccessToken
	err := c.BindJSON(&token)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if configReader.CheckAccessToken(&token) == true {

		from := "wg@digital-2.com"

		to := []string{
			"wg@wogra.com",
		}

		smtpsender.SmtpSender(from, to, "Testmail", "Email body")
		c.IndentedJSON(http.StatusOK, "Mail sent")
	} else {
		errorMsg := fmt.Sprintf("Accesstoken %v not allowed.", token)
		log.Println(errorMsg)
		c.IndentedJSON(http.StatusMethodNotAllowed, "Sending mail not allowed. "+errorMsg)
	}
}

func SendMail(c *gin.Context) {

	var mailData Mail
	err := c.BindJSON(&mailData)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	log.Println(mailData)

	if configReader.CheckAccessToken(&mailData.Token) == true {

		err = smtpsender.SmtpSender(mailData.Sender, mailData.Receiver, mailData.Subject, mailData.Message)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, err)
		} else {
			c.IndentedJSON(http.StatusOK, "Mail sent")
		}

	} else {
		errorMsg := fmt.Sprintf("Accesstoken %v not allowed.", mailData.Token)
		log.Println(errorMsg)
		c.IndentedJSON(http.StatusMethodNotAllowed, "Sending mail not allowed. "+errorMsg)
	}
}
