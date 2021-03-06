package genctx

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/engine"
)

const DateTimeFormatDefault string = "2006-01-02T15:04:05-07:00"

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Output{})

// Activity is an Activity that is used to flow context
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New function for the activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	err := metadata.MapToStruct(ctx.Settings(), &Settings{}, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{}

	return act, nil
}

func generateUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	output := &Output{}
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if input.Flow != "" {
		output.Flow = input.Flow
	} else {
		output.Flow = ctx.ActivityHost().Name()
	}
	output.FlowID = ctx.ActivityHost().ID()
	output.ApplicationName = engine.GetAppName()
	output.ApplicationVersion = engine.GetAppVersion()
	if input.ExternalID != "" {
		output.ExternalID = input.ExternalID
	} else {
		output.ExternalID = ""
	}
	output.RequestID = generateUUID()
	output.RequestTimestamp = time.Now().UTC().Format(DateTimeFormatDefault)
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
