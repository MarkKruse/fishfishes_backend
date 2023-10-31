package repository

import (
	common "fishfishes_backend/common"
	log "github.com/sirupsen/logrus"
)

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
}

func NewRepo() Repo {
	return Repo{}
}

func (r Repo) GetAllSpots(id string) common.Fish_spots {
	log.Infof("User: %s", id)
	return spots
}
