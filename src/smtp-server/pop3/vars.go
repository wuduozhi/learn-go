package pop3

var (
	// DefaultProcessors holds processor functions
	DefaultProcessors = map[string]Processor{
		"USER":userProcessor,
		"PASS":passProcessor,
		"STAT":statProcessor,
		"LIST":listProcessor,
		"RETR":retrProcessor,
	}
)