// Code generated by go-enum
// DO NOT EDIT!

package controler

import (
	"fmt"
	"strings"
)

const (
	// EventCreated is a Event of type Created
	EventCreated Event = iota
	// EventRemoved is a Event of type Removed
	EventRemoved
	// EventStarted is a Event of type Started
	EventStarted
	// EventRestarted is a Event of type Restarted
	EventRestarted
	// EventStopped is a Event of type Stopped
	EventStopped
	// EventUpdated is a Event of type Updated
	EventUpdated
	// EventEnabled is a Event of type Enabled
	EventEnabled
	// EventDisabled is a Event of type Disabled
	EventDisabled
	// EventJoined is a Event of type Joined
	EventJoined
	// EventLeaved is a Event of type Leaved
	EventLeaved
)

const _EventName = "CreatedRemovedStartedRestartedStoppedUpdatedEnabledDisabledJoinedLeaved"

var _EventMap = map[Event]string{
	0: _EventName[0:7],
	1: _EventName[7:14],
	2: _EventName[14:21],
	3: _EventName[21:30],
	4: _EventName[30:37],
	5: _EventName[37:44],
	6: _EventName[44:51],
	7: _EventName[51:59],
	8: _EventName[59:65],
	9: _EventName[65:71],
}

// String implements the Stringer interface.
func (x Event) String() string {
	if str, ok := _EventMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Event(%d)", x)
}

var _EventValue = map[string]Event{
	_EventName[0:7]:                    0,
	strings.ToLower(_EventName[0:7]):   0,
	_EventName[7:14]:                   1,
	strings.ToLower(_EventName[7:14]):  1,
	_EventName[14:21]:                  2,
	strings.ToLower(_EventName[14:21]): 2,
	_EventName[21:30]:                  3,
	strings.ToLower(_EventName[21:30]): 3,
	_EventName[30:37]:                  4,
	strings.ToLower(_EventName[30:37]): 4,
	_EventName[37:44]:                  5,
	strings.ToLower(_EventName[37:44]): 5,
	_EventName[44:51]:                  6,
	strings.ToLower(_EventName[44:51]): 6,
	_EventName[51:59]:                  7,
	strings.ToLower(_EventName[51:59]): 7,
	_EventName[59:65]:                  8,
	strings.ToLower(_EventName[59:65]): 8,
	_EventName[65:71]:                  9,
	strings.ToLower(_EventName[65:71]): 9,
}

// ParseEvent attempts to convert a string to a Event
func ParseEvent(name string) (Event, error) {
	if x, ok := _EventValue[name]; ok {
		return x, nil
	}
	return Event(0), fmt.Errorf("%s is not a valid Event", name)
}

// MarshalText implements the text marshaller method
func (x *Event) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Event) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseEvent(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
