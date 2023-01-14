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

func toFQName(name string) string {
	return prometheus.BuildFQName(*namespace, "", name)
}

func newGauge(name string, help string) prometheus.Gauge {
	return promauto.With(reg).NewGauge(prometheus.GaugeOpts{
		Name: toFQName(name),
		Help: help,
	})
}

var (
	state = newGauge("state", "")
	info  = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: toFQName("info"),
		Help: "",
	}, []string{"model_name"})
	reg                          = prometheus.NewRegistry()
	battery_voltage              = newGauge("battery_voltage", "")
	input_rating_voltage         = newGauge("input_rating_voltage", "")
	output_rating_watts          = newGauge("output_rating_watts", "Watt rating of the UPS")
	is_avr_supported             = newGauge("is_avr_supported", "")
	is_online_type               = newGauge("is_online_type", "")
	diagnostic_result            = newGauge("diagnostic_result", "")
	diagnostic_date_unix         = newGauge("diagnostic_date_unix", "")
	power_event_result           = newGauge("power_event_result", "")
	power_event_date_unix        = newGauge("power_event_date_unix", "")
	power_event_duration_seconds = newGauge("power_event_duration_seconds", "")
	battery_remaining_seconds    = newGauge("battery_remaining_seconds", "Remaing runtime in seconds")
	is_battery_charging          = newGauge("is_battery_charging", "")
	is_battery_discharging       = newGauge("is_battery_discharging", "")
	is_ac_present                = newGauge("is_ac_present", "")
	is_boost                     = newGauge("is_boost", "")
	utility_voltage              = newGauge("utility_voltage", "")
	output_voltage               = newGauge("output_voltage", "")
	load_percent                 = newGauge("load_percent", "Percent load on the UPS")
	battery_capacity_percent     = newGauge("battery_capacity_percent", "")
	load_watts                   = newGauge("load_watts", "Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)")
)

func writer(c net.Conn) {
	for {
		_, err := c.Write([]byte("STATUS\n\n"))
		if err != nil {
			log.Fatal("write error: ", err)
		}

		time.Sleep(time.Duration(*collectDelay) * time.Second)
	}
}

func reader(r io.Reader) {
	buf := make([]byte, 512)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			log.Fatal("read error: ", err)
		}
		result := string(buf[0:n])

		if *debug {
			fmt.Println("Read from socket:", result)
		}

		parsedResult := pwrstat_parser.ParseToResult(result)

		if *debug {
			fmt.Println("Parsed result:", parsedResult)
		}

		updatePrometheus(parsedResult)
	}
}

func updatePrometheus(parsedResults pwrstat_parser.PwrstatResult) {
	state.Set(parsedResults.State)
	info.WithLabelValues(parsedResults.ModelName)
	battery_voltage.Set(parsedResults.BatteryVoltage)
	input_rating_voltage.Set(parsedResults.InputRatingVoltage)
	output_rating_watts.Set(parsedResults.OutputRatingWatts)
	is_avr_supported.Set(parsedResults.IsAvrSupported)
	is_online_type.Set(parsedResults.IsOnlineType)
	diagnostic_result.Set(parsedResults.DiagnosticResult)
	diagnostic_date_unix.Set(parsedResults.DiagnosticDateUnix)
	power_event_result.Set(parsedResults.PowerEventResult)
	power_event_date_unix.Set(parsedResults.PowerEventDateUnix)
	power_event_duration_seconds.Set(parsedResults.PowerEventDurationSeconds)
	battery_remaining_seconds.Set(parsedResults.BatteryRemainingSeconds)
	is_battery_charging.Set(parsedResults.IsBatteryCharging)
	is_battery_discharging.Set(parsedResults.IsBatteryDischarging)
	is_ac_present.Set(parsedResults.IsAcPresent)
	is_boost.Set(parsedResults.IsBoost)
	utility_voltage.Set(parsedResults.UtilityVoltage)
	output_voltage.Set(parsedResults.OutputVoltage)
	load_percent.Set(parsedResults.LoadPercent)
	battery_capacity_percent.Set(parsedResults.BatteryCapacityPercent)
	load_watts.Set(parsedResults.LoadWatts)
}

var (
	port          = flag.String("port", "10100", "The address to listen on for HTTP requests.")
	namespace     = flag.String("namespace", "pwrstat", "Namespace prefix for prometheus metrics")
	collectDelay  = flag.Int("collect.delay", 15, "The delay between each sensor reading in seconds.")
	debug         = flag.Bool("debug", false, "Turn on debug logging.")
	pwrstatSocket = flag.String("pwrstat.socket", "/var/pwrstatd.ipc", "Where to find the pwrstatd.ipc socket.")
)

func main() {
	flag.Parse()

	reg.MustRegister(info)

	c, err := net.Dial("unix", *pwrstatSocket)
	if err != nil {
		log.Fatal("Could not connect to pwrstatd socket at: ", pwrstatSocket)
	}

	defer c.Close()
	go reader(c)
	go writer(c)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
