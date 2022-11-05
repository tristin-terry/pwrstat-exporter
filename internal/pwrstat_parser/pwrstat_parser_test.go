package pwrstat_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sampleUps1500 = `STATUS
	state=1
	model_name= CP 1500C
	battery_volt=20600
	input_rating_volt=120000
	output_rating_watt=900000
	avr_supported=yes
	online_type=yes
	diagnostic_result=1
	diagnostic_date=2022/11/02 19:12:21
	power_event_result=1
	power_event_date=2022/11/02 19:29:14
	power_event_during=4 sec.
	battery_remainingtime=3030
	battery_charging=yes
	battery_discharging=yes
	ac_present=yes
	boost=yes
	utility_volt=121000
	output_volt=119000
	load=10000
	battery_capacity=100`
)

func TestParseToResult(t *testing.T) {
	expected := PwrstatResult{
		State:                     1,
		ModelName:                 "CP 1500C",
		BatteryVoltage:            20.6,
		InputRatingVoltage:        120,
		OutputRatingWatts:         900,
		IsAvrSupported:            1,
		IsOnlineType:              1,
		DiagnosticResult:          1,
		DiagnosticDateUnix:        1667416341,
		PowerEventResult:          1,
		PowerEventDateUnix:        1667417354,
		PowerEventDurationSeconds: 4,
		BatteryRemainingSeconds:   3030,
		IsBatteryCharging:         1,
		IsBatteryDischarging:      1,
		IsAcPresent:               1,
		IsBoost:                   1,
		UtilityVoltage:            121,
		OutputVoltage:             119,
		LoadPercent:               0.1,
		BatteryCapacityPercent:    1,
		LoadWatts:                 90,
	}

	actual := ParseToResult(sampleUps1500)

	assert.EqualValues(t, expected, actual)
}

func TestParseDateStringToUnix(t *testing.T) {
	actual := ParseDateStringToUnix("2022/11/02 19:12:21")

	assert.Equal(t, float64(1667416341), actual)
}
