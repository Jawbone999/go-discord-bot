package utils

// HandleError formats a context message with an error message
func HandleError(msg string, err error) {
	Logger.Fatal().Msg(msg + ": " + err.Error())
	// Logger.Fatal().Stack().Msg(msg + ": " + err.Error())
}
