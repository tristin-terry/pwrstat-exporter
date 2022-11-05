package pwrstat_parser

import (
	"strconv"
	"strings"
)

const (
	debug = false
)

type PwrstatResult struct {
	ModelName               string
	OutputRatingWatts       float64
	LoadPercent             float64
	LoadWatts               float64
	BatteryRemainingSeconds float64
	IsAcPresent             float64
	IsBatteryCharging       float64
	IsBatteryDischarging    float64
	UtilityVoltage          float64
	OutputVoltage           float64
	DiagnosticResult        float64
	BatteryCapacityPercent  float64
	BatteryVoltage          float64
	InputRatingVoltage      float64
}

func Parse(input string) map[string]string {
	measures := make(map[string]string)

	splits := strings.Split(input, "\n")
	// remove header element
	skipFirst := strings.EqualFold(splits[0], "STATUS")
	if skipFirst {
		splits = splits[1:]
	}

	// clean line and split into key/val
	for _, line := range splits {
		cleanLine := strings.TrimSpace(line)
		lineSplit := strings.Split(cleanLine, "=")

		if len(lineSplit) != 2 {
			continue
		}

		if strings.EqualFold(lineSplit[0], "") {
			continue
		}

		key := strings.TrimSpace(lineSplit[0])
		value := strings.TrimSpace(lineSplit[1])

		measures[key] = value
	}

	return measures
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
	mappedData := Parse(input)
	result := PwrstatResult{}

	result.ModelName = mappedData["model_name"]
	result.BatteryVoltage = toFloat(mappedData["battery_volt"]) / 1000
	result.InputRatingVoltage = toFloat(mappedData["input_rating_volt"]) / 1000
	result.UtilityVoltage = toFloat(mappedData["utility_volt"]) / 1000
	result.OutputVoltage = toFloat(mappedData["output_volt"]) / 1000
	result.OutputRatingWatts = toFloat(mappedData["output_rating_watt"]) / 1000
	result.DiagnosticResult = toFloat(mappedData["diagnostic_result"])
	result.BatteryRemainingSeconds = toFloat(mappedData["battery_remainingtime"])
	result.IsBatteryCharging = toBoolFloat(mappedData["battery_charging"])
	result.IsAcPresent = toBoolFloat(mappedData["ac_present"])
	result.IsBatteryDischarging = toBoolFloat(mappedData["battery_discharging"])
	result.LoadPercent = toFloat(mappedData["load"]) / 100000
	result.LoadWatts = result.OutputRatingWatts * result.LoadPercent
	result.BatteryCapacityPercent = toFloat(mappedData["battery_capacity"]) / 100

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
