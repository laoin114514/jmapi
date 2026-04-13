# jmapi (Go)

Go 模块名：`github.com/laoin114514/jmapi`

这是从 `JMComic-Crawler-Python` 迁移出的 Go 版本（库），提供 **API(APP)** 与 **HTML(网页)** 双客户端能力，以及一个可扩展的下载器与插件体系。

本 README 目标是：**详细到每个公开 API 都能直接复制示例运行**。

> 注意：这是一个用于学习/研究的非官方实现。目标站点可能随时更改页面或接口行为，导致部分能力失效。

## 安装

```bash
go get github.com/laoin114514/jmapi
```

## 快速开始

### 1) 创建客户端（API 模式 / 默认）

```go
package main

import (
	"fmt"
	"github.com/laoin114514/jmapi"
)

func main() {
	client := jmapi.NewClient(jmapi.Config{
		ClientType: jmapi.ClientTypeAPI,
	})

	album, err := client.GetAlbumDetail("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(album.ID, album.Name)
}
```

### 2) 创建客户端（HTML 模式 / 网页解析）

```go
client := jmapi.NewClient(jmapi.Config{
	ClientType: jmapi.ClientTypeHTML,
	// Domains 留空会使用内置默认 HTML 域名
})
album, _ := client.GetAlbumDetail("123456")
fmt.Println(album.Name, album.CommentCount)
```

### 3) 使用 Downloader（带插件与并发下载图片）

```go
opt := jmapi.DefaultOption()
opt.ClientConfig.ClientType = jmapi.ClientTypeAPI
opt.Download.Threading.Photo = 4
opt.Download.Threading.Image = 16

d := jmapi.NewDownloader(opt)
album, err := d.DownloadAlbum("123456")
if err != nil {
	panic(err)
}
_ = album
if err := d.RaiseIfHasFailures(); err != nil {
	// 表示“部分失败”（例如某些图片失败），可按需处理
	fmt.Println(err)
}
```

---

## 所有公开接口与示例

## 0. 构造与配置

### `NewClient(cfg Config) *Client`

```go
client := jmapi.NewClient(jmapi.Config{
	ClientType: jmapi.ClientTypeAPI,
	Domains:    []string{"www.cdnaspa.vip"},
	RetryTimes: 3,
})
```

#### `Config` 字段说明（常用）

- **ClientType**：`api` 或 `html`（`jmapi.ClientTypeAPI` / `jmapi.ClientTypeHTML`）
- **Domains**：域名列表（留空将使用内置默认域名）
- **Timeout**：HTTP 超时（默认 25s）
- **RetryTimes**：重试次数（默认 2，语义为“额外重试次数”，实现中会尝试 \(RetryTimes+1\) 次）
- **Headers / Cookies**：自定义请求头、cookies（用于登录、访问受限内容等）
- **AutoUpdateHost**：仅 API 模式有效，启动时尝试从“域名服务器”更新 API 域名
- **AutoEnsureCookies**：仅 API 模式有效，启动时请求 `/setting` 以确保具备必要 cookies
- **UseFixedTimestamp**：API 模式 tokenparam 的 ts 使用固定值（减少某些环境下波动）

### `SetDomains(domains []string)` / `Domains() []string`

```go
client.SetDomains([]string{"www.cdnaspa.club", "www.cdnplaystation6.vip"})
fmt.Println(client.Domains())
```

### `UpdateCookies(cookies map[string]string)`

```go
client.UpdateCookies(map[string]string{
	"AVS": "your_avs_cookie",
})
```

### `Option`（Downloader 配置）与 YAML

Downloader 使用 `Option`，你可以直接用 `DefaultOption()`，也可以从 YAML 加载：

```go
opt, err := jmapi.LoadOption("option.yml")
if err != nil {
	panic(err)
}
d := jmapi.NewDownloader(opt)
_, _ = d.DownloadAlbum("123456")
```

---

## 1. 详情接口

### `GetAlbumDetail(albumID string) (*AlbumDetail, error)`

```go
album, err := client.GetAlbumDetail("123456")
if err != nil {
	panic(err)
}
fmt.Println(album.Name, album.Tags)
fmt.Println("comment_count:", album.CommentCount)
fmt.Println("episode_ids:", len(album.EpisodeIDs))
```

