package ping

import "time"

func checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "Pong"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}
