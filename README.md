# Prometheus exporter for Confluent Cloud Metrics API

A simple prometheus exporter that can be used to extract metrics from [Confluent Cloud Metric API](https://docs.confluent.io/current/cloud/metrics-api.html).
By default, the exporter will be exposing the metrics on [port 2112](http://localhost:2112)
To use the exporter, the following environment variables need to be specified:


* `CCLOUD_APIKEY`: Your Confluent Cloud base64 encoded API ID-Secret

`CCLOUD_APIKEY` environment variable will be used to invoke the https://api.telemetry.confluent.cloud endpoint.

## Usage
```
./ccloudexporter -cluster <cluster_id>
````

## Examples

### Building and executing
```
go get github.com/Vinsser/ccloudexporter/cmd/ccloudexporter
go install github.com/Vinsser/ccloudexporter/cmd/ccloudexporter
export CCLOUD_APIKEY=<apikey>
./ccloudexporter -cluster lkc-abc123
```

### Using docker
```
docker run \
  -e CCLOUD_APIKEY=$CCLOUD_APIKEY \
  -p 2112:2112
  docvin55/ccloudexporter:0.1.1 -cluster lkc-abc123

```

## How to build
```
go get github.com/Vinsser/ccloudexporter/cmd/ccloudexporter
```

