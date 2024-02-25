#!/bin/sh

set -eou pipefail

JSON_FILE="data/hot-it.json"
MODEL_FILE="data/hot-it.model"
RSS_FILE="public/hot-it.xml"

cd "$(dirname "$0")"

if [ ! -f "${JSON_FILE}" ]; then echo "error: ${JSON_FILE} not found"; exit 1; fi
if [ ! -f "${MODEL_FILE}" ]; then echo "error: ${MODEL_FILE} not found"; exit 1; fi
if [ -z "${SYNC_CLOUD_STORAGE}" ]; then echo "error: SYNC_CLOUD_STORAGE is not set"; exit 1; fi
if [ "${SYNC_CLOUD_STORAGE}" ]; then
  if [ -z "${RSS_URL}" ]; then echo "error: RSS_URL is not set"; exit 1; fi
  if [ -z "${RCLONE_DESTINATION}" ]; then echo "error: RCLONE_DESTINATION is not set"; exit 1; fi
  if [ -z "${RCLONE_CONFIG}" ]; then echo "error: RCLONE_CONFIG is not set"; exit 1; fi
fi

# fetch
./rssfilter fetch -u "https://b.hatena.ne.jp/hotentry/it.rss" -f "${JSON_FILE}"

# classify
./rssfilter classify -f "${JSON_FILE}" -m "${MODEL_FILE}"

if [ -n "${SYNC_CLOUD_STORAGE}" ]; then
  # setup rclone
  echo "$RCLONE_CONFIG" | base64 -d > rclone.conf
  # generate RSS
  ./rssfilter export -f "${JSON_FILE}" -r "${RSS_FILE}" -u "${RSS_URL}"

  # upload
  /usr/local/bin/rclone -v sync --config rclone.conf "${RSS_FILE}" "${RCLONE_DESTINATION}"
fi

