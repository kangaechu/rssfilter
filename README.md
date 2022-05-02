# RssFilter

RSSFilterはお気に入りのRSSをナイーブベイズによりフィルタします。
これで読みたい記事だけフィルタされますね。


# Installation

Download from GitHub releases.

# Usage

Retrieve Hatena bookmark IT hot entries.

```bash
rssfilter fetch -u "https://b.hatena.ne.jp/hotentry/it.rss" -f hot-it.json
```

```shell
rssfilter classify -f hot-it.json -m hot-it.model
```

# License

"rssfilter" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).

