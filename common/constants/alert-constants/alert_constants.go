package alert_constants

const (
	SLOW_SQL                int16 = 1
	DELAYED_SERVICE         int16 = 2
	JAVA_METASPACE_USAGE    int16 = 3
	JAVA_TOO_MANY_MAP_ENTRY int16 = 4
	RESULTSET_LEAK          int16 = 5
	STMT_LEAK               int16 = 6
	CONNECTION_LEAK         int16 = 7
	JAVA_HEAP_USED          int16 = 8
	JAVA_HEAP_USED_PCT      int16 = 9
	FD_USED                 int16 = 10
	STUCK_SERVICE           int16 = 11
	MEMORY_USED_PCT         int16 = 12
	DISK_USED_PCT           int16 = 13
	CPU_USED_PCT            int16 = 14

	K8S_NODE_MEMORY_PRESSURE     int16 = 50
	K8S_NODE_DISK_PRESSURE       int16 = 51
	K8S_NODE_PID_PRESSURE        int16 = 52
	K8S_NODE_NETWORK_UNAVAILABLE int16 = 53
)
