package collection

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestUniqBy(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	list := []User{
		{
			Name: "a",
			Age:  0,
		},
		{
			Name: "b",
			Age:  1,
		},
		{
			Name: "b",
			Age:  1,
		},
	}

	result := lo.UniqBy(list, func(i User) string {
		return i.Name
	})

	strings := lo.Map(result, func(item User, index int) string {
		return item.Name
	})

	assert.Equal(t, []string{"a", "b"}, strings, "user name repeat")
}
