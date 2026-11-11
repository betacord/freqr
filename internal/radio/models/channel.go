// Radio protocol-related models implementations.

package models

import (
	"fmt"
	"strconv"
	"strings"
)

type modType string

const (
	AM modType = "AM"
	FM         = "FM"
)

// A struct representing frequency channel.
type Channel struct {
	Index    int
	Name     string
	Freq     string
	Mod      modType
	Delay    bool
	Lockout  bool
	Priority bool
}

// A function creating a new `Channel` structure object.
func newChannel(
	index int,
	name string,
	freq string,
	mod modType,
	delay bool,
	lockout bool,
	priority bool,
) *Channel {
	return &Channel{
		Index:    index,
		Name:     name,
		Freq:     freq,
		Mod:      mod,
		Delay:    delay,
		Lockout:  lockout,
		Priority: priority,
	}
}

// A function creating a new `Channel` structure object from text device response.
func NewChannelFromText(text string) (*Channel, error) {
	parts := strings.Split(text, ",")
	chId, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	delay := parts[5] == "1"
	lockout := parts[6] == "1"
	priority := parts[7] == "1"

	ch := newChannel(
		chId,
		parts[2],
		parts[3],
		modType(parts[4]),
		delay,
		lockout,
		priority,
	)

	return ch, nil
}

func NewChannelFromCSVRecord(record []string) (*Channel, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	return newChannel(
		id,
		record[1],
		record[2],
		modType(record[3]),
		strToBool(record[4]),
		strToBool(record[5]),
		strToBool(record[6]),
	), nil
}

// A function returning text representation of frequency channel.
func (c *Channel) ToString() string {
	return fmt.Sprintf(
		"CIN,%d,%s,%s,%s,,%d,%d,%d",
		c.Index,
		c.Name,
		c.Freq,
		c.Mod,
		boolToInt(c.Delay),
		boolToInt(c.Lockout),
		boolToInt(c.Priority),
	)
}

func (c *Channel) ToCSVRow() []string {
	return []string{
		strconv.Itoa(c.Index),
		c.Name,
		c.Freq,
		string(c.Mod),
		fmt.Sprintf("%v", c.Delay),
		fmt.Sprintf("%v", c.Lockout),
		fmt.Sprintf("%v", c.Priority),
	}
}

func boolToInt(val bool) uint8 {
	if val == false {
		return 0
	} else {
		return 1
	}
}

func strToBool(val string) bool {
	return strings.ToLower(val) == "true"
}
