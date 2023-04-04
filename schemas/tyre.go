package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tyre struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	Code                      string             `bson:"code,omitempty"`
	Name                      string             `bson:"name,omitempty"`
	Reference                 string             `bson:"reference,omitempty"`
	Brand                     string             `bson:"brand,omitempty"`
	Width                     string             `bson:"width,omitempty"`
	AspectRatio               string             `bson:"aspect_ratio,omitempty"`
	Construction              string             `bson:"construction,omitempty"`
	Rim                       string             `bson:"rim,omitempty"`
	Load                      string             `bson:"load,omitempty"`
	Speed                     string             `bson:"speed,omitempty"`
	Season                    string             `bson:"season,omitempty"`
	EprelId                   string             `bson:"eprel_id,omitempty"`
	ExternalRollingNoiseClass string             `bson:"external_rolling_noise_class,omitempty"`
	ExternalRollingNoiseLevel string             `bson:"external_rolling_noise_level,omitempty"`
	LoadVersion               string             `bson:"load_version,omitempty"`
	FuelEfficiency            string             `bson:"fuel_efficiency,omitempty"`
	WetGripClass              string             `bson:"wet_grip_class,omitempty"`
	ImageUrl                  string             `bson:"image_url,omitempty"`
}
