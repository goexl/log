package log_test

import (
	"testing"

	"github.com/goexl/log"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, log.New().Apply())
}