### `GetPhotoDetail(photoID string, fetchAlbum bool, fetchScrambleID bool) (*PhotoDetail, error)`

```go
photo, err := client.GetPhotoDetail("654321", true, true)
if err != nil {
	panic(err)
}
fmt.Println(photo.ID, photo.AlbumID, photo.ScrambleID)
fmt.Println("images:", len(photo.PageArr))
```

### `GetScrambleID(photoID string) (string, error)`

```go
sid, err := client.GetScrambleID("654321")
if err != nil {
	panic(err)
}
fmt.Println("scramble id:", sid)
```

### `CheckPhoto(photo *PhotoDetail) error`

当你的 `PhotoDetail` 可能缺少必要字段（例如 `AlbumID`、`PageArr`）时，可调用此方法自动补齐：

```go
photo := &jmapi.PhotoDetail{ID: "654321"}
if err := client.CheckPhoto(photo); err != nil {
	panic(err)
}
fmt.Println(photo.AlbumID, len(photo.PageArr))
```

---

## 2. 搜索接口

### `Search(searchQuery string, page int, mainTag int, orderBy, timeRange, category, subCategory string) (*SearchResult, error)`

```go
res, err := client.Search("MANA", 1, 0, "mr", "a", "0", "")
if err != nil {
	panic(err)
}
fmt.Println(res.Total, len(res.Items))
if len(res.Items) > 0 {
	fmt.Println(res.Items[0].ID, res.Items[0].Name)
}
```

#### mainTag 参数（简要）

- `0`：站内（Site）
- `1`：作品（Work）
- `2`：作者（Author）
- `3`：标签（Tag）
- `4`：角色（Actor）

#### orderBy / timeRange 常用值

- `orderBy`: `mr`(最新) / `mv`(观看) / `tf`(喜欢) / `md`(评论) 等，见 `constants.go`
- `timeRange`: `a`(全部) / `m`(月) / `w`(周) / `t`(日)，见 `constants.go`

### `SearchSite(searchQuery string, page int, orderBy, timeRange, category, subCategory string)`

```go
res, _ := client.SearchSite("人妻", 1, "mr", "a", "0", "")
fmt.Println(len(res.Items))
```

### `SearchWork(searchQuery string, page int, orderBy, timeRange, category, subCategory string)`

```go
res, _ := client.SearchWork("MANA", 1, "mr", "a", "0", "")
fmt.Println(res.Total)
```

### `SearchAuthor(searchQuery string, page int, orderBy, timeRange, category, subCategory string)`

```go
res, _ := client.SearchAuthor("作者名", 1, "mr", "a", "0", "")
fmt.Println(len(res.Items))
```

### `SearchTag(searchQuery string, page int, orderBy, timeRange, category, subCategory string)`

```go
res, _ := client.SearchTag("無修正", 1, "mr", "a", "0", "")
fmt.Println(len(res.Items))
```

### `SearchActor(searchQuery string, page int, orderBy, timeRange, category, subCategory string)`

```go
res, _ := client.SearchActor("角色名", 1, "mr", "a", "0", "")
fmt.Println(len(res.Items))
```

---

## 3. 分类/排行接口

### `CategoriesFilter(page int, timeRange, category, orderBy, subCategory string) (*SearchResult, error)`

```go
res, err := client.CategoriesFilter(1, "a", "0", "mv", "")
if err != nil {
	panic(err)
}
fmt.Println(len(res.Items))
```

### `MonthRanking(page int, category string)`

```go
res, _ := client.MonthRanking(1, "0")
fmt.Println(len(res.Items))
```

### `WeekRanking(page int, category string)`

```go
res, _ := client.WeekRanking(1, "0")
fmt.Println(len(res.Items))
```

### `DayRanking(page int, category string)`

```go
res, _ := client.DayRanking(1, "0")
fmt.Println(len(res.Items))
```

---

## 4. 用户接口

### `Login(username, password string) (map[string]any, error)`

```go
profile, err := client.Login("your_username", "your_password")
if err != nil {
	panic(err)
}
fmt.Println(profile["username"], profile["uid"])
```

> HTML 模式下 `FavoriteFolder` 需要额外提供 `username`（因为网页端路径是 `/user/{username}/favorite/albums`）。

### `FavoriteFolder(page int, orderBy, folderID, username string) (*FavoriteResult, error)`

