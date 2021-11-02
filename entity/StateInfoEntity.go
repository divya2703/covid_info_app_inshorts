package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:parameters getStateReport
type StateName struct {

	// The name of the state
	// in: path
	// required: true
	State string `json:"state"`
}

// swagger:model StateReport
type StateReport struct {

	// swagger:ignore
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	//name of the state
	State string `json:"State" bson:"State"`

	//last updated time of the report
	Last_Updated_Time string `json:"Last_Updated_Time" bson:"Last_Updated_Time"`

	//number of active covid cases in that state
	Active int `json:"Active" bson:"Active"`
}
