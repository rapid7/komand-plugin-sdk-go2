package message

import "encoding/json"

// TriggerEvent This an event type for triggering a single workflow.
// Used for API triggers and Republishing events.
type TriggerEvent struct {
	Type    string           `json:"type"`     // Is it a single trigger or a multi-cast trigger?
	ID      string           `json:"id"`       // This will become the JobID for the resulting job
	GroupID string           `json:"group_id"` // This is the id of the originating job, in the event a job has been re-run multiple times.
	Meta    TriggerEventMeta `json:"meta"`
	Output  interface{}      `json:"output"`
}

// TriggerEventMeta is the base information needed to run a step. It contains information both for
// single and multi-cast triggers. The type of trigger will determine if it uses the TriggerConfigID
// or the Step/Workflow information
type TriggerEventMeta struct {
	// TriggerConfigID is used for multi-cast triggers
	TriggerConfigID string `json:"trigger_config_id"`
	// StepUID, WorkflowUID, and WorkflowVersionUID are used for single-cast triggers
	StepUID            string `json:"step_uid"`
	WorkflowUID        string `json:"workflow_uid,omitempty"`
	WorkflowVersionUID string `json:"workflow_version_uid,omitempty"`
}

// RawTriggerEvent This an event type for triggering a single workflow.
// Used for API triggers and Republishing events.
// RawTriggerEvent is the consumer version where the consumer will want to make a choice about the data
// structure it marshals the output into
type RawTriggerEvent struct {
	Type    string           `json:"type"`     // Is it a single trigger or a multi-cast trigger?
	ID      string           `json:"id"`       // This will become the JobID for the resulting job
	GroupID string           `json:"group_id"` // This is the id of the originating job, in the event a job has been re-run multiple times.
	Meta    TriggerEventMeta `json:"meta"`
	Output  json.RawMessage  `json:"output"`
}
