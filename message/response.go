package message

import "encoding/json"

// Response is the format of the message from an Actions result
type Response struct {
	Meta   *json.RawMessage `json:"meta"`
	Status string           `json:"status"` // Status identifies the result status from the Action
	Error  string           `json:"error"`  // Error identifies any error that occured during the Action
	Output interface{}      `json:"output"` // Output contains the output of the Action
	Log    string           `json:"log"`    // Log contains any captured log information
}

// ResponseWrapper wraps the Response object because thats how we do it in komand/komand
type ResponseWrapper struct {
	Type    string   `json:"type"`
	Version string   `json:"version"`
	Body    Response `json:"body"`
}
