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

func Publish(ctx context.Context, event Event) error {
	url := fmt.Sprintf("%s/event", eventHandlerDomain)

	domainEvent := event.ToDomain()
	data, _ := json.Marshal(domainEvent)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error publishing event: ", err)
		return err
	}

	return nil
}
