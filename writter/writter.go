package writter

import (
	"strings"

	"firebase.google.com/go/messaging"
	"github.com/pamungkaski/camar/datamodel"
)

type Writer struct{}

func (a *Writer) CreateAlertMessage(disaster datamodel.GeoJSON, alerts []string) (messaging.Message, error) {
	// Cut AlertBody unnecessary part
	alen := len(alerts)
	alerts = append(alerts[:0], alerts[1:]...)
	alerts = append(alerts[:alen-4], alerts[alen-1:]...)
	alertBody := strings.Join(alerts, " ")

	message := messaging.Message{
		Data: map[string]string{
			"URL": disaster.URL,
		},
		Notification: &messaging.Notification{
			Title: "Warning",
			Body:  alertBody,
		},
	}
	return message, nil
}
