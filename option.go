package jmapi

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

type Option struct {
	ClientConfig Config
	DirRule      DirRule
	Download     DownloadOptions
	Plugins      PluginGroup
	UseCache     bool
}

type DownloadOptions struct {
	Image     ImageOptions
	Threading ThreadingOptions
}

type ImageOptions struct {
	Decode bool
	Suffix string
}

type ThreadingOptions struct {
	Image int
	Photo int
}

type DirRule struct {
	Rule        string
	BaseDir     string
	NormalizeZH string
}

type PluginGroup struct {
	BeforeAlbum []PluginConfig
	AfterAlbum  []PluginConfig
	BeforePhoto []PluginConfig
	AfterPhoto  []PluginConfig
	BeforeImage []PluginConfig
	AfterImage  []PluginConfig
	AfterInit   []PluginConfig
}

type PluginConfig struct {
	Plugin string
	Kwargs map[string]any
	Safe   bool
	Log    bool
	Valid  string
}

func DefaultOption() Option {
	wd, _ := os.Getwd()
	photoWorkers := runtime.NumCPU()
	return Option{
		ClientConfig: Config{
			ClientType:        ClientTypeAPI,
			RetryTimes:        2,
			AutoUpdateHost:    true,
			AutoEnsureCookies: true,
			UseFixedTimestamp: true,
		},
		DirRule: DirRule{
			Rule:    "Bd_Pname",
			BaseDir: wd,
		},
		Download: DownloadOptions{
			Image: ImageOptions{
				Decode: true,
				Suffix: "",
			},
			Threading: ThreadingOptions{
				Image: 30,
				Photo: photoWorkers,
			},
		},
		Plugins: PluginGroup{},
	}
}

func LoadOption(path string) (Option, error) {
	opt := DefaultOption()
	if path == "" {
		return opt, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return opt, err
	}
	if err := mergeOptionFromYAML(&opt, data); err != nil {
		return opt, err
	}
	return opt, nil
}

func mergeOptionFromYAML(opt *Option, data []byte) error {
	var raw map[string]any
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return err
	}
	if v, ok := raw["dir_rule"].(map[string]any); ok {
		if rule, ok := v["rule"].(string); ok {
			opt.DirRule.Rule = rule
		}
		if base, ok := v["base_dir"].(string); ok {
			opt.DirRule.BaseDir = base
		}
		if nzh, ok := v["normalize_zh"].(string); ok {
			opt.DirRule.NormalizeZH = nzh
		}
	}
	if v, ok := raw["download"].(map[string]any); ok {
		if cache, ok := v["cache"].(bool); ok {
			opt.UseCache = cache
		}
		if img, ok := v["image"].(map[string]any); ok {
			if dec, ok := img["decode"].(bool); ok {
				opt.Download.Image.Decode = dec
			}
			if suf, ok := img["suffix"].(string); ok {
				opt.Download.Image.Suffix = suf
			}
		}
		if th, ok := v["threading"].(map[string]any); ok {
			if img, ok := th["image"].(int); ok {
				opt.Download.Threading.Image = img
			}
			if photo, ok := th["photo"].(int); ok {
				opt.Download.Threading.Photo = photo
			}
		}
	}
	if v, ok := raw["client"].(map[string]any); ok {
		if impl, ok := v["impl"].(string); ok {
			opt.ClientConfig.ClientType = ClientType(impl)
		}
		if retry, ok := v["retry_times"].(int); ok {
			opt.ClientConfig.RetryTimes = retry
		}
		if domain, ok := v["domain"].([]any); ok {
			var domains []string
			for _, item := range domain {
				if s, ok := item.(string); ok {
					domains = append(domains, s)
				}
			}
			opt.ClientConfig.Domains = domains
		}
	}
	return nil
}

func (o Option) NewClient() *Client {
	return NewClient(o.ClientConfig)
}

func (o Option) DecideImageSuffix(original string) string {
	if o.Download.Image.Suffix != "" {
		return o.Download.Image.Suffix
	}
	return original
}

func (o Option) DecideImageFilename(index int) string {
	return fmt.Sprintf("%05d", index)
}

func (o Option) DecideImageSaveDir(album AlbumDetail, photo PhotoDetail) (string, error) {
	rule := o.DirRule.Rule
	if rule == "" {
		rule = "Bd_Pname"
	}
	base := o.DirRule.BaseDir
	if base == "" {
		base, _ = os.Getwd()
	}

	parts := splitRule(rule)
	paths := []string{}
	for _, r := range parts {
		switch r {
		case "Bd":
			paths = append(paths, base)
		case "Pname":
			paths = append(paths, sanitizePath(photo.Name))
		case "Pid":
			paths = append(paths, sanitizePath(photo.ID))
		case "Aname":
			paths = append(paths, sanitizePath(album.Name))
		case "Aid":
			paths = append(paths, sanitizePath(album.ID))
		default:
			paths = append(paths, r)
		}
	}

	out := filepath.Join(paths...)
	return out, nil
}

func splitRule(rule string) []string {
	if rule == "" {
		return []string{"Bd"}
	}
	if strings.Contains(rule, "/") {
		return strings.Split(rule, "/")
	}
	return strings.Split(rule, "_")
}

func sanitizePath(name string) string {
	name = strings.TrimSpace(name)
	replacer := strings.NewReplacer(
		"\\", "_",
		"/", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	name = replacer.Replace(name)
	if runtime.GOOS == "windows" {
		name = strings.TrimRight(name, ". ")
	}
	if name == "" {
		return "unknown"
	}
	return name
}
