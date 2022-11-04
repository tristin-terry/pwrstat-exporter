package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tristin-terry/pwrstat-prometheus-exporter/internal/pwrstat_parser"
)

const (
	debug = false
)

var (
	info = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pwrstat_info",
		Help: "",
	}, []string{"model_name"})
	output_rating_watts = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_output_rating_watts",
		Help: "Watt rating of the UPS",
	})
	load_percent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_load_percent",
		Help: "Percent load on the UPS",
	})
	load_watts = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_load_watts",
		Help: "Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)",
	})
	battery_remaining_seconds = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_battery_remaining_seconds",
		Help: "",
	})
	is_ac_present = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_is_ac_present",
		Help: "",
	})
	is_battery_charging = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_is_battery_charging",
		Help: "",
	})
	is_battery_discharging = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_is_battery_discharging",
		Help: "",
	})
	utility_voltage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_utility_voltage",
		Help: "",
	})
	output_voltage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_output_voltage",
		Help: "",
	})
	diagnostic_result = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_diagnostic_result",
		Help: "",
	})
	battery_capacity_percent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_battery_capacity_percent",
		Help: "",
	})
	battery_voltage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_battery_voltage",
		Help: "",
	})
	input_rating_voltage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pwrstat_input_rating_voltage",
		Help: "",
	})
	model_name = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "pwrstat_model_name",
		Help: "model_name",
	})
)

func writer(c net.Conn) {
	for {
		_, err := c.Write([]byte("STATUS\n\n"))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}

		time.Sleep(15 * time.Second)
	}
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		result := string(buf[0:n])
		parsedResult := pwrstat_parser.ParseToResult(result)

		if debug {
			fmt.Println(parsedResult)
		}

		updatePrometheus(parsedResult)
	}
}

func updatePrometheus(parsedResults pwrstat_parser.PwrstatResult) {
	info.WithLabelValues(parsedResults.ModelName)
	output_rating_watts.Set(parsedResults.OutputRatingWatts)
	load_percent.Set(parsedResults.LoadPercent)
	load_watts.Set(parsedResults.LoadWatts)
	battery_remaining_seconds.Set(parsedResults.BatteryRemainingSeconds)
	is_ac_present.Set(parsedResults.IsAcPresent)
	is_battery_charging.Set(parsedResults.IsBatteryCharging)
	is_battery_discharging.Set(parsedResults.IsBatteryDischarging)
	utility_voltage.Set(parsedResults.UtilityVoltage)
	output_voltage.Set(parsedResults.OutputVoltage)
	diagnostic_result.Set(parsedResults.DiagnosticResult)
	battery_capacity_percent.Set(parsedResults.BatteryCapacityPercent)
	battery_voltage.Set(parsedResults.BatteryVoltage)
	input_rating_voltage.Set(parsedResults.InputRatingVoltage)
}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(info)
	prometheus.MustRegister(model_name)
	prometheus.MustRegister(output_rating_watts)
	prometheus.MustRegister(load_percent)
	prometheus.MustRegister(load_watts)
	prometheus.MustRegister(battery_remaining_seconds)
	prometheus.MustRegister(is_ac_present)
	prometheus.MustRegister(is_battery_charging)
	prometheus.MustRegister(is_battery_discharging)
	prometheus.MustRegister(utility_voltage)
	prometheus.MustRegister(output_voltage)
	prometheus.MustRegister(diagnostic_result)
	prometheus.MustRegister(battery_capacity_percent)
	prometheus.MustRegister(battery_voltage)
	prometheus.MustRegister(input_rating_voltage)
}

func main() {
	c, err := net.Dial("unix", "/var/pwrstatd.ipc")
	if err != nil {
		panic(err)
	}

	defer c.Close()
	go reader(c)
	go writer(c)

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":10100", nil))
}
