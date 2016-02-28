package bugtracker

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
)

// SendBug is called every time an unhandled error occurs, particularly when
// that error happens as a result of the message a user sends to Ava
func (sg *MailClient) SendBug(err error) {
	if os.Getenv("AVA_ENV") != "production" {
		return
	}
	subj := "[Bug] " + err.Error()
	if len(subj) > 30 {
		subj = subj[0:27] + "..."
	}
	text := "<html><body>"
	text += fmt.Sprintf("<p>%s</p>", err.Error())
	text += "</body></html>"
	if err := sg.Send(subj, text, Admin()); err != nil {
		log.Debug("could not send bug report", err)
	}
}
