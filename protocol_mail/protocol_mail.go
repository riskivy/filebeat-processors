package protocol_mail

import (
	"fmt"

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
	_, err := event.PutValue(p.TargetField, "test")
	if err != nil {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", p.TargetField, "test")
	}

	return event, nil
}

func (p *ProtocolMail) String() string {
	return fmt.Sprintf("protocol_mail=[source_field=%s,target_field=%s]", p.SourceField, p.TargetField)
}
