package entity

import (
	"encoding/json"
	"fmt"
)

type Category uint

const (
	CategoryTech Category = iota + 1
	CategoryBusiness
	CategoryMarketing
)

func (c Category) String() string {
	switch c {
	case CategoryTech:
		return "tech"
	case CategoryBusiness:
		return "business"
	case CategoryMarketing:
		return "marketing"
	default:
		return "unknown"
	}
}

func (c Category) IsValid() bool {
	return c >= CategoryTech && c <= CategoryMarketing
}

func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "tech":
		*c = CategoryTech
	case "business":
		*c = CategoryBusiness
	case "marketing":
		*c = CategoryMarketing
	default:
		return fmt.Errorf("invalid category: %s", s)
	}
	return nil
}
