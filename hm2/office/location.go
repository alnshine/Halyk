package office

import (
	"errors"
	"fmt"
)

type Location interface {
	GetLocationTitle() string
	CheckMoveToArea(Location) bool
}

func (b *BaseLocationParams) GetLocationTitle() string {
	return b.title
}
func (b *BaseLocationParams) CheckMoveToArea(l Location) bool {
	if l == nil {
		return false
	}
	res := Con(b.nearWith, l.GetLocationTitle())
	return res
}

type BaseLocationParams struct {
	title    string
	nearWith []string
}

var ErrUnknownLocationType = errors.New("unknown location type")

func NewLocationFactory(title string) (Location, error) {
	switch title {
	case "office":
		return newOffice(), nil
	case "workArea":
		return newWorkArea(), nil
	case "servers":
		return newServers(), nil
	}
	return nil, fmt.Errorf("newLocation: %w", ErrUnknownLocationType)
}

// TODO:: impl Location
type office struct {
	BaseLocationParams
}

func newOffice() Location {
	nearWith := []string{"workArea"}

	return &office{
		BaseLocationParams{
			title:    "office",
			nearWith: nearWith,
		},
	}
}

// TODO:: impl Location
type workArea struct {
	BaseLocationParams
}

func newWorkArea() Location {
	nearWith := []string{"office", "servers"}

	return &workArea{
		BaseLocationParams{
			title:    "workArea",
			nearWith: nearWith,
		},
	}
}

// TODO:: impl Location
type servers struct {
	BaseLocationParams
}

func newServers() Location {
	nearWith := []string{"workArea"}

	return &servers{
		BaseLocationParams{
			title:    "servers",
			nearWith: nearWith,
		},
	}
}
