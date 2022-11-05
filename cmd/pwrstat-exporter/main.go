package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tristin-terry/pwrstat-exporter/internal/pwrstat_parser"
)

const (
	namespace = "pwrstat"
)

func toFQName(name string) string {
	return prometheus.BuildFQName(namespace, "", name)
}

func newGauge(name string, help string) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Name: toFQName(name),
		Help: help,
	})
}

var (
	output_rating_watts       = newGauge("output_rating_watts", "Watt rating of the UPS")
	load_percent              = newGauge("load_percent", "Percent load on the UPS")
	load_watts                = newGauge("load_watts", "Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)")
	battery_remaining_seconds = newGauge("battery_remaining_seconds", "Remaing runtime in seconds")
	is_ac_present             = newGauge("is_ac_present", "")
	is_battery_charging       = newGauge("is_battery_charging", "")
	is_battery_discharging    = newGauge("is_battery_discharging", "")
	utility_voltage           = newGauge("utility_voltage", "")
	output_voltage            = newGauge("output_voltage", "")
	diagnostic_result         = newGauge("diagnostic_result", "")
	battery_capacity_percent  = newGauge("battery_capacity_percent", "")
	battery_voltage           = newGauge("battery_voltage", "")
	input_rating_voltage      = newGauge("input_rating_voltage", "")
	info                      = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: toFQName("info"),
		Help: "",
	}, []string{"model_name"})
)

func writer(c net.Conn) {
	for {
		_, err := c.Write([]byte("STATUS\n\n"))
		if err != nil {
			log.Fatal("write error:", err)
			return
		}

		time.Sleep(time.Duration(*collectDelay) * time.Second)
	}
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			log.Fatal("read error:", err)
			return
		}
		result := string(buf[0:n])
		parsedResult := pwrstat_parser.ParseToResult(result)

		if *debug {
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

var (
	listenAddress = flag.String("listen-address", "10100", "The address to listen on for HTTP requests.")
	collectDelay  = flag.Int("collect-delay", 15, "The delay between each sensor reading in seconds.")
	debug         = flag.Bool("debug", false, "Toggle debug logs.")
)

func main() {
	flag.Parse()

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
	log.Fatal(http.ListenAndServe(":"+*listenAddress, nil))
}
