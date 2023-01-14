# Sample GET /metrics reply
```
# HELP pwrstat_battery_capacity_percent 
# TYPE pwrstat_battery_capacity_percent gauge
pwrstat_battery_capacity_percent 1
# HELP pwrstat_battery_remaining_seconds Remaing runtime in seconds
# TYPE pwrstat_battery_remaining_seconds gauge
pwrstat_battery_remaining_seconds 3257
# HELP pwrstat_battery_voltage 
# TYPE pwrstat_battery_voltage gauge
pwrstat_battery_voltage 20.6
# HELP pwrstat_diagnostic_date_unix 
# TYPE pwrstat_diagnostic_date_unix gauge
pwrstat_diagnostic_date_unix 1.667416341e+09
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
pwrstat_is_ac_present 1
# HELP pwrstat_is_avr_supported 
# TYPE pwrstat_is_avr_supported gauge
pwrstat_is_avr_supported 1
# HELP pwrstat_is_battery_charging 
# TYPE pwrstat_is_battery_charging gauge
pwrstat_is_battery_charging 0
# HELP pwrstat_is_battery_discharging 
# TYPE pwrstat_is_battery_discharging gauge
pwrstat_is_battery_discharging 0
# HELP pwrstat_is_boost 
# TYPE pwrstat_is_boost gauge
pwrstat_is_boost 0
# HELP pwrstat_is_online_type 
# TYPE pwrstat_is_online_type gauge
pwrstat_is_online_type 0
# HELP pwrstat_load_percent Percent load on the UPS
# TYPE pwrstat_load_percent gauge
pwrstat_load_percent 0.09
# HELP pwrstat_load_watts Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)
# TYPE pwrstat_load_watts gauge
pwrstat_load_watts 81
# HELP pwrstat_output_rating_watts Watt rating of the UPS
# TYPE pwrstat_output_rating_watts gauge
pwrstat_output_rating_watts 900
# HELP pwrstat_output_voltage 
# TYPE pwrstat_output_voltage gauge
pwrstat_output_voltage 120
# HELP pwrstat_power_event_date_unix 
# TYPE pwrstat_power_event_date_unix gauge
pwrstat_power_event_date_unix 1.673610628e+09
# HELP pwrstat_power_event_duration_seconds 
# TYPE pwrstat_power_event_duration_seconds gauge
pwrstat_power_event_duration_seconds 28
# HELP pwrstat_power_event_result 
# TYPE pwrstat_power_event_result gauge
pwrstat_power_event_result 1
# HELP pwrstat_state 
# TYPE pwrstat_state gauge
pwrstat_state 0
# HELP pwrstat_utility_voltage 
# TYPE pwrstat_utility_voltage gauge
pwrstat_utility_voltage 120
```