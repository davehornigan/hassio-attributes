package scalars

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalTime(t)
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("time must be string")
	}
	return time.Parse(time.RFC3339, str)
}
