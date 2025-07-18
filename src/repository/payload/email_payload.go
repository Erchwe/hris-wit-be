package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
)

type SendReportEmailPayload struct {
	RecipientEmail string `json:"recipient_email" valid:"required,email"`
	Subject        string `json:"subject" valid:"required"`
	SenderName     string `json:"sender_name" valid:"required"`
	BodyHTML       string `json:"body_html" valid:"required"`
}

func (payload *SendReportEmailPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}
