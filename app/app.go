package app

import "errors"

var (
	ErrEmptyActivity = errors.New("activity is empty")
	ErrNotFound      = errors.New("activity not found")
)

type SchoolActivity struct {
	name        string
	isCompleted bool
}

func (s *SchoolActivity) Name() string {
	return s.name
}

func (s *SchoolActivity) ToggleCompletion() {
	s.isCompleted = !s.isCompleted
}

func (s *SchoolActivity) IsCompleted() bool {
	return s.isCompleted
}

func NewSchoolActivity(name string) (SchoolActivity, error) {
	if name == "" {
		return SchoolActivity{}, ErrEmptyActivity
	}
	return SchoolActivity{name: name, isCompleted: false}, nil
}

var activities []SchoolActivity

func InitializeActivities() {
	activities = []SchoolActivity{}
}

func InsertActivity(name string) error {
	if activities == nil {
		InitializeActivities()
	}

	act, err := NewSchoolActivity(name)
	if err != nil {
		return err
	}

	activities = append(activities, act)
	return nil
}

func GetAllActivities() []SchoolActivity {
	if activities == nil {
		InitializeActivities()
	}
	return activities
}

func GetActivity(id int) (SchoolActivity, error) {
	if activities == nil {
		InitializeActivities()
	}
	if id < 0 || id >= len(activities) {
		return SchoolActivity{}, ErrNotFound
	}
	return activities[id], nil
}

func RemoveActivity(id int) error {
	if activities == nil {
		InitializeActivities()
	}
	if id < 0 || id >= len(activities) {
		return ErrNotFound
	}
	activities = append(activities[:id], activities[id+1:]...)
	return nil
}
