package alert

import (
	"time"
)

type SimplifiedAlertInfo struct {
	Onset             time.Time `json:"onset"`
	Ends              time.Time `json:"ends"`
	Severity          string    `json:"severity"`
	Certainty         string    `json:"certainty"`
	Urgency           string    `json:"urgency"`
	EventType         string    `json:"event_type"`
	Headline          string    `json:"headline"`
	Description       string    `json:"description"`
	RecommendedAction string    `json:"recommended_action"`
}

func ProcessNWSActiveAlertResponse(response *NWSActiveAlertsResponse) (activeAlerts []SimplifiedAlertInfo) {
	activeAlerts = []SimplifiedAlertInfo{}

	if len(response.Features) == 0 {
		return
	}

	for _, feature := range response.Features {
		if ShouldIncludeActiveAlert(feature.Properties) {
			activeAlerts = append(activeAlerts, TransformNWSFeature(feature))
		}
	}
	return
}

func TransformNWSFeature(nwsFeature Feature) SimplifiedAlertInfo {
	return SimplifiedAlertInfo{
		Onset:             nwsFeature.Properties.Onset,
		Ends:              nwsFeature.Properties.Ends,
		Severity:          nwsFeature.Properties.Severity,
		Certainty:         nwsFeature.Properties.Certainty,
		Urgency:           nwsFeature.Properties.Urgency,
		EventType:         nwsFeature.Properties.Event,
		Headline:          nwsFeature.Properties.Headline,
		Description:       nwsFeature.Properties.Description,
		RecommendedAction: nwsFeature.Properties.Response,
	}
}

func ShouldIncludeActiveAlert(nwsAlertProps Properties) (ok bool) {
	const ValidAlertStatus = "Actual"
	ok = nwsAlertProps.Status == ValidAlertStatus && validateMessageType(nwsAlertProps.MessageType)
	return
}

func validateMessageType(messageType string) (ok bool) {
	ok = !(messageType == "Cancel" || messageType == "Error")
	return
}
