package repository

import (
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepository_StartActivity(t *testing.T) {
	a := Activities{
		ActivityType:   100,
		StartTimestamp: time.Now(),
	}

	result, err := testRepo.StartActivity(a)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back:", result.ID)
	}
}

func TestSQLiteRepository_AllActivities(t *testing.T) {
	a, err := testRepo.AllActivities()
	if err != nil {
		t.Error("get all failed:", err)
	}

	if len(a) != 1 {
		t.Error("wrong number of rows returned; expected 1, but got ", len(a))
	}
}

func TestSQLiteRepository_GetActivityByID(t *testing.T) {
	a, err := testRepo.GetActivityByID(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	if a.ActivityType != 100 {
		t.Error("wrong type returned; expected 100 but got", a.ActivityType)

	}

	_, err = testRepo.GetActivityByID(2)
	if err == nil {
		t.Error("get one returned value for non-existend id")
	}
}

func TestSQLiteRepository_UpdateActivity(t *testing.T) {
	a, err := testRepo.GetActivityByID(1)
	if err != nil {
		t.Error(err)
	}

	a.ActivityType = 200

	err = testRepo.UpdateActivity(1, *a)
	if err != nil {
		t.Error("update failed:", err)
	}

	a.EndTimestamp = time.Now()

	err = testRepo.UpdateActivity(1, *a)
	if err != nil {
		t.Error("update failed:", err)
	}
}

func TestSQLiteRepository_FloodActivity(t *testing.T) {
	a := Activities{
		ActivityType:   200,
		StartTimestamp: time.Now(),
	}

	at, err := testRepo.AllActivityType()
	if err != nil {
		t.Error("Error getting Activity Type: ", err)
	}

	if len(at) < 1 {
		t.Error("No activityType found!")
	}

	for i := 1; i <= 2000; i++ {
		result, err := testRepo.StartActivity(a)
		if err != nil {
			t.Error("insert failed:", err)
		}

		if result.ID <= 0 {
			t.Error("invalid id sent back:", result.ID)
		}

		a.EndTimestamp = time.Now()
		atIndex := rand.Intn(len(at))
		a.ActivityType = at[atIndex].ID

		err = testRepo.UpdateActivity(result.ID, a)
		if err != nil {
			t.Error("update failed:", err)
		}
	}
}

func TestSQLiteRepository_DeleteActivity(t *testing.T) {
	err := testRepo.DeleteActivity(1)
	if err != nil {
		t.Error("failed to delete activity", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	err = testRepo.DeleteActivity(-4)
	if err == nil {
		t.Error("no error when trying to delete non-existend record")
	}
}
