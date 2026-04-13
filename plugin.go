package jmapi

import "log"

type PluginContext struct {
	Option     *Option
	Downloader *Downloader
	Client     *Client
}

type Plugin interface {
	Key() string
	AfterInit(ctx PluginContext) error
	BeforeAlbum(ctx PluginContext, album *AlbumDetail) error
	AfterAlbum(ctx PluginContext, album *AlbumDetail) error
	BeforePhoto(ctx PluginContext, photo *PhotoDetail) error
	AfterPhoto(ctx PluginContext, photo *PhotoDetail) error
	BeforeImage(ctx PluginContext, photo *PhotoDetail, imageURL string, savePath string) error
	AfterImage(ctx PluginContext, photo *PhotoDetail, imageURL string, savePath string) error
}

type PluginAdapter struct{}

func (PluginAdapter) Key() string { return "adapter" }
func (PluginAdapter) AfterInit(ctx PluginContext) error { return nil }
func (PluginAdapter) BeforeAlbum(ctx PluginContext, album *AlbumDetail) error { return nil }
func (PluginAdapter) AfterAlbum(ctx PluginContext, album *AlbumDetail) error { return nil }
func (PluginAdapter) BeforePhoto(ctx PluginContext, photo *PhotoDetail) error { return nil }
func (PluginAdapter) AfterPhoto(ctx PluginContext, photo *PhotoDetail) error { return nil }
func (PluginAdapter) BeforeImage(ctx PluginContext, photo *PhotoDetail, imageURL string, savePath string) error { return nil }
func (PluginAdapter) AfterImage(ctx PluginContext, photo *PhotoDetail, imageURL string, savePath string) error { return nil }

type PluginManager struct {
	plugins []Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{}
}

func (pm *PluginManager) Register(p Plugin) {
	if p == nil {
		return
	}
	pm.plugins = append(pm.plugins, p)
}

func (pm *PluginManager) run(fn func(Plugin) error) error {
	for _, p := range pm.plugins {
		if err := fn(p); err != nil {
			return err
		}
	}
	return nil
}

func (pm *PluginManager) SafeRun(fn func(Plugin) error) {
	for _, p := range pm.plugins {
		if err := fn(p); err != nil {
			log.Printf("plugin [%s] error: %v", p.Key(), err)
		}
	}
}

func (pm *PluginManager) AfterInit(ctx PluginContext) error {
	return pm.run(func(p Plugin) error { return p.AfterInit(ctx) })
}
func (pm *PluginManager) BeforeAlbum(ctx PluginContext, album *AlbumDetail) error {
	return pm.run(func(p Plugin) error { return p.BeforeAlbum(ctx, album) })
}
func (pm *PluginManager) AfterAlbum(ctx PluginContext, album *AlbumDetail) error {
	return pm.run(func(p Plugin) error { return p.AfterAlbum(ctx, album) })
}
func (pm *PluginManager) BeforePhoto(ctx PluginContext, photo *PhotoDetail) error {
	return pm.run(func(p Plugin) error { return p.BeforePhoto(ctx, photo) })
}
func (pm *PluginManager) AfterPhoto(ctx PluginContext, photo *PhotoDetail) error {
	return pm.run(func(p Plugin) error { return p.AfterPhoto(ctx, photo) })
}
func (pm *PluginManager) BeforeImage(ctx PluginContext, photo *PhotoDetail, imageURL, savePath string) error {
	return pm.run(func(p Plugin) error { return p.BeforeImage(ctx, photo, imageURL, savePath) })
}
func (pm *PluginManager) AfterImage(ctx PluginContext, photo *PhotoDetail, imageURL, savePath string) error {
	return pm.run(func(p Plugin) error { return p.AfterImage(ctx, photo, imageURL, savePath) })
}
