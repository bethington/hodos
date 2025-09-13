package asset

import (
	"nostos/common/cache"
	"nostos/common/d2loader"
	"nostos/common/d2util"
	d2tbl "nostos/common/fileformats/tbl"
	"nostos/core/d2records"
)

// NewAssetManager creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func NewAssetManager(logLevel d2util.LogLevel) (*AssetManager, error) {
	loader, err := d2loader.NewLoader(logLevel)
	if err != nil {
		return nil, err
	}

	records, err := d2records.NewRecordManager(logLevel)
	if err != nil {
		return nil, err
	}

	logger := d2util.NewLogger()
	logger.SetPrefix(logPrefix)
	logger.SetLevel(logLevel)

	manager := &AssetManager{
		Logger:     logger,
		Loader:     loader,
		tables:     make([]d2tbl.TextDictionary, 0),
		animations: cache.CreateCache(animationBudget),
		fonts:      cache.CreateCache(fontBudget),
		palettes:   cache.CreateCache(paletteBudget),
		transforms: cache.CreateCache(paletteTransformBudget),
		dt1s:       cache.CreateCache(dt1Budget),
		ds1s:       cache.CreateCache(ds1Budget),
		cofs:       cache.CreateCache(cofBudget),
		dccs:       cache.CreateCache(dccBudget),
		Records:    records,
	}

	return manager, err
}
