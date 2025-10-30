package logs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/logger"
)

type HttpLogger struct {
	domain string
}

func NewHttpLogger(domain string) *HttpLogger {
	return &HttpLogger{
		domain: domain,
	}
}

func (l *HttpLogger) Flush(ctx context.Context, log *logger.Log) error {
	url := fmt.Sprintf("%s/logs", l.domain)

	data, _ := json.Marshal(log)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error creating log: ", err)
		return err
	}

	return nil
}
