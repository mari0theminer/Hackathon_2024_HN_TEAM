package service

import (
	"encoding/base64"
	"fmt"
	"github.com/castai/promwrite"
)

func SetupProm(username string, password string, URL string) (*promwrite.Client, map[string]string) {
	client := promwrite.NewClient(URL)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", encodedCredentials),
	}
	return client, headers
}
