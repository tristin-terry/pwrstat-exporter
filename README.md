# Pwrstat Prometheus Exporter
A Prometheus exporter for PowerPanel Personal Linux (pwrstat cli) built in Go.

## Installation
* Install and configure [PowerPanel Personal Linux](https://www.cyberpowersystems.com/product/software/power-panel-personal/powerpanel-for-linux/)
* run `pwrstat_exporter`
  * optionally set it up as a service by creating `/etc/systemd/system/pwrstat_exporter.service`:   
``` 
[Unit]
Description=Pwrstat Exporter
After=network-online.target

[Service]
Type=simple
User=root
Group=root
ExecStart=/usr/local/bin/pwrstat_exporter
SyslogIdentifier=pwrstat_exporter
Restart=always

[Install]
WantedBy=multi-user.target
```
* `curl localhost:10100/metrics` to verify it is working
* setup prometheus, example scrape config:

``` 
- job_name: pwrstat_exporter
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:10100
```

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

## Sample Readings
```
# HELP pwrstat_battery_capacity_percent 
# TYPE pwrstat_battery_capacity_percent gauge
pwrstat_battery_capacity_percent 1
# HELP pwrstat_battery_remaining_seconds 
# TYPE pwrstat_battery_remaining_seconds gauge
pwrstat_battery_remaining_seconds 2832
# HELP pwrstat_battery_voltage 
# TYPE pwrstat_battery_voltage gauge
pwrstat_battery_voltage 20.7
# HELP pwrstat_diagnostic_result 
# TYPE pwrstat_diagnostic_result gauge
pwrstat_diagnostic_result 1
# HELP pwrstat_info 
# TYPE pwrstat_info gauge
pwrstat_info{model_name="CP 1500C"} 0
# HELP pwrstat_input_rating_voltage 
# TYPE pwrstat_input_rating_voltage gauge
pwrstat_input_rating_voltage 120
# HELP pwrstat_is_ac_present 
# TYPE pwrstat_is_ac_present gauge
pwrstat_is_ac_present 2832
# HELP pwrstat_is_battery_charging 
# TYPE pwrstat_is_battery_charging gauge
pwrstat_is_battery_charging 0
# HELP pwrstat_is_battery_discharging 
# TYPE pwrstat_is_battery_discharging gauge
pwrstat_is_battery_discharging 0
# HELP pwrstat_load_percent Percent load on the UPS
# TYPE pwrstat_load_percent gauge
pwrstat_load_percent 0.12
# HELP pwrstat_load_watts Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)
# TYPE pwrstat_load_watts gauge
pwrstat_load_watts 108
# HELP pwrstat_model_name model_name
# TYPE pwrstat_model_name summary
pwrstat_model_name_sum 0
pwrstat_model_name_count 0
# HELP pwrstat_output_rating_watts Watt rating of the UPS
# TYPE pwrstat_output_rating_watts gauge
pwrstat_output_rating_watts 900
# HELP pwrstat_output_voltage 
# TYPE pwrstat_output_voltage gauge
pwrstat_output_voltage 120
# HELP pwrstat_utility_voltage 
# TYPE pwrstat_utility_voltage gauge
pwrstat_utility_voltage 120
```