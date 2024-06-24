package health_test

import (
	"testing"
	"time"

	"api.default.marincor.pt/app/repository/health"
	"api.default.marincor.pt/pkg/app"
	"github.com/stretchr/testify/assert"
)

var healthRepo *health.Repository

func TestMain(m *testing.M) {
	app.ApplicationInit()

	healthRepo = health.New()

	m.Run()
}

func TestGetOne(t *testing.T) {
	t.Parallel()

	t.Run("get correctly sync", func(t *testing.T) {
		t.Parallel()

		now := time.Now()

		err := healthRepo.Insert(now)
		if err != nil {
			t.Errorf("health not insert, error: %v", err)
		}

		health, err := healthRepo.GetOne(now)
		if err != nil {
			t.Errorf("health not retrieved, error: %v", err)
		}

		// remove microseconds
		now = now.Truncate(time.Second)
		*health.Sync = health.Sync.Truncate(time.Second)

		assert.Equal(t, &now, health.Sync, "health is not synced")
	})

	t.Run("get wrong sync", func(t *testing.T) {
		t.Parallel()

		now := time.Now()

		err := healthRepo.Insert(now)
		if err != nil {
			t.Errorf("health not insert, error: %v", err)
		}

		now = now.AddDate(0, 0, 1)

		health, err := healthRepo.GetOne(now)
		if err != nil {
			t.Errorf("health not retrieved, error: %v", err)
		}

		assert.NotEqual(t, nil, health, "health should not have been retreived")
	})
}
