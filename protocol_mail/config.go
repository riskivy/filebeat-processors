package protocol_mail

// config for protocol mail processor.
type config struct {
	ProcessorsField string `config:"processors_field"`
	SourceField     string `config:"source_field"`
}

func defaultConfig() config {
	return config{
		ProcessorsField: "processors.protocol_mail",
		SourceField:     "raw_mail",
	}
}
