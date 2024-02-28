package utilities

import "log"

func CheckError(str string, err error) {
	if err != nil {
		log.Printf("%s:%s\n", str, err)
		panic(err)
	}
}
