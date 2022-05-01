#!/usr/bin/env bash

# import
rssfilter import -u "https://b.hatena.ne.jp/hotentry/it.rss" -f hot-it.json

# classify
rssfilter classify -f hot-it.json -m hot-it.model

