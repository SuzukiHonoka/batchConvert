package converter

import "strconv"

// DefaultOptions aims to provide the high quality h265 encoding
var DefaultOptions = []Option{
	WithRateControl(RateControlCQP),
	WithQuantizationParameter(DefaultQP),
	WithAppleCompatibility(),
	WithHideBanner(),
	WithLogLevel("warning"),
	WithProgress("-"),
}

type Option func(*Config)

func WithQuantizationParameter(qp QuantizationParameter) Option {
	return func(config *Config) {
		config.extraFlags["qp_i"] = strconv.Itoa(qp.KeyFrame)
		config.extraFlags["qp_p"] = strconv.Itoa(qp.PredictiveFrame)
	}
}

func WithRateControl(rc RateControl) Option {
	return func(config *Config) {
		config.extraFlags["rc"] = string(rc)
	}
}

// WithAppleCompatibility refers to https://trac.ffmpeg.org/wiki/Encode/H.265#FinalCutandApplestuffcompatibility
func WithAppleCompatibility() Option {
	return func(config *Config) {
		config.extraFlags["tag:v"] = "hvc1"
	}
}

func WithHideBanner() Option {
	return func(config *Config) {
		config.extraFlags["hide_banner"] = ""
	}
}

func WithLogLevel(level string) Option {
	return func(config *Config) {
		config.extraFlags["loglevel"] = level
	}
}

func WithProgress(url string) Option {
	return func(config *Config) {
		config.extraFlags["progress"] = url
	}
}

func WithStripMetadata() Option {
	return func(config *Config) {
		config.extraFlags["map_metadata"] = "-1"
	}
}

func WithRawFlags(raw string) Option {
	return func(config *Config) {
		config.rawFlags = raw
	}
}
