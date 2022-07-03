package Config

type SetConfig struct {
	SetIPFilter     bool
	SetWordsFilter  bool
	SetImagesFilter bool
}

var (
	Config SetConfig
)
