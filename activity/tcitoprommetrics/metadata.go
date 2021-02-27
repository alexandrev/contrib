package tcitoprommetrics

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
}

// Input for the activity
type Input struct {
	Metrics []AppMetrics `md:"appMetrics"`
}

//AppData object
type AppData struct {
	AppID                string `md:"appId"`
	AppName              string `md:"appName"`
	AppType              string `md:"appType"`
	Category             string `md:"category"`
	CreatedTime          int    `md:"createdTime"`
	DeploymentStage      string `md:"deploymentStage"`
	DeploymentType       string `md:"deploymentType"`
	DesiredInstanceCount int    `md:"desiredInstanceCount"`
	EndpointVisibility   string `md:"endpointVisibility"`
	LastStartedTime      int64  `md:"lastStartedTime"`
	ModifiedTime         int64  `md:"modifiedTime"`
}

type ApplicationMetrics struct {
	TciAppExecutions []struct {
		Labels struct {
			Status string `md:"status"`
		} `md:"labels"`
		Value int `md:"value"`
	} `md:"tci_app_executions"`
	TciAppInstancesCPU []struct {
		Labels struct {
			Status string `md:"status"`
		} `md:"labels"`
		Value float64 `md:"value"`
	} `md:"tci_app_instances_cpu"`
	TciAppInstancesMemory []struct {
		Labels struct {
			Status string `md:"status"`
		} `md:"labels"`
		Value float64 `md:"value"`
	} `md:"tci_app_instances_memory"`
	TciAppSinceLastExecution []struct {
		Value int `md:"value"`
	} `md:"tci_app_since_last_execution"`
}

type FlowData struct {
	AvgExecTime float64 `md:"avg_exec_time"`
	Completed   int     `md:"completed"`
	Failed      int     `md:"failed"`
	FlowName    string  `md:"flow_name"`
	MaxExecTime float64 `md:"max_exec_time"`
	MinExecTime float64 `md:"min_exec_time"`
	Started     int     `md:"started"`
}

type ApplicationInstanceMetrics struct {
	AppInstance        string `md:"appInstance"`
	AppInstanceMetrics struct {
		AppName    string     `md:"app_name"`
		AppVersion string     `md:"app_version"`
		Flows      []FlowData `md:"flows"`
		Triggers   []struct {
			Completed int `md:"completed"`
			Failed    int `md:"failed"`
			Handlers  struct {
				Test struct {
					Completed int `md:"completed"`
					Config    struct {
						Interval     string `md:"interval"`
						IntervalUnit string `md:"interval_unit"`
						Repeating    string `md:"repeating"`
					} `md:"config"`
					Failed      int    `md:"failed"`
					HandlerName string `md:"handler_name"`
					Started     int    `md:"started"`
				} `md:"test"`
			} `md:"handlers"`
			Started     int    `md:"started"`
			Status      string `md:"status"`
			TriggerName string `md:"trigger_name"`
		} `md:"triggers"`
	}
}

type Sample struct {
	Labels map[string]string
	Value  float64
}

type MetricType struct {
	Name        string
	Description string
	Type        string
	Samples     []Sample
}

type MetricList struct {
	Metrics []*MetricType
}

func (r *MetricList) Create(name string, description string, metrictype string) *MetricType {
	metric := r.Get(name)
	if metric == nil {
		m := MetricType{Name: name, Description: description, Type: metrictype}
		r.Metrics = append(r.Metrics, &m)
		return &m
	}
	return metric
}

func (r MetricList) Get(name string) *MetricType {
	for _, metric := range r.Metrics {
		if metric.Name == name {
			return metric
		}
	}
	return nil
}

func (r *MetricType) Add(labels map[string]string, value float64) bool {
	s := Sample{labels, value}
	r.Samples = append(r.Samples, s)
	return true
}

//AppMetrics object as input for the activity
type AppMetrics struct {
	App                AppData                      `md:"app"`
	AppInstanceMetrics []ApplicationInstanceMetrics `md:"appInstanceMetrics"`
	AppMetrics         ApplicationMetrics           `md:"appInstanceMetrics"`
}

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"appMetrics": i.Metrics,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Metrics, _ = values["appMetrics"].([]AppMetrics)
	return nil
}

// Output for the activity
type Output struct {
	Data string `md:"data"`
}

// ToMap conver to object
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

// FromMap convert to object
func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.Data, err = coerce.ToString(values["data"])
	if err != nil {
		return err
	}

	return nil
}
