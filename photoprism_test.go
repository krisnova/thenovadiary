package thenovadiary

import (
	"testing"
	"time"

	"github.com/kris-nova/logger"

	"github.com/kris-nova/photoprism-client-go/api/v1"
)

func TestFavoriteFindPhoto(t *testing.T) {
	expected := api.Photo{
		PhotoTitle:    "Winner",
		PhotoFavorite: true,
	}
	loser1 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	photos := []api.Photo{
		{
			PhotoTitle:    "Loser",
			PhotoFavorite: false,
		},
		{
			PhotoTitle:    "Loser2",
			PhotoFavorite: false,
		},
		{
			PhotoTitle:    "Loser4",
			PhotoFavorite: false,
		},
		{
			PhotoTitle:    "Loser3",
			PhotoFavorite: false,
		},
	}

	SetCustomData(&CustomData{}, &loser1)
	photos = append(photos, expected)
	photos = append(photos, loser1)
	actual, err := FindPhoto(photos)
	if err != nil {
		t.Errorf("unable to find favorite photo: %v", err)
		t.FailNow()
	}
	if actual.PhotoTitle != expected.PhotoTitle {
		t.Errorf("unable to match actual (%s) and expected (%s)", actual.PhotoTitle, expected.PhotoTitle)
		t.FailNow()
	}

}

func TestSadNewFindPhoto(t *testing.T) {
	var photos []api.Photo
	expected := api.Photo{
		PhotoTitle:    "Winner",
		PhotoFavorite: false,
	}
	loser1 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: true,
	}
	loser2 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	now := time.Now()
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser1)
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser2)
	photos = append(photos, loser2)
	photos = append(photos, expected)
	photos = append(photos, loser1)
	actual, err := FindPhoto(photos)
	if err != nil {
		t.Errorf("unable to find favorite photo: %v", err)
		t.FailNow()
	}
	if actual.PhotoTitle == expected.PhotoTitle {
		t.Errorf("unable to match actual (%s) and expected (%s)", actual.PhotoTitle, expected.PhotoTitle)
		t.FailNow()
	}
}

func TestHappyNewFindPhoto(t *testing.T) {
	var photos []api.Photo
	expected := api.Photo{
		PhotoTitle:    "Winner",
		PhotoFavorite: false,
	}
	loser1 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	loser2 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	now := time.Now()
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser1)
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser2)
	photos = append(photos, loser2)
	photos = append(photos, expected)
	photos = append(photos, loser1)
	actual, err := FindPhoto(photos)
	if err != nil {
		t.Errorf("unable to find favorite photo: %v", err)
		t.FailNow()
	}
	if actual.PhotoTitle != expected.PhotoTitle {
		t.Errorf("unable to match actual (%s) and expected (%s)", actual.PhotoTitle, expected.PhotoTitle)
		t.FailNow()
	}
}

func TestFavoriteOldest(t *testing.T) {
	logger.Level = 4
	var photos []api.Photo
	expected := api.Photo{
		PhotoTitle:    "Winner",
		PhotoFavorite: false,
	}
	loser1 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	loser2 := api.Photo{
		PhotoTitle:    "Loser",
		PhotoFavorite: false,
	}
	now := time.Now()
	yesterday := TimeDeltaDaysFromNow(-100)
	SetCustomData(&CustomData{
		LastTweet: &yesterday,
	}, &expected)
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser1)
	SetCustomData(&CustomData{
		LastTweet: &now,
	}, &loser2)
	photos = append(photos, loser1)
	photos = append(photos, expected)
	photos = append(photos, loser2)
	actual, err := FindPhoto(photos)
	if err != nil {
		t.Errorf("unable to find favorite photo: %v", err)
		t.FailNow()
	}
	if actual.PhotoTitle != expected.PhotoTitle {
		t.Errorf("unable to match actual (%s) and expected (%s)", actual.PhotoTitle, expected.PhotoTitle)
		t.FailNow()
	}
}
