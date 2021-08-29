package services

import (
	"fmt"
	"github.com/harborchurchla/harborbot/internal/entities"
	"gopkg.in/Iwark/spreadsheet.v2"
	"strings"
	"time"
)

type ScheduleService struct {
	client  *spreadsheet.Service
	sheetId string
}

func NewScheduleService(client *spreadsheet.Service, sheetId string) *ScheduleService {
	return &ScheduleService{client, sheetId}
}

func (s *ScheduleService) GetByTeam(team string) (*entities.Schedule, error) {
	team = strings.Title(strings.ToLower(team))
	ss, err := s.client.FetchSpreadsheet(s.sheetId)
	if err != nil {
		return nil, fmt.Errorf("error fetching schedule spreadsheet: %v", err)
	}
	sheet, err := ss.SheetByTitle(strings.Title(strings.ToLower(team)))
	if err != nil {
		return nil, fmt.Errorf("error fetching sheet for team %s - %v", team, err)
	}

	var scheduleEntries []*entities.ScheduleEntry
	for i, row := range sheet.Rows {
		// First row is the columns
		if i == 0 {
			continue
		}

		t, err := time.Parse("1/2/2006", row[0].Value)
		if err != nil {
			return nil, fmt.Errorf("error parsing time for data %s - %v", row[0].Value, err)
		}

		scheduleEntries = append(scheduleEntries, &entities.ScheduleEntry{Date: entities.ScheduleDate(t), TeamMembers: row[1].Value, Notes: row[2].Value})
	}

	return &entities.Schedule{Team: team, Entries: scheduleEntries}, nil
}
