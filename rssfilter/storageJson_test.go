package rssfilter

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		rss1 *[]RSSEntry
		rss2 *[]RSSEntry
	}
	item1 := RSSEntry{
		Link: "https://url1",
	}
	item2 := RSSEntry{
		Link: "https://url2",
	}
	item1New := RSSEntry{
		Link:       "https://url1",
		Reputation: "Good",
	}
	tests := []struct {
		name string
		args args
		want *[]RSSEntry
	}{
		{
			name: "どちらにもエントリが存在し、重複するものがない",
			args: args{
				rss1: &[]RSSEntry{item1},
				rss2: &[]RSSEntry{item2},
			},
			want: &[]RSSEntry{item1, item2},
		},
		{
			name: "rss1のみnil",
			args: args{
				rss1: nil,
				rss2: &[]RSSEntry{item2},
			},
			want: &[]RSSEntry{item2},
		},
		{
			name: "rss2のみnil",
			args: args{
				rss1: &[]RSSEntry{item1},
				rss2: nil,
			},
			want: &[]RSSEntry{item1},
		},
		{
			name: "重複するものがある",
			args: args{
				rss1: &[]RSSEntry{item1},
				rss2: &[]RSSEntry{item1, item2},
			},
			want: &[]RSSEntry{item1, item2},
		},
		{
			name: "キー以外の変更があれば、上書きされる",
			args: args{
				rss1: &[]RSSEntry{item1},
				rss2: &[]RSSEntry{item1New, item2},
			},
			want: &[]RSSEntry{item1New, item2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.rss1, tt.args.rss2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
