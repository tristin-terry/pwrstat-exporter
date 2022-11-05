package pwrstat_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	battery_charging=yes
	battery_discharging=no
	ac_present=yes
	boost=no
	utility_volt=121000
	output_volt=119000
	load=10000
	battery_capacity=100`
)

func TestParser(t *testing.T) {
	expected := PwrstatResult{
		ModelName:               "CP 1500C",
		BatteryVoltage:          20.6,
		InputRatingVoltage:      120,
		OutputRatingWatts:       900,
		DiagnosticResult:        1,
		LoadPercent:             0.1,
		LoadWatts:               90,
		BatteryRemainingSeconds: 3030,
		IsAcPresent:             1,
		IsBatteryCharging:       1,
		IsBatteryDischarging:    0,
		UtilityVoltage:          121,
		OutputVoltage:           119,
		BatteryCapacityPercent:  1,
	}

	result := ParseToResult(sampleUps1500)
	assert.EqualValues(t, expected, result)
}
