package error

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
