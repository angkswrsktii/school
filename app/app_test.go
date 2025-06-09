package app_test

import (
	"Goland/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchoolActivityApp(t *testing.T) {
	t.Run("successfully add activity", func(t *testing.T) {
		app.InitializeActivities()
		err := app.InsertActivity("Math Club")
		assert.NoError(t, err)
	})

	t.Run("cannot add empty activity", func(t *testing.T) {
		err := app.InsertActivity("")
		assert.Error(t, err)
		assert.ErrorIs(t, err, app.ErrEmptyActivity)
	})

	t.Run("successfully list all activities", func(t *testing.T) {
		app.InitializeActivities()
		_ = app.InsertActivity("Math Club")
		_ = app.InsertActivity("Science Fair")

		activities := app.GetAllActivities()
		assert.Equal(t, 2, len(activities))
		assert.Equal(t, "Math Club", activities[0].Name())
		assert.Equal(t, "Science Fair", activities[1].Name())
	})

	t.Run("mark activity as completed", func(t *testing.T) {
		app.InitializeActivities()
		err := app.InsertActivity("Math Club")
		assert.NoError(t, err)

		activityItem, err := app.GetActivity(0)
		assert.NoError(t, err)

		assert.Equal(t, false, activityItem.IsCompleted())
		activityItem.ToggleCompletion()
		assert.Equal(t, true, activityItem.IsCompleted())

		activityItem, err = app.GetActivity(0)
		assert.NoError(t, err)
		assert.Equal(t, true, activityItem.IsCompleted())
	})

	t.Run("unmark activity as completed", func(t *testing.T) {
		app.InitializeActivities()
		_ = app.InsertActivity("Math Club")

		activities := app.GetAllActivities()
		assert.False(t, activities[0].IsCompleted())

		activities[0].ToggleCompletion()
		assert.True(t, activities[0].IsCompleted())

		activities[0].ToggleCompletion()
		assert.False(t, activities[0].IsCompleted())
	})

	t.Run("successfully delete canceled activity", func(t *testing.T) {
		app.InitializeActivities()
		err := app.RemoveActivity(0)
		assert.Error(t, err)
		assert.ErrorIs(t, err, app.ErrNotFound)

		_ = app.InsertActivity("Math Club")
		err = app.RemoveActivity(0)
		assert.NoError(t, err)

		err = app.RemoveActivity(0)
		assert.ErrorIs(t, err, app.ErrNotFound)
	})
}
