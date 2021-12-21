package protocol_mail

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/jhillyerd/go.enmime"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("protocol_mail", New)
	jsprocessor.RegisterPlugin("ProtocolMail", New)
}

type ProtocolMail struct {
	config
	log *logp.Logger
}

const (
	processorName = "protocol_mail"
	logName       = "processor.protocol_mail"
)

// New constructs a new protocol_mail processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	p := &ProtocolMail{
		config: config,
		log:    logp.NewLogger(logName),
	}

	return p, nil
}

func (p *ProtocolMail) Run(event *beat.Event) (*beat.Event, error) {
	rawMail, err := event.GetValue(p.SourceField)
	if err != nil {
		return event, errors.Wrapf(err, "failed to get source field %s", p.SourceField)
	}

	mailMessage := strings.NewReader(rawMail)
	msg, _ := mail.ReadMessage(mailMessage)
	mime, _ := enmime.ParseMIMEBody(msg)

	event.PutValue("subject", mime.GetHeader("Subject"))

	return event, nil
}

func (p *ProtocolMail) String() string {
	return fmt.Sprintf("protocol_mail=[source_field=%s,target_field=%s]", p.SourceField, p.TargetField)
}
