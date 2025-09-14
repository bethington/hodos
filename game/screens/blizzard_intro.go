package screens

import (
	"nostos/common/data/video"
	d2asset "nostos/core/asset"
	d2screen "nostos/core/screens"
)

// BlizzardIntro represents the Blizzard Intro screen
type BlizzardIntro struct {
	asset        *d2asset.AssetManager
	videoDecoder *video.BinkDecoder
}

// CreateBlizzardIntro creates a Blizzard Intro screen
func CreateBlizzardIntro(asset *d2asset.AssetManager) *BlizzardIntro {
	return &BlizzardIntro{
		asset: asset,
	}
}

// OnLoad loads the resources for the Blizzard Intro screen
func (v *BlizzardIntro) OnLoad(loading d2screen.LoadingState) {
	videoBytes, err := v.asset.LoadFile("/data/local/video/BlizNorth640x480.bik")
	if err != nil {
		loading.Error(err)
		return
	}

	loading.Progress(fiftyPercent)

	v.videoDecoder, err = video.CreateBinkDecoder(videoBytes)
	if err != nil {
		loading.Error(err)
	}
}
