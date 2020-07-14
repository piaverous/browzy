package health

import (
	"net/http"
	"sync"
	"time"
)

var HealthCheckFrequency int = 1000 // Milliseconds interval

type HealthCheck struct {
	time         time.Time
	url          string
	err          error
	status       int
	responseTime time.Duration
}

type Checker struct {
	url        string
	httpClient *http.Client
	running    bool
	results    chan<- HealthCheck
}

func (c *Checker) check() {
	startTime := time.Now()
	res, err := c.httpClient.Get(c.url)
	responseTime := time.Now().Sub(startTime)

	statusCode := 0
	if err == nil {
		statusCode = res.StatusCode
	}
	if c.running {
		c.results <- HealthCheck{
			startTime,
			c.url,
			err,
			statusCode,
			responseTime,
		}
	}
}

func (c *Checker) start(wg *sync.WaitGroup) {
	defer wg.Done()

	tick := time.NewTicker(time.Duration(HealthCheckFrequency) * time.Millisecond)
	for c.running {
		c.check()
		<-tick.C
	}
}

func (c *Checker) stop() {
	c.running = false
	close(c.results)
}

func createChecker(url string, channel chan<- HealthCheck) Checker {
	tr := &http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Duration(HealthCheckFrequency) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	return Checker{
		url,
		client,
		true,
		channel,
	}
}
