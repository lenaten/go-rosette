package rosette

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntities(t *testing.T) {
	c := client()
	content := "Bill Murray will appear in new Ghostbusters film: Dr. Peter Venkman was spotted filming a cameo in Boston this… http://dlvr.it/BnsFfS"
	in := c.NewEntitiesInput(content)
	out, err := c.Entities(in)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(out.Entities))
	entity := out.Entities[0]
	assert.NotEqual(t, "", entity.Type)
	assert.NotEqual(t, 0, entity.Count)
}
