package entity

// swagger:model LocationInfo
type LocationInfo struct {

	//latitude of the location coordinates
	Lat string `json:"lat" bson:"lat"`
	//longitude of the location coordinates
	Lon string `json:"lon" bson:"lon"`
	//AddressInfo
	Address *AddressInfo `json:"address" bson:"address"`
}
