package helpers

import (
	"errors"
	"sort"
	"time"

	"backend/datamodel"
)

// SortBy Sort data by time
type SortBy []datamodel.DockerImage

var (
	layout = "2006-01-02T15:04:05.000Z"
)

func (p SortBy) Len() int {
	return len(p)
}

func (p SortBy) Less(i, j int) bool {
	ti, erri := time.Parse(layout, p[i].Created)
	tj, errj := time.Parse(layout, p[j].Created)
	if erri != nil || errj != nil {
		return false
	}
	return ti.Before(tj)
}

func (p SortBy) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetLatest(docks SortBy) ([]datamodel.DockerImage, error) {
	if docks == nil || len(docks) == 0 {
		return nil, errors.New("No docker images")
	}
	sort.Sort(docks)
	if len(docks) > 5 {
		return docks[0:4], nil
	}
	result := make([]datamodel.DockerImage, 0)
	result = append(result, docks[0])
	return result, nil
}
