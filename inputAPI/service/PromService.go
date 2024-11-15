package service

import (
	"context"
	"fmt"
	"github.com/castai/promwrite"
	"os"
)

func SendToProm(metrics []promwrite.TimeSeries, client *promwrite.Client, headers map[string]string) {

	_, err := client.Write(context.Background(), &promwrite.WriteRequest{
		TimeSeries: metrics,
	}, promwrite.WriteHeaders(headers))
	if err != nil {
		fmt.Println(fmt.Sprintf("could not send metrics to prometheus: %v", err.Error()))
		os.Exit(1)
	}
	fmt.Println("Metrics were send successfully")

}
