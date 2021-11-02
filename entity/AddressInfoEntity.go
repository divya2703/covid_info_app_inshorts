package entity

//swagger: model
type AddressInfo struct {

	//name of the village
	Village string `json:"village" bson:"village"`
	//name of the county
	County string `json:"county" bson:"county"`
	//name of the state_district
	State_district string `json:"state_district" bson:"state_district"`
	//name of the district
	State string `json:"state" bson:"state"`
	//name of the country
	Country string `json:"country" bson:"country"`
	//postcode of the location
	Postcode string `json:"postcode" bson:"postcode"`
	//country code for the country
	Country_code string `json:"country_code" bson:"country_code"`
}
