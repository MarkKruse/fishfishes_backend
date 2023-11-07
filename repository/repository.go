package repository

import (
	"context"
	common "fishfishes_backend/common"
	"fishfishes_backend/common/mongo"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongoClient "go.mongodb.org/mongo-driver/mongo"
)

type SpotEntity struct {
	Id     string           `bson:"_id"`
	UserId string           `bson:"userId"`
	Spot   common.Fish_spot `bson:"spot"`
}

const Spot string = "spot"

// albums slice to seed record album data.
var spots = common.Fish_spots{
	Fish_spots: []common.Fish_spot{
		{
			Id: "12345",
			Marker: common.Marker{
				Title: "title1",
				Coordinates: common.Coordinates{
					Latitude:  54.110943,
					Longitude: 6.713647,
				},
			},
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
			Id: "12345",
			Marker: common.Marker{
				Title: "title2",
				Coordinates: common.Coordinates{
					Latitude:  55.078367,
					Longitude: 7.809639,
				},
			},
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

func (r Repo) InstallIndexes() error {

	//err := r.db.InstallIndex(Spot, "spot_idx", bson.D{
	//	{Key: "_id", Value: 1},
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}

func (r Repo) GetAllSpots(ctx context.Context, id string) (*[]common.Fish_spot, error) {

	filter := bson.D{
		{
			Key:   "userId",
			Value: fmt.Sprintf("%s", id),
		},
	}

	cur, err := r.db.Database.Collection(Spot).Find(ctx, filter)
	if err != nil {
		if err == mongoClient.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	defer mongo.CloseCursor(cur, ctx)

	var entities []common.Fish_spot

	for cur.Next(ctx) {
		var entity SpotEntity
		err := cur.Decode(&entity)
		if err != nil {
			if err == mongoClient.ErrNoDocuments {
				return nil, nil
			}
			return nil, err
		}
		entities = append(entities, entity.Spot)
	}

	return &entities, nil
}

func (r Repo) SaveSpot(ctx context.Context, userId string, spot common.Fish_spot) error {

	spotEntity := SpotEntity{
		Id:     uuid.New().String(),
		UserId: userId,
		Spot:   spot,
	}

	_, err := r.db.Database.Collection(Spot).InsertOne(ctx, spotEntity)
	if err != nil {
		return err
	}

	return nil
}
