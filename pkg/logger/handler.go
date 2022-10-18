package logger

func HandleError(err error, fatal bool) bool {
	if err != nil {
		Error(err)
		if fatal {
			panic(err)
		}

		return true
	}

	return false
}
