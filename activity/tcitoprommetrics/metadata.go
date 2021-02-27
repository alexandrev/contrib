package tcitoprommetrics

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

// Settings for the activity
type Settings struct {
}

// Input for the activity
type Input struct {
	Metrics []AutoGenerated `json:"appMetrics"`
}

type AutoGenerated struct {
	App                App                  `json:"app"`
	AppInstanceMetrics []AppInstanceMetrics `json:"appInstanceMetrics"`
	AppMetrics         AppMetrics           `json:"appMetrics"`
}
type App struct {
	AppID                string `json:"appId"`
	AppName              string `json:"appName"`
	AppType              string `json:"appType"`
	Category             string `json:"category"`
	CreatedTime          int    `json:"createdTime"`
	DeploymentStage      string `json:"deploymentStage"`
	DeploymentType       string `json:"deploymentType"`
	DesiredInstanceCount int    `json:"desiredInstanceCount"`
	EndpointVisibility   string `json:"endpointVisibility"`
	LastStartedTime      int64  `json:"lastStartedTime"`
	ModifiedTime         int64  `json:"modifiedTime"`
}

type Flows struct {
	AvgExecTime float64 `json:"avg_exec_time"`
	Completed   int     `json:"completed"`
	Failed      int     `json:"failed"`
	FlowName    string  `json:"flow_name"`
	MaxExecTime float64 `json:"max_exec_time"`
	MinExecTime float64 `json:"min_exec_time"`
	Started     int     `json:"started"`
}
type Config struct {
	Interval     string `json:"interval"`
	IntervalUnit string `json:"interval_unit"`
	Repeating    string `json:"repeating"`
}

type Triggers struct {
	Completed   int    `json:"completed"`
	Failed      int    `json:"failed"`
	Started     int    `json:"started"`
	Status      string `json:"status"`
	TriggerName string `json:"trigger_name"`
}
type AppInstanceMetricsIn struct {
	AppName    string     `json:"app_name"`
	AppVersion string     `json:"app_version"`
	Flows      []Flows    `json:"flows"`
	Triggers   []Triggers `json:"triggers"`
}
type AppInstanceMetrics struct {
	AppInstance        string               `json:"appInstance"`
	AppInstanceMetrics AppInstanceMetricsIn `json:"appInstanceMetrics"`
}
type Labels struct {
	Status string `json:"status"`
}
type TciAppExecutions struct {
	Labels Labels `json:"labels"`
	Value  int    `json:"value"`
}
type TciAppInstancesCPU struct {
	Labels Labels  `json:"labels"`
	Value  float64 `json:"value"`
}
type TciAppInstancesMemory struct {
	Labels Labels  `json:"labels"`
	Value  float64 `json:"value"`
}
type TciAppSinceLastExecution struct {
	Value float64 `json:"value"`
}
type AppMetrics struct {
	TciAppExecutions         []TciAppExecutions         `json:"tci_app_executions"`
	TciAppInstancesCPU       []TciAppInstancesCPU       `json:"tci_app_instances_cpu"`
	TciAppInstancesMemory    []TciAppInstancesMemory    `json:"tci_app_instances_memory"`
	TciAppSinceLastExecution []TciAppSinceLastExecution `json:"tci_app_since_last_execution"`
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

// ToMap for Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"appMetrics": i.Metrics,
	}
}

// FromMap for input
func (i *Input) FromMap(values map[string]interface{}) error {
	for i, k := range values {
		println(i)
		println(k)
		println(fmt.Sprintf("%T", k))

	}
	var err bool
	i.Metrics, err = values["appMetrics"].([]AutoGenerated)

	println(i.Metrics)
	println(err)

	return nil
}

// Output for the activity
type Output struct {
	Data string `json:"data"`
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