package notifications

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendNotification(ctx context.Context) error {
	log.Println("########Inside Activity")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	return nil
}
