package pwrstat_parser

import (
	"strconv"
	"strings"
	"time"
)

const (
	debug           = false
	separator       = "="
	timeParseLayout = "2006/01/02 15:04:05"
)

type PwrstatResult struct {
	State                     float64
	ModelName                 string
	BatteryVoltage            float64
	InputRatingVoltage        float64
	OutputRatingWatts         float64
	IsAvrSupported            float64
	IsOnlineType              float64
	DiagnosticResult          float64
	DiagnosticDateUnix        float64
	PowerEventResult          float64
	PowerEventDateUnix        float64
	PowerEventDurationSeconds float64
	BatteryRemainingSeconds   float64
	IsBatteryCharging         float64
	IsBatteryDischarging      float64
	IsAcPresent               float64
	IsBoost                   float64
	UtilityVoltage            float64
	OutputVoltage             float64
	LoadPercent               float64
	BatteryCapacityPercent    float64

	// computed by OutputRatingWatts * LoadPercent
	LoadWatts float64
}

// Sample Data:
// state=0
// model_name= CP 1500C
// battery_volt=20600
// input_rating_volt=120000
// output_rating_watt=900000
// avr_supported=yes
// online_type=no
// diagnostic_result=1
// diagnostic_date=2022/11/02 19:12:21
// power_event_result=1
// power_event_date=2022/11/02 19:29:14
// power_event_during=4 sec.
// battery_remainingtime=3030
// battery_charging=no
// battery_discharging=no
// ac_present=yes
// boost=no
// utility_volt=121000
// output_volt=121000
// load=10000
// battery_capacity=100
func ParseToResult(input string) PwrstatResult {
	result := PwrstatResult{}

	// mapped properties
	splits := strings.Split(input, "\n")
	for _, line := range splits {
		if strings.Count(line, separator) != 1 {
			continue
		}

		cleanLine := strings.TrimSpace(line)
		lineSplit := strings.Split(cleanLine, separator)

		key := strings.TrimSpace(lineSplit[0])
		value := strings.TrimSpace(lineSplit[1])

		switch key {
		case "state":
			result.State = toFloat(value)
		case "model_name":
			result.ModelName = value
		case "battery_volt":
			result.BatteryVoltage = toFloat(value) / 1000
		case "input_rating_volt":
			result.InputRatingVoltage = toFloat(value) / 1000
		case "output_rating_watt":
			result.OutputRatingWatts = toFloat(value) / 1000
		case "avr_supported":
			result.IsAvrSupported = toBoolFloat(value)
		case "online_type":
			result.IsOnlineType = toBoolFloat(value)
		case "diagnostic_result":
			result.DiagnosticResult = toFloat(value)
		case "diagnostic_date":
			result.DiagnosticDateUnix = ParseDateStringToUnix(value)
		case "power_event_result":
			result.PowerEventResult = toFloat(value)
		case "power_event_date":
			result.PowerEventDateUnix = ParseDateStringToUnix(value)
		case "power_event_during":
			result.PowerEventDurationSeconds = toFloat(strings.Split(value, " ")[0])
		case "battery_remainingtime":
			result.BatteryRemainingSeconds = toFloat(value)
		case "battery_charging":
			result.IsBatteryCharging = toBoolFloat(value)
		case "battery_discharging":
			result.IsBatteryDischarging = toBoolFloat(value)
		case "ac_present":
			result.IsAcPresent = toBoolFloat(value)
		case "boost":
			result.IsBoost = toBoolFloat(value)
		case "utility_volt":
			result.UtilityVoltage = toFloat(value) / 1000
		case "output_volt":
			result.OutputVoltage = toFloat(value) / 1000
		case "load":
			result.LoadPercent = toFloat(value) / 100000
		case "battery_capacity":
			result.BatteryCapacityPercent = toFloat(value) / 100
		}
	}

	// computed properties
	result.LoadWatts = result.OutputRatingWatts * result.LoadPercent

	return result
}

func toBoolFloat(in string) float64 {
	if strings.EqualFold(in, "yes") {
		return 1
	} else {
		return 0
	}
}

func toFloat(in string) float64 {
	result, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0
	}

	return result
}

func ParseDateStringToUnix(in string) float64 {
	result, err := time.Parse(timeParseLayout, in)
	if err != nil {
		return 0
	}

	return float64(result.Unix())
}
