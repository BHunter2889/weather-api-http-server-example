package alert

import "time"

type NWSActiveAlertsResponse struct {
	Type       string     `json:"type"`
	Features   []Feature  `json:"features"`
	Title      string     `json:"title"`
	Updated    time.Time  `json:"updated"`
	Pagination Pagination `json:"pagination"`
}

type Geocode struct {
	Ugc  []string `json:"UGC"`
	Same []string `json:"SAME"`
}

type References struct {
	ID         string    `json:"@id"`
	Identifier string    `json:"identifier"`
	Sender     string    `json:"sender"`
	Sent       time.Time `json:"sent"`
}

type Parameters struct {
	AdditionalProp1 []any `json:"additionalProp1"`
	AdditionalProp2 []any `json:"additionalProp2"`
	AdditionalProp3 []any `json:"additionalProp3"`
}

type Properties struct {
	ID            string       `json:"id"`
	AreaDesc      string       `json:"areaDesc"`
	Geocode       Geocode      `json:"geocode"`
	AffectedZones []string     `json:"affectedZones"`
	References    []References `json:"references"`
	Sent          time.Time    `json:"sent"`
	Effective     time.Time    `json:"effective"`
	Onset         time.Time    `json:"onset"`
	Expires       time.Time    `json:"expires"`
	Ends          time.Time    `json:"ends"`
	Status        string       `json:"status"`
	MessageType   string       `json:"messageType"`
	Category      string       `json:"category"`
	Severity      string       `json:"severity"`
	Certainty     string       `json:"certainty"`
	Urgency       string       `json:"urgency"`
	Event         string       `json:"event"`
	Sender        string       `json:"sender"`
	SenderName    string       `json:"senderName"`
	Headline      string       `json:"headline"`
	Description   string       `json:"description"`
	Instruction   string       `json:"instruction"`
	Response      string       `json:"response"`
	Parameters    Parameters   `json:"parameters"`
}

type Feature struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Pagination struct {
	Next string `json:"next"`
}
