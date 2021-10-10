package memorystore

import (
	"testing"

	"github.com/adlerhsieh/simple_search_engine/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGet(t *testing.T) {
	m := New()
	document := &entities.Document{
		ID:      "foo",
		Content: "bar",
	}

	err := m.Set(document)
	require.NoError(t, err)

	d, err := m.Get(document.ID)
	require.NoError(t, err)
	assert.Equal(t, "bar", d.Content)
}
