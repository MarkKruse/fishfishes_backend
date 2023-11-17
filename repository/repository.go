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
var fishListSaltwWater = []string{
	"Sword Fish", "Atlantic Cod", "Mackerel",
	"Trout", "Atlantic Salmon", "Tuna", "Shark",
	"Red Mullet", "Barramundi", "Mahi-Mahi", "Anchovy", "Haddock", "Red Seabream Fish", "Gold Line Fish",
	"Pollack", "Ocean Sunfish", "Northern Red Snapper", "Bonito", "Bluefish",
}

var fishListFreshWater = []string{
	"Common Carp", "Oscar Fish", "Wels Catfish", "Sauger Fish", "The Northern Pike", "Tench", "European Eel", "Cisco Fish", "Black Crappie",
	"Brown Bullhead Catfish", "Golden Shiner", "Largemouth Bass", "Fathead Minnow", "Walleye Fish", "Common Dace",
	"European Chub", "", "", "", "", ""}

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

func (r Repo) GetFishListSalt() []string {
	return fishListSaltwWater
}

func (r Repo) GetFishListFresh() []string {
	return fishListFreshWater
}
