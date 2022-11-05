# Sample GET /metrics reply
```
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 2.1292e-05
go_gc_duration_seconds{quantile="0.25"} 2.6698e-05
go_gc_duration_seconds{quantile="0.5"} 3.1576e-05
go_gc_duration_seconds{quantile="0.75"} 4.8366e-05
go_gc_duration_seconds{quantile="1"} 0.000154472
go_gc_duration_seconds_sum 0.023321188
go_gc_duration_seconds_count 529
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 10
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.19.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 3.335088e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 8.71598e+08
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4262
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 2.283067e+06
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 9.010912e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 3.335088e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.823936e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 5.103616e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 2955
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 5.40672e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.1927552e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.667656462509554e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 2.286022e+06
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 4800
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 79968
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 97920
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 6.53068e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 815386
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 655360
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 655360
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 2.2526992e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 9
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 3.51
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 524288
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 11
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.601536e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.66762211018e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.267822592e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 2295
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP pwrstat_battery_capacity_percent 
# TYPE pwrstat_battery_capacity_percent gauge
pwrstat_battery_capacity_percent 1
# HELP pwrstat_battery_remaining_seconds Remaing runtime in seconds
# TYPE pwrstat_battery_remaining_seconds gauge
pwrstat_battery_remaining_seconds 2961
# HELP pwrstat_battery_voltage 
# TYPE pwrstat_battery_voltage gauge
pwrstat_battery_voltage 20.7
# HELP pwrstat_diagnostic_date_unix 
# TYPE pwrstat_diagnostic_date_unix gauge
pwrstat_diagnostic_date_unix 1.667416341e+09
# HELP pwrstat_diagnostic_result 
# TYPE pwrstat_diagnostic_result gauge
pwrstat_diagnostic_result 1
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
pwrstat_load_percent 0.08
# HELP pwrstat_load_watts Total calculated load watts (pwrstat_output_rating_watts * pwrstat_load_percent)
# TYPE pwrstat_load_watts gauge
pwrstat_load_watts 72
# HELP pwrstat_output_rating_watts Watt rating of the UPS
# TYPE pwrstat_output_rating_watts gauge
pwrstat_output_rating_watts 900
# HELP pwrstat_output_voltage 
# TYPE pwrstat_output_voltage gauge
pwrstat_output_voltage 121
# HELP pwrstat_power_event_date_unix 
# TYPE pwrstat_power_event_date_unix gauge
pwrstat_power_event_date_unix 1.667417354e+09
# HELP pwrstat_power_event_duration_seconds 
# TYPE pwrstat_power_event_duration_seconds gauge
pwrstat_power_event_duration_seconds 4
# HELP pwrstat_power_event_result 
# TYPE pwrstat_power_event_result gauge
pwrstat_power_event_result 1
# HELP pwrstat_state 
# TYPE pwrstat_state gauge
pwrstat_state 0
# HELP pwrstat_utility_voltage 
# TYPE pwrstat_utility_voltage gauge
pwrstat_utility_voltage 121
```