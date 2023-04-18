package office

import (
	"errors"
	"fmt"
)

type Employee interface {
	GetCurrentLocation() Location
	MoveToLocation(Location) error
}

func (b *baseEmployeeParams) GetCurrentLocation() Location {
	return b.location
}

func (b *baseEmployeeParams) MoveToLocation(l Location) error {
	if l == nil {
		return fmt.Errorf("MoveToLocation:no location")
	}

	locTitle := l.GetLocationTitle()

	res := Con(b.accesses, locTitle)
	if b.location == nil {
		vrem, err := NewLocationFactory(locTitle)
		if err != nil {
			return fmt.Errorf("movetoloc%w", err)
		}
		b.location = vrem
	}
	if locTitle == b.location.GetLocationTitle() {
		fmt.Println("move to", locTitle)
		return nil
	}
	if res {
		if b.location.CheckMoveToArea(l) {
			b.location = l
			fmt.Println("move to", locTitle)
			return nil
		}
		fmt.Errorf("movetolocation check")
	}
	return fmt.Errorf("movetolock:con")
}

type baseEmployeeParams struct {
	location Location
	accesses []string
}

var ErrUnknownEmplType = errors.New("unknown employee type")

func NewEmployeeFactory(title string) (Employee, error) {
	switch title {
	case "hr":
		return newHr(), nil
	case "itSecurity":
		return newItSecurity(), nil
	}
	return nil, fmt.Errorf("newEmployee: %w", ErrUnknownEmplType)
}

// TODO:: must impl Employee
type hr struct {
	baseEmployeeParams
}

func newHr() Employee {
	accesses := []string{"office", "workArea"}

	return &hr{
		baseEmployeeParams{
			accesses: accesses,
		},
	}
}

// TODO:: must impl Employee
type itSecurity struct {
	baseEmployeeParams
}

func newItSecurity() Employee {
	accesses := []string{"office", "workArea", "servers"}

	return &itSecurity{
		baseEmployeeParams{
			accesses: accesses,
		},
	}
}
