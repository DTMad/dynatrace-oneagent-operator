package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaults_OneAgentSpec(t *testing.T) {
	oa := newOneAgentSpec()
	SetDefaults_OneAgentSpec(oa)
	assert.NotNil(t, oa.WaitReadySeconds)
	assert.NotEmpty(t, oa.Image)
}

func newOneAgentSpec() *OneAgentSpec {
	return &OneAgentSpec{}
}
