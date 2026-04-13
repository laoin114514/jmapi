# jmapi (Go)

Go 模块名：`github.com/laoin114514/jmapi`

这是从 `JMComic-Crawler-Python` 的 API 能力迁移出的 Go 版本（不含命令行）。

## 安装

```bash
go get github.com/laoin114514/jmapi
```

## 快速开始

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

---

## 所有公开接口与示例

### 构造与配置

### `NewClient(cfg Config) *Client`

```go
client := jmapi.NewClient(jmapi.Config{
	ClientType: jmapi.ClientTypeAPI,
	Domains:    []string{"www.cdnaspa.vip"},
	RetryTimes: 3,
})
```

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

---

### 详情接口

### `GetAlbumDetail(albumID string) (*AlbumDetail, error)`

```go
album, err := client.GetAlbumDetail("123456")
if err != nil {
	panic(err)
}
fmt.Println(album.Name, album.Tags)
```

### `GetPhotoDetail(photoID string, fetchAlbum bool, fetchScrambleID bool) (*PhotoDetail, error)`

```go
photo, err := client.GetPhotoDetail("654321", true, true)
if err != nil {
	panic(err)
}
fmt.Println(photo.ID, photo.AlbumID, photo.ScrambleID)
```

### `GetScrambleID(photoID string) (string, error)`

```go
sid, err := client.GetScrambleID("654321")
if err != nil {
	panic(err)
}
fmt.Println("scramble id:", sid)
```

---

### 搜索接口

### `Search(searchQuery string, page int, mainTag int, orderBy, timeRange, category, subCategory string) (*SearchResult, error)`

```go
res, err := client.Search("MANA", 1, 0, "mr", "a", "0", "")
if err != nil {
	panic(err)
}
fmt.Println(res.Total, len(res.Items))
```

### `SearchSite(searchQuery string, page int, orderBy, timeRange string)`

```go
res, _ := client.SearchSite("人妻", 1, "mr", "a")
fmt.Println(len(res.Items))
```

### `SearchAuthor(searchQuery string, page int, orderBy, timeRange string)`

```go
res, _ := client.SearchAuthor("作者名", 1, "mr", "a")
fmt.Println(len(res.Items))
```

### `SearchTag(searchQuery string, page int, orderBy, timeRange string)`

```go
res, _ := client.SearchTag("無修正", 1, "mr", "a")
fmt.Println(len(res.Items))
```

### `SearchActor(searchQuery string, page int, orderBy, timeRange string)`

```go
res, _ := client.SearchActor("角色名", 1, "mr", "a")
fmt.Println(len(res.Items))
```

---

### 分类/排行接口

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

### 用户接口

### `Login(username, password string) (map[string]any, error)`

```go
profile, err := client.Login("your_username", "your_password")
if err != nil {
	panic(err)
}
fmt.Println(profile["username"], profile["uid"])
```

### `FavoriteFolder(page int, orderBy, folderID string) (*FavoriteResult, error)`

```go
fav, err := client.FavoriteFolder(1, "mr", "0")
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

### 图片接口

### `DownloadImage(imgURL string) ([]byte, error)`

```go
b, err := client.DownloadImage("https://cdn-msp.jmapiproxy1.cc/media/photos/123456/00001.jpg")
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

## 与 Python 版本关系说明

- 本 Go 版本聚焦“API 提供器”能力，不包含命令行。
- 已覆盖核心：详情、搜索、分类/排行、登录/收藏/评论、图片下载。
- Python 项目中的高级特性（插件系统、复杂下载调度、HTML 解析端完整兼容、图片分片重组解码链路）暂未全部迁移。
