package thenovadiary

import "testing"

func TestHappyCacheEmpty(t *testing.T) {
	cache := NewCache("Test")
	t.Logf("Cache directory: %s", cache.path.Name())
	delta, err := cache.Recover()
	if err == nil {
		// We have an existing cache
		t.Logf("Found existing cache: cleaning...")
		err := cache.Clean()
		if err != nil {
			t.Errorf("unable to clean cache: %v", err)
			t.FailNow()
		}
	} else {
		t.Logf("No existing test cache found")
	}
	// We get here, we know we have an empty cache
	records := []*Record{
		{
			Key:   "drink",
			Value: "kombucha",
		},
		{
			Key: "mountain",
			Value: struct {
				HasGlacier bool
			}{
				HasGlacier: true,
			},
		},
		{
			Key:   "hasClimbed",
			Value: 1,
		},
	}
	for _, record := range records {
		cache.Set(record.Key, record)
	}
	err = cache.Persist()
	if err != nil {
		t.Errorf("unable to persist cache: %v", err)
		t.FailNow()
	}
	// Manually flush records
	cache.Records = map[string]*Record{}

	delta, err = cache.Recover()
	if err != nil {
		t.Errorf("unable to recover cache: %v", err)
		t.FailNow()
	}
	t.Logf("Delta in cache recovery: %d", delta)
	if delta != 3 {
		t.Errorf("Delta (3) expected, have Delta (%d)", delta)
	}

	// Manually flush records, add 1 record
	cache.Records = map[string]*Record{
		"Name": {
			Key:   "Name",
			Value: "Nóva",
		},
	}

	delta, err = cache.Recover()
	if err != nil {
		t.Errorf("unable to recover cache: %v", err)
		t.FailNow()
	}
	t.Logf("Delta in cache recovery: %d", delta)
	if delta != 2 {
		t.Errorf("Delta (2) expected, have Delta (%d)", delta)
	}

	err = cache.Clean()
	if err != nil {
		t.Errorf("unable to clean cache: %v", err)
	}
}
