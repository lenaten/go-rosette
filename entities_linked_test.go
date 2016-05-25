package rosette

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntitiesLinked(t *testing.T) {
	c := client()
	content := "Apple"
	in := c.NewEntitiesLinkedInput(content)
	out, err := c.EntitiesLinked(in)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(out.Entities))
	entity := out.Entities[0]
	assert.Equal(t, "Q312", entity.EntityId)
}
