package converter

import "maps"

// Presets for VideoCodec
type Presets map[VideoCodec]Flags

// DefaultPresets for VideoCodec
var DefaultPresets = Presets{
	VideoCodecHevcAmf: {
		"quality": "quality", // options: balanced, speed, quality
	},
}

// Get gets the preset of specific VideoCodec
func (p Presets) Get(vc VideoCodec) (Flags, bool) {
	preset, ok := DefaultPresets[vc]
	if !ok {
		return preset, false
	}
	// shallow copy prevents user changes the default preset
	return maps.Clone(preset), ok
}
