package entities

import (
	"fmt"
	"time"
)

type ScheduleDate time.Time

type Schedule struct {
	Team    string           `json:"team"`
	Entries []*ScheduleEntry `json:"entries"`
}

type ScheduleEntry struct {
	Date        ScheduleDate `json:"date"`
	TeamMembers string       `json:"team_members"`
	Notes       string       `json:"notes"`
}

func (s *Schedule) GetFutureEntries() []*ScheduleEntry {
	var future []*ScheduleEntry
	for _, entry := range s.Entries {
		if time.Time(entry.Date).YearDay() >= time.Now().YearDay() {
			future = append(future, entry)
		}
	}
	return future
}

func (t ScheduleDate) MarshalJSON() ([]byte, error) {
	formattedDate := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(formattedDate), nil
}
