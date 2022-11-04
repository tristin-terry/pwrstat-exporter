# Pwrstat Prometheus Exporter
A Prometheus exporter for PowerPanel Personal Linux (pwrstat cli) built in Go.

See sample output [here](/../../blob/main/samples/SAMPLE_METRICS.md)

## Installation
* Install and configure [PowerPanel Personal Linux](https://www.cyberpowersystems.com/product/software/power-panel-personal/powerpanel-for-linux/)
* run `pwrstat-exporter`
* `curl localhost:10100/metrics` to verify it is working
* setup prometheus, example scrape config:
``` 
- job_name: pwrstat-exporter
  scrape_interval: 15s
  static_configs:
  - targets:
    - localhost:10100
```
* optionally set it up as a service by creating `/etc/systemd/system/pwrstat-exporter.service`:   
``` 
[Unit]
Description=Pwrstat Exporter
After=network-online.target

[Service]
Type=simple
User=root
Group=root
ExecStart=/usr/local/bin/pwrstat-exporter
SyslogIdentifier=pwrstat-exporter
Restart=always

[Install]
WantedBy=multi-user.target
```
  * `systemctl daemon-reload`
  * `systemctl enable pwrstat-exporter.service`
  * `systemctl start pwrstat-exporter.service`
  * to verify it is running: `systemctl | grep pwrstat`

## Usage
* `pwrstat_exporter` starts an http server on port `10100` with endpoints:
  * `GET /metrics`

## Local Development
* Install [Go](https://go.dev/doc/install)
* Clone the repository
* Available makefile commands:
  * `make run` to run the app locally
  * `make test` to run tests
  * `make build` to build an executable