package events

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var eventHandlerDomain string

func SetEventHandlerDomain(domain string) {
	eventHandlerDomain = domain
}

func Publish(ctx context.Context, event *Event) {
	go func() {

		url := fmt.Sprintf("%s/event", eventHandlerDomain)

		data, _ := json.Marshal(event)

		req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating request: ", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error publishing event: ", err)
		}
	}()
}
