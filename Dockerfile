FROM busybox:1.36-uclibc as busybox

FROM rclone/rclone:1 as rclone

FROM golang:1.23-bookworm AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY rssfilter/ ./rssfilter
COPY cmd/ ./cmd

RUN go build -o /app/main -ldflags '-s -w' main.go


FROM gcr.io/distroless/base-debian12

COPY --from=busybox /bin/busybox /bin/busybox
RUN ["/bin/busybox", "--install", "/bin"]

USER nonroot:nonroot
WORKDIR /app
COPY --chown=nonroot:nonroot --from=build /app/main /app/rssfilter
COPY --chown=nonroot:nonroot --from=rclone /usr/local/bin/rclone /usr/local/bin/rclone
COPY --chown=nonroot:nonroot run.sh /app/run.sh

# ENV SYNC_CLOUD_STORAGE="true"
# ENV RSS_URL
# ENV RCLONE_DESTINATION
# ENV RCLONE_CONFIG

ENTRYPOINT ["/bin/sh", "-c", "/app/run.sh"]
