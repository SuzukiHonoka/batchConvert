package converter

import (
	"fmt"
	"maps"
)

// Config of conversion
type Config struct {
	// VC for VideoCodec
	VC VideoCodec
	// AC for AudioCodec
	AC         AudioCodec
	extraFlags Flags
	rawFlags   string
}

func NewConfig(vc VideoCodec, ac AudioCodec, opts ...Option) *Config {
	config := &Config{
		VC:         vc,
		AC:         ac,
		extraFlags: make(Flags),
	}
	for _, opt := range opts {
		opt(config)
	}
	config.ApplyPreset()
	return config
}

// ApplyPreset applies preset if any
func (p *Config) ApplyPreset() {
	if preset, ok := DefaultPresets.Get(p.VC); ok {
		if p.extraFlags != nil {
			// allow user overrides the preset
			flags := maps.Clone(p.extraFlags)
			p.extraFlags = preset
			maps.Copy(p.extraFlags, flags)
			return
		}
		p.extraFlags = preset
	}
}

func (p *Config) Command(in, out string) string {
	if p.rawFlags == "" {
		return fmt.Sprintf("ffmpeg -y -i %s -c:v %s -c:a %s %s %s", in, p.VC, p.AC, p.extraFlags, out)
	} else {
		return fmt.Sprintf("ffmpeg -y -i %s -c:v %s -c:a %s %s %s %s", in, p.VC, p.AC, p.rawFlags, p.extraFlags, out)
	}
}
