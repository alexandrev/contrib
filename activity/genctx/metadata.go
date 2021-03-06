package genctx

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	Flow       string `md:"flow"`
	ExternalID string `md:"externalID"`
}

type Output struct {
	RequestID          string `md:"requestID"`
	RequestTimestamp   string `md:"requestTimestamp"`
	Flow               string `md:"flow"`
	FlowID             string `md:"flowID"`
	ApplicationName    string `md:"applicationName"`
	ApplicationVersion string `md:"applicationVersion"`
	ExternalID         string `md:"externalID"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"flow":               o.Flow,
		"flowID":             o.FlowID,
		"applicationName":    o.ApplicationName,
		"applicationVersion": o.ApplicationVersion,
		"requestID":          o.RequestID,
		"externalID":         o.ExternalID,
		"requestTimestamp":   o.RequestTimestamp,
	}
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"flow":       i.Flow,
		"externalID": i.ExternalID,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Flow, err = coerce.ToString(values["flow"])
	if err != nil {
		return err
	}
	i.ExternalID, err = coerce.ToString(values["externalID"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Flow, err = coerce.ToString(values["flow"])
	if err != nil {
		return err
	}
	o.FlowID, err = coerce.ToString(values["flowID"])
	if err != nil {
		return err
	}
	o.ApplicationName, err = coerce.ToString(values["applicationName"])
	if err != nil {
		return err
	}
	o.ApplicationVersion, err = coerce.ToString(values["applicationVersion"])
	if err != nil {
		return err
	}

	o.ExternalID, err = coerce.ToString(values["externalID"])
	if err != nil {
		return err
	}

	o.RequestTimestamp, err = coerce.ToString(values["requestTimestamp"])
	if err != nil {
		return err
	}

	o.RequestID, err = coerce.ToString(values["requestID"])
	if err != nil {
		return err
	}

	return nil
}
