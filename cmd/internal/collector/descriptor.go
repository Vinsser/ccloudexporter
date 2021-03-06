package collector

//
// descriptor.go
// Copyright (C) 2020 gaspar_d </var/spool/mail/gaspar_d>
//
// Distributed under terms of the MIT license.
//

import "strings"
import "errors"
import "os"
import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"

// Response from Confluent Cloud API metric endpoint
// This is the JSON structure for the endpoint
// https://api.telemetry.confluent.cloud/v1/metrics/cloud/descriptors
type DescriptorResponse struct {
	Data []MetricDescription `json:"data"`
}

// Metric from the  https://api.telemetry.confluent.cloud/v1/metrics/cloud/descriptors
// response
type MetricDescription struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Unit        string        `json:"unit"`
	Description string        `json:"description"`
	Labels      []MetricLabel `json:"labels"`
}

// Label of a metric, should contain a key and a description
// e.g.
//  {
//      "description": "Name of the Kafka topic",
//      "key": "topic"
//  }
type MetricLabel struct {
	Key         string `json:"key"`
	Description string `json:"description"`
}

var (
	excludeListForMetric = map[string]string{
		"io.confluent.kafka.server": "",
		"delta":                     "",
	}
	descriptorUri = "/v1/metrics/cloud/descriptors"
)

// Return true if the metric has this label
func (metric MetricDescription) hasLabel(label string) bool {
	for _, l := range metric.Labels {
		if l.Key == label {
			return true
		}
	}
	return false
}

// Return a human friendly metric name from a Confluent Cloud API metric
func GetNiceNameForMetric(metric MetricDescription) string {
	splits := strings.Split(metric.Name, "/")
	for _, split := range splits {
		_, contain := excludeListForMetric[split]
		if !contain {
			return split
		}
	}

	panic(errors.New("Invalid metric: " + metric.Name))
}

// Call the https://api.telemetry.confluent.cloud/v1/metrics/cloud/descriptors endpoint
// to retrieve the list of metrics
func SendDescriptorQuery() DescriptorResponse {
	// user, present := os.LookupEnv("CCLOUD_USER")
	// if !present || user == "" {
	// 	fmt.Print("CCLOUD_USER environment variable has not been specified")
	// 	os.Exit(1)
	// }
	// password, present := os.LookupEnv("CCLOUD_PASSWORD")
	// if !present || password == "" {
	// 	fmt.Print("CCLOUD_PASSWORD environment variable has not been specified")
	// 	os.Exit(1)
	// }
	apikey, present := os.LookupEnv("CCLOUD_APIKEY")
	if !present || apikey == "" {
		fmt.Print("CCLOUD_APIKEY environment variable has not been specified")
		os.Exit(1)
	}

	endpoint := HttpBaseUrl + descriptorUri
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	// req.SetBasicAuth(user, password)
	req.Header.Set("Authorization", "Basic " + apikey)
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		fmt.Printf("Received status code %d instead of 200", res.StatusCode)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	response := DescriptorResponse{}
	json.Unmarshal(body, &response)

	return response
}
