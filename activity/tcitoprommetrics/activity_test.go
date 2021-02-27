/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package tcitoprommetrics

import (
	"testing"

	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestPromMetrics(t *testing.T) {

	var a []AppMetrics

	var b AppMetrics

	b.App.AppName = "test"
	b.App.AppType = "flogo"
	b.App.AppID = "id1"

	var c ApplicationInstanceMetrics

	c.AppInstanceMetrics.AppName = "test"
	c.AppInstanceMetrics.AppVersion = "1.0"

	var d FlowData

	d.Completed = 1
	d.Failed = 0
	d.Started = 1
	d.AvgExecTime = 1.1
	d.MaxExecTime = 1.2
	d.MinExecTime = 1.0

	c.AppInstanceMetrics.Flows = append(c.AppInstanceMetrics.Flows, d)

	b.AppInstanceMetrics = append(b.AppInstanceMetrics, c)

	a = append(a, b)

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(nil, mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	input := &Input{
		Metrics: a,
	}
	tc.SetInputObject(input)

	act.Eval(tc)

	output := &Output{}
	tc.GetOutputObject(output)
	assert.NotEmpty(t, output.Data)

	println(output.Data)

}
