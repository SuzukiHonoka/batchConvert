package converter

type VideoCodec string

const (
	VideoCodecCopy    VideoCodec = "copy"
	VideoCodecHevcAmf VideoCodec = "hevc_amf"
)

type AudioCodec string

const (
	AudioCodecCopy AudioCodec = "copy"
)