```go
fav, err := client.FavoriteFolder(1, "mr", "0", "your_username")
if err != nil {
	panic(err)
}
fmt.Println(fav.Total, len(fav.Items))
```

### `AddFavoriteAlbum(albumID, folderID string) (map[string]any, error)`

```go
ret, err := client.AddFavoriteAlbum("123456", "0")
if err != nil {
	panic(err)
}
fmt.Println(ret)
```

### `AlbumComment(videoID, comment, originator, status, commentID string) (map[string]any, error)`

```go
ret, err := client.AlbumComment("123456", "测试评论", "", "true", "")
if err != nil {
	panic(err)
}
fmt.Println(ret)
```

回复评论示例：

```go
ret, err := client.AlbumComment("123456", "回复内容", "", "", "999999")
if err != nil {
	panic(err)
}
fmt.Println(ret)
```

---

## 5. 图片接口

### `DownloadImage(imgURL string) ([]byte, error)`

```go
b, err := client.DownloadImage("https://cdn-msp.jmapiproxy1.cc/media/photos/123456/00001.jpg")
if err != nil {
	panic(err)
}
fmt.Println("bytes:", len(b))
```

### `DownloadByImageDetail(photoID, imageName string) ([]byte, error)`

`imageName` 留空会默认下载该 photo 的第一张图：

```go
b, err := client.DownloadByImageDetail("654321", "")
if err != nil {
	panic(err)
}
fmt.Println("bytes:", len(b))
```

### `DownloadAlbumCover(albumID string) ([]byte, error)`

```go
b, err := client.DownloadAlbumCover("123456")
if err != nil {
	panic(err)
}
fmt.Println("cover bytes:", len(b))
```

---

## 6. Downloader（下载整本/章节）

### `NewDownloader(option Option) *Downloader`

```go
opt := jmapi.DefaultOption()
opt.ClientConfig.ClientType = jmapi.ClientTypeAPI
opt.DirRule.BaseDir = "./downloads"

d := jmapi.NewDownloader(opt)
album, err := d.DownloadAlbum("123456")
if err != nil {
	panic(err)
}
fmt.Println(album.Name)

if err := d.RaiseIfHasFailures(); err != nil {
	fmt.Println("partial failures:", err)
}
```

### `DownloadPhoto(photoID string) (*PhotoDetail, error)`

```go
opt := jmapi.DefaultOption()
d := jmapi.NewDownloader(opt)
photo, err := d.DownloadPhoto("654321")
if err != nil {
	panic(err)
}
fmt.Println(photo.Name)
```

---

## 7. 插件系统（可配置参数 + 策略控制）

### 插件配置入口（Option.YAML）

```yaml
plugins:
  valid: log
  after_init:
    - plugin: retry_tuning
      safe: true
      log: true
      valid: log
      kwargs:
        retry_times: 6

  before_image:
    - plugin: image_suffix_filter
      safe: true
      kwargs:
        suffixes: [".jpg", ".png"]

  after_image:
    - plugin: topic_filter
      safe: true
      kwargs:
        allow: ["album", "photo", "image"]
```

### 内置插件（首批）

1. `topic_filter`：日志主题过滤（album/photo/image）
2. `image_suffix_filter`：图片后缀过滤
3. `retry_tuning`：启动时调整 `RetryTimes`

---

## 8. API / HTML 两种模式差异（实用提示）

- **API 模式（ClientTypeAPI）**
  - 更适合：详情/搜索/排行/登录/收藏/评论/图片下载的“接口化”路径
  - 支持：`AutoUpdateHost`、`AutoEnsureCookies`
- **HTML 模式（ClientTypeHTML）**
  - 更适合：当 API 域名不稳定或 API 行为变化时的备用方案
  - 注意：网页端某些功能（如收藏列表）可能需要你提供 `username`，且更依赖 cookies/headers

---

## 与 Python 版本关系说明

- 本 Go 版本聚焦“API 提供器”能力，不包含命令行。
- 已覆盖核心：详情、搜索、分类/排行、登录/收藏/评论、图片下载、Downloader、插件体系、API/HTML 双客户端门面。
- Python 版本的一些更高级工程化能力（更完整的 HTML 兼容、更多域名策略、更多插件、CLI 生态等）仍在逐步迁移中。
