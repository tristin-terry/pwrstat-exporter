package pwrstat_parser

import (
	"testing"
)

const (
	sampleUps1500 = `STATUS
	state=0
	model_name= CP 1500C
	battery_volt=20600
	input_rating_volt=120000
	output_rating_watt=900000
	avr_supported=yes
	online_type=no
	diagnostic_result=1
	diagnostic_date=2022/11/02 19:12:21
	power_event_result=1
	power_event_date=2022/11/02 19:29:14
	power_event_during=4 sec.
	battery_remainingtime=3030
	battery_charging=no
	battery_discharging=no
	ac_present=yes
	boost=no
	utility_volt=121000
	output_volt=121000
	load=10000
	battery_capacity=100`
)

func TestParser(t *testing.T) {
	result := Parse(sampleUps1500)

	if len(result) != 21 {
		t.Fail()
	}

	if result["battery_capacity"] != "100" {
		t.Fail()
	}
}
