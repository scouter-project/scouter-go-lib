package packconstants

// Pack type
const (
	// M
	MAP                 byte = 10
	XLOG                byte = 21
	DROPPED_XLOG        byte = 22
	XLOG_PROFILE        byte = 26
	XLOG_PROFILE2       byte = 27
	TEXT                byte = 50
	METRIC              byte = 59
	PERFCOUNTER         byte = 60
	PERF_STATUS         byte = 61
	PERFCOUNTER_K8S     byte = 65
	K8S_CLUSTER__INFO   byte = 66
	ALERT               byte = 70
	ALERT2              byte = 71
	OBJECT              byte = 80
	CLOUD_WATCH         byte = 11
	XRay                byte = 12
	CLOUD_OBJECT        byte = 13
	CLOUD_ITEMS         byte = 14
	STACKDRIVER         byte = 15
	XRAY_HEATMAP        byte = 16
	AWS_LAMBDA          byte = 17
	AZUREMETRIC         byte = 30
	CLOUD_RESOURCE_DATA byte = 18
	CLOUD_METRIC_DATA   byte = 19
	K8S_CONTAINER       byte = 1
	K8S_NODE            byte = 2
	K8S_NAMESPACE       byte = 3
	K8S_WATCH_PACK      byte = 4
	K8S_CLUSTER_PACK    byte = 5

	K8S_EVENT_PACK byte = 6
)
