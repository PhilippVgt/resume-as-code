package model

import (
	"encoding/json"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return nil
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t Date) MarshalYAML() (interface{}, error) {
	return t.Time.Format("2006-01-02"), nil
}

func (t *Date) UnmarshalJSON(data []byte) error {
	var buf string
	err := json.Unmarshal(data, &buf)
	if err != nil {
		return nil
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("2006-01-02"))
}
