package genctx

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(nil, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	println(output.ApplicationName)
	println(output.RequestID)
	println(output.ExternalID)
	println(output.RequestTimestamp)
	assert.NotEmpty(t, output.RequestID)
}
