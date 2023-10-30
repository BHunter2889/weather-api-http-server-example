package alert_test

import (
	"fmt"
	"github.com/BHunter2889/weather-api-http-server-example/alert"

	"testing"
	"time"
)

func TestTransformNWSFeature(t *testing.T) {
	nwsFeature := alert.Feature{
		ID:   "test_id",
		Type: "test_type",
		Properties: alert.Properties{
			Onset:       time.Now(),
			Ends:        time.Now().Add(time.Hour),
			Severity:    "Moderate",
			Certainty:   "Likely",
			Urgency:     "Immediate",
			Event:       "TestEvent",
			Headline:    "TestHeadline",
			Description: "TestDescription",
			Response:    "TestAction",
		},
	}

	result := alert.TransformNWSFeature(nwsFeature)

	expectedResult := alert.SimplifiedAlertInfo{
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

	equals(t, expectedResult, result)
}

func TestShouldIncludeActiveAlert(t *testing.T) {
	validAlert := alert.Properties{
		Status:      "Actual",
		MessageType: "Alert",
	}

	invalidAlert := alert.Properties{
		Status:      "TestStatus",
		MessageType: "Cancel",
	}

	const ShouldBeTrueMsg = "ShouldIncludeActiveAlert should return true for valid alert, but returned false."
	const ShouldBeFalseMsg = "ShouldIncludeActiveAlert should return false for invalid alert, but returned true."

	// valid
	assert(t, alert.ShouldIncludeActiveAlert(validAlert), ShouldBeTrueMsg)
	// invalid
	assert(t, !alert.ShouldIncludeActiveAlert(invalidAlert), ShouldBeFalseMsg)
}

func TestProcessNWSActiveAlertResponse(t *testing.T) {
	validAlertFeature := alert.Feature{
		ID:   "test_id_1",
		Type: "test_type_1",
		Properties: alert.Properties{
			Onset:       time.Now(),
			Ends:        time.Now().Add(time.Hour),
			Severity:    "Moderate",
			Certainty:   "Likely",
			Urgency:     "Immediate",
			Event:       "TestEvent1",
			Headline:    "TestHeadline1",
			Status:      "Actual",
			MessageType: "Alert",
		},
	}

	invalidAlertFeature := alert.Feature{
		ID:   "test_id_2",
		Type: "test_type_2",
		Properties: alert.Properties{
			Onset:       time.Now(),
			Ends:        time.Now().Add(time.Hour),
			Severity:    "Moderate",
			Certainty:   "Likely",
			Urgency:     "Immediate",
			Event:       "TestEvent2",
			Headline:    "TestHeadline2",
			Status:      "TestStatus",
			MessageType: "Cancel",
		},
	}

	nwsResponse := alert.NWSActiveAlertsResponse{
		Features: []alert.Feature{validAlertFeature, invalidAlertFeature},
	}
	result := alert.ProcessNWSActiveAlertResponse(&nwsResponse)

	// Test length is correct
	assert(
		t,
		len(result) == 1,
		fmt.Sprintf(
			"ProcessNWSActiveAlertResponse should include only and all valid alerts.\nExpected: %d\n Got: %d\n",
			1,
			len(result),
		),
	)

	expectedResult := alert.SimplifiedAlertInfo{
		Onset:             validAlertFeature.Properties.Onset,
		Ends:              validAlertFeature.Properties.Ends,
		Severity:          validAlertFeature.Properties.Severity,
		Certainty:         validAlertFeature.Properties.Certainty,
		Urgency:           validAlertFeature.Properties.Urgency,
		EventType:         validAlertFeature.Properties.Event,
		Headline:          validAlertFeature.Properties.Headline,
		Description:       validAlertFeature.Properties.Description,
		RecommendedAction: validAlertFeature.Properties.Response,
	}

	equals(t, expectedResult, result[0])
}
