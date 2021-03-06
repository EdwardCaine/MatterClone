// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InferenceClassification undocumented
type InferenceClassification struct {
	// Entity is the base model of InferenceClassification
	Entity
	// Overrides undocumented
	Overrides []InferenceClassificationOverride `json:"overrides,omitempty"`
}

// InferenceClassificationOverride undocumented
type InferenceClassificationOverride struct {
	// Entity is the base model of InferenceClassificationOverride
	Entity
	// ClassifyAs undocumented
	ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`
	// SenderEmailAddress undocumented
	SenderEmailAddress *EmailAddress `json:"senderEmailAddress,omitempty"`
}

// InferenceData undocumented
type InferenceData struct {
	// Object is the base model of InferenceData
	Object
	// ConfidenceScore undocumented
	ConfidenceScore *float64 `json:"confidenceScore,omitempty"`
	// UserHasVerifiedAccuracy undocumented
	UserHasVerifiedAccuracy *bool `json:"userHasVerifiedAccuracy,omitempty"`
}
