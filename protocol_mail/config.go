package protocol_mail

// config for protocol mail processor.
type config struct {
	ProcessorsField string `config:"processors_field"`
	SourceField     string `config:"source_field"`
	TargetField     string `config:"target_field"`
	IgnoreMissing   bool   `config:"ignore_missing"`
	IgnoreFailure   bool   `config:"ignore_failure"`
}

func defaultConfig() config {
	return config{
		ProcessorsField: "processors.protocol_mail",
		SourceField:     "row_mail",
		TargetField:     "mail",
		IgnoreFailure:   true,
	}
}
