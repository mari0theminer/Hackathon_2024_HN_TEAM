package service

import (
	"github.com/castai/promwrite"
	"time"
)

func Write(metricName string, metricValue float64, username string, password string, URL string) {

	labels := []promwrite.Label{
		{Name: "__name__", Value: metricName},
	}
	metric := promwrite.TimeSeries{
		Labels: labels,
		Sample: promwrite.Sample{
			Time:  time.Now(),
			Value: metricValue,
		},
	}
	metrics := []promwrite.TimeSeries{metric}
	client, headers := SetupProm(username, password, URL)

	SendToProm(metrics, client, headers)
}
