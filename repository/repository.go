package repository

import (
	"context"
	common "fishfishes_backend/common"
	"fishfishes_backend/common/mongo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type SpotEntity struct {
	UserId string           `bson:"_id"`
	Spot   common.Fish_spot `bson:"spot"`
}

const Spot string = "spot"

// albums slice to seed record album data.
var spots = common.Fish_spots{
	Fish_spots: []common.Fish_spot{
		{
			ID:  "12345",
			LAT: 54.110943,
			LNG: 6.713647,
			Catches: []common.Catch{
				{
					Fish:      "",
					Number:    1,
					Size:      80.45,
					Deep:      500,
					Time:      "Morning",
					Equipment: common.Equipment{Name: "Angel 1"},
				},
			},
		},
		{
			ID:  "12345",
			LAT: 55.078367,
			LNG: 7.809639,
			Catches: []common.Catch{
				{
					Fish:      "",
					Number:    1,
					Size:      80.45,
					Deep:      500,
					Time:      "Morning",
					Equipment: common.Equipment{Name: "Angel 3"},
				},
			},
		},
	},
}

type Repo struct {
	db *mongo.Database
}

func NewRepo(db *mongo.Database) Repo {
	return Repo{
		db: db,
	}
}

func (r Repo) GetAllSpots(ctx context.Context, id string) common.Fish_spots {

	filter := bson.D{
		{
			Key:   "_id",
			Value: fmt.Sprintf("%s_%s_%s", id),
		},
	}

	r.db.Database.Collection(Spot).Find(ctx, filter)

	return spots
}

func (r Repo) SaveSpot(ctx context.Context, userId string, spot common.Fish_spot) error {

	spotEntity := SpotEntity{
		UserId: userId,
		Spot:   spot,
	}

	_, err := r.db.Database.Collection(Spot).InsertOne(ctx, spotEntity)
	if err != nil {
		return err
	}

	return nil
}
