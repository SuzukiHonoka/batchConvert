package converter

// QuantizationParameter for I-frame and P-frame, from -1 to 51 (hevc_amf)
// As a general guideline, QP values in the range of 18 to 30 for typical video encoding tasks
type QuantizationParameter struct {
	KeyFrame        int
	PredictiveFrame int
}

// DefaultQP is best for video quality
// average bitrate about ~1.8M in 4K 60fps video encoding (hevc_amf)
var DefaultQP = QuantizationParameter{
	KeyFrame:        28,
	PredictiveFrame: 28,
}
