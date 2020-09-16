package timeutils

import "time"

func getNowWithoutDelimiter()(now string){
	return time.Now().Format("20060102150405000")
}
