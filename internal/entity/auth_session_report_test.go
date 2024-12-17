package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession_Report(t *testing.T) {
	m := FindSessionByRefID("sessxkkcabcd")

	r, _ := m.Report(false)
	assert.GreaterOrEqual(t, len(r), 1)

	r2, _ := m.Report(true)
	assert.GreaterOrEqual(t, len(r2), 1)
}
