package cpuanalyzer

const (
	defaultSegmentSize = 40
)

type Config struct {
	// SamplingInterval is the sampling interval for the same url.
	// The unit is seconds.
	SamplingInterval int `mapstructure:"sampling_interval"`
	// OpenJavaTraceSampling a switch for whether to use Java-Trace to trigger sampling.
	// The default is false.
	OpenJavaTraceSampling bool `mapstructure:"open_java_trace_sampling"`
	//JavaTraceSlowTime is used to identify the threshold of slow requests recognized by the apm side
	// The unit is seconds.
	JavaTraceSlowTime int `mapstructure:"java_trace_slow_time"`
	// SegmentSize defines how many segments(seconds) can be cached to wait for sending.
	// The elder segments will be overwritten by the newer ones, so don't set it too low.
	SegmentSize int `mapstructure:"segment_size"`
	// EdgeEventsWindowSize is the size of the duration window that seats the edge events.
	// The unit is seconds. The greater it is, the more data will be stored.
	EdgeEventsWindowSize int `mapstructure:"edge_events_window_size"`
}

func NewDefaultConfig() *Config {
	return &Config{
		SamplingInterval:      5,
		OpenJavaTraceSampling: false,
		JavaTraceSlowTime:     500,
		SegmentSize:           40,
		EdgeEventsWindowSize:  2,
	}
}
