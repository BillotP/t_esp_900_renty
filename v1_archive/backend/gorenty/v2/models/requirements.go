package models

// Requirement is a rent offer candidature requirement
type Requirement struct {
	Base
	Name      string   `json:"name"`
	Required  bool     `json:"required"`
	Condition *string  `json:"condition,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}

// Requirements is a rent offer candidature requirement
type Requirements struct {
	Base
	Candidate        []Requirement `json:"candidate"`
	Warrants         []Requirement `json:"warrants"`
	OtherConditions  string        `json:"other_conditions"`
}
