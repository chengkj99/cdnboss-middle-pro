package alarm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func Config(rg *echo.Group) {
	rg.GET("/config", getConfigAction)
}

func getConfigAction(c echo.Context) error {
	MetricUrl := host + "/v1/metric/all"
	GroupMetricUrl := host + "/v1/groupmetric/all"
	RuleMetricUrl := host + "/v1/rule/all"

	// fmt.Println(data)

	MetricRes := getResOfByte(c, MetricUrl)
	GroupMetricRes := getResOfByte(c, GroupMetricUrl)
	RuleMetricRes := getResOfByte(c, RuleMetricUrl)

	var metricD map[string][]Metric
	var GroupMetricD []GroupMetric
	var RuleMetricD []RuleMetric

	err := json.Unmarshal(MetricRes, &metricD)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}
	err = json.Unmarshal(GroupMetricRes, &GroupMetricD)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}
	err = json.Unmarshal(RuleMetricRes, &RuleMetricD)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}

	MetricAll := &MetricAll{metricD, GroupMetricD, RuleMetricD}

	c.JSON(200, MetricAll)

	return nil
}

// MetricAll 数据整合
type MetricAll struct {
	Metric      map[string][]Metric
	GroupMetric []GroupMetric
	RuleMetric  []RuleMetric
}

// Metric 来源系统下的监控项
type Metric struct {
	ID           int    `json:"Id"`
	MetricName   string `json:"MetricName"`
	SystemName   string `json:"SystemName"`
	Describe     string `json:"Describe"`
	ConfirmState string `json:"ConfirmState"`
}

// GroupMetric 监控项和用户组的关系
type GroupMetric struct {
	GroupID    int                 `json:"GroupId"`
	GroupName  string              `json:"GroupName"`
	SysMetrics map[string][]string `json:"SysMetrics"`
}

// RuleMetric 监控项和升级规则的关系
type RuleMetric struct {
	ID          int      `json:"Id"`
	Metric      string   `json:"Metric"`
	OldType     int      `json:"OldType"`
	SystemName  string   `json:"SystemName"`
	Endpoint    string   `json:"Endpoint"`
	Type        int      `json:"Type"`
	CreateTime  string   `json:"CreateTime"`
	UpdateTime  string   `json:"UpdateTime"`
	Period      []string `json:"Period"`
	TargetLevel []string `json:"TargetLevel"`
	Value       []string `json:"Value"`
}

func getResOfByte(c echo.Context, url string) []byte {
	req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger().Errorf("error: %v", err)
	}
	return data
}
