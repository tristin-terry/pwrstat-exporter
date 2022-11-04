# Pwrstat Prometheus Exporter
A Prometheus exporter for PowerPanel Personal Linux (pwrstat cli) built in Go.

See sample output [here](/../../blob/main/samples/SAMPLE_METRICS.md)

## Installation
* Install and configure [PowerPanel Personal Linux](https://www.cyberpowersystems.com/product/software/power-panel-personal/powerpanel-for-linux/)
* Download `pwrstat-exporter` from the most recent build artifacts or build locally (`make build`)
* Run `pwrstat-exporter`
* `curl localhost:10100/metrics` to verify it is working
* Setup prometheus, example scrape config:
``` 
- job_name: pwrstat-exporter
  scrape_interval: 15s
  static_configs:
  - targets:
    - localhost:10100
```
* A sample Grafana dashboard is available at `/grafana/export.json`
* Optionally set it up as a service by creating `/etc/systemd/system/pwrstat-exporter.service`:   
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
  * Run:
```
systemctl daemon-reload
systemctl enable pwrstat-exporter.service
systemctl start pwrstat-exporter.service
systemctl status pwrstat-exporter.service
```

# Arguments
```
pwrstat-exporter --help
Usage of pwrstat-exporter:
  -collect-delay int
        The delay between each sensor reading in seconds. (default 15)
  -listen-address string
        The address to listen on for HTTP requests. (default "10100")
```

## Local Development
* Install [Go](https://go.dev/doc/install)
* Clone the repository
* Available makefile commands:
  * `make run` to run the app locally
    * note: the server runs on `localhost:10101/metrics` to not collide with the default
  * `make test` to run tests
  * `make build` to build an executable