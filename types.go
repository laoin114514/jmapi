package jmapi

import "time"

type ClientType string

const (
	ClientTypeAPI  ClientType = "api"
	ClientTypeHTML ClientType = "html"
)

type Config struct {
	ClientType          ClientType
	Domains             []string
	Timeout             time.Duration
	RetryTimes          int
	AppVersion          string
	Proxies             map[string]string
	Cookies             map[string]string
	Headers             map[string]string
	AutoUpdateHost      bool
	AutoEnsureCookies   bool
	UseFixedTimestamp   bool
}

type AlbumDetail struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Author       []string       `json:"author,omitempty"`
	Description  string         `json:"description,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Works        []string       `json:"works,omitempty"`
	Actors       []string       `json:"actors,omitempty"`
	PageCount    int            `json:"page_count,omitempty"`
	CommentCount int            `json:"comment_count,omitempty"`
	Likes        string         `json:"likes,omitempty"`
	Views        string         `json:"views,omitempty"`
	EpisodeIDs   []string       `json:"episode_ids,omitempty"`
	Raw          map[string]any `json:"raw,omitempty"`
}

type PhotoDetail struct {
	ID          string         `json:"id"`
	AlbumID     string         `json:"album_id,omitempty"`
	Name        string         `json:"name"`
	SeriesID    string         `json:"series_id,omitempty"`
	Sort        string         `json:"sort,omitempty"`
	ScrambleID  string         `json:"scramble_id,omitempty"`
	PageArr     []string       `json:"page_arr,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
}

type SearchResult struct {
	Total int            `json:"total"`
	Items []AlbumListItem `json:"items"`
	Raw   map[string]any  `json:"raw,omitempty"`
}

type AlbumListItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Label     string `json:"label,omitempty"`
	Category  string `json:"category,omitempty"`
	TagList   []string `json:"tag_list,omitempty"`
}

type FavoriteResult struct {
	Total int             `json:"total"`
	Items []AlbumListItem `json:"items"`
	Raw   map[string]any  `json:"raw,omitempty"`
}

type APIResponse struct {
	Code     int             `json:"code"`
	ErrorMsg string          `json:"errorMsg"`
	Data     map[string]any  `json:"data"`
}
