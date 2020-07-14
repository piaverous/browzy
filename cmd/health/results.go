package health

import (
	"strconv"
	"time"
)

var HISTORY_SIZE int = 60

type ResultsMaster struct {
	in            chan HealthCheck
	results       []HealthCheck
	statusCodes   map[int]int
	responseTimes []float64
}

func (rm *ResultsMaster) consumeResults() {
	for {
		healthcheck, open := <-rm.in
		if open == false {
			break
		}
		rm.results = append(rm.results, healthcheck)
		rm.addStatusCode(healthcheck.status)
		rm.responseTimes = append(
			rm.responseTimes,
			float64(healthcheck.responseTime/time.Millisecond),
		)

		rm.truncateArrays(HISTORY_SIZE)
	}
}

func (rm *ResultsMaster) addStatusCode(statusCode int) {
	if _, ok := rm.statusCodes[statusCode]; ok {
		rm.statusCodes[statusCode] += 1
	} else {
		rm.statusCodes[statusCode] = 1
	}
}

func (rm *ResultsMaster) getStatusCodeResults() ([]string, []float64) {
	keys := make([]string, 0, len(rm.statusCodes))
	values := make([]float64, 0, len(rm.statusCodes))
	for k, v := range rm.statusCodes {
		keys = append(keys, strconv.Itoa(k))
		values = append(values, float64(v))
	}
	return keys, values
}

func (rm *ResultsMaster) truncateArrays(maxLength int) {
	for len(rm.results) > maxLength {
		rm.results = rm.results[1:]
	}
	for len(rm.responseTimes) > maxLength {
		rm.responseTimes = rm.responseTimes[1:]
	}
}

func createResultsMaster() ResultsMaster {
	return ResultsMaster{
		make(chan HealthCheck, 10),
		make([]HealthCheck, 0),
		make(map[int]int, 0),
		make([]float64, 0),
	}
}
