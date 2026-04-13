# jmapi examples

## 准备

1. 进入目录：

```bash
cd goRebuild
```

2. 可选：设置环境变量（含登录接口时建议）

- `JM_USERNAME`
- `JM_PASSWORD`

## 运行方式

每个示例都是独立 `main` 包，使用 `go run` 运行。

```bash
go run ./example/album_detail -id 123456
```

## 示例列表

- `example/album_detail`：获取本子详情
- `example/photo_detail`：获取章节详情
- `example/search`：搜索本子
- `example/ranking`：排行榜（月/周/日）
- `example/favorite`：登录后读取收藏夹
- `example/comment`：登录后发表评论
- `example/image`：下载封面图
- `example/option_usage`：Option 配置与目录规则
- `example/downloader_basic`：Downloader 下载一本
- `example/plugin_simple`：插件回调示例
