// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package opsworksiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/opsworks"
	"github.com/awslabs/aws-sdk-go/service/opsworks/opsworksiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*opsworksiface.OpsWorksAPI)(nil), opsworks.New(nil))
}