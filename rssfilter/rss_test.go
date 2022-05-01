package rssfilter

import (
	"reflect"
	"testing"
)

func TestRSSEntry_GenerateModel(t *testing.T) {
	tests := []struct {
		name      string
		fields    RSSEntry
		className string
		words     *[]string
		wantErr   bool
	}{
		{
			name: "通常ケース",
			fields: RSSEntry{
				Title:       "走れメロス",
				Description: "メロスは激怒した。必ず、かの邪智暴虐の王を除かなければならぬと決意した。メロスには政治がわからぬ。メロスは、村の牧人である。笛を吹き、羊と遊んで暮して来た。けれども邪悪に対しては、人一倍に敏感であった。",
				Link:        "https://www.aozora.gr.jp/cards/000035/files/1567_14913.html",
				Categories:  []string{"青空文庫", "王様"},
				Reputation:  "Good",
			},
			className: "Good",
			words:     &[]string{"メロス", "激怒", "邪智", "暴虐", "王", "決意", "政治", "村", "牧人", "笛", "羊", "邪悪", "人一倍", "敏感", "www.aozora.gr.jp", "category:青空文庫", "category:王様"},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := RSSEntry{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Link:        tt.fields.Link,
				Published:   tt.fields.Published,
				Categories:  tt.fields.Categories,
				Retrieved:   tt.fields.Retrieved,
				Reputation:  tt.fields.Reputation,
			}
			got, got1, err := e.GenerateLearnData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.className {
				t.Errorf("GenerateModel() got = %v, className %v", got, tt.className)
			}
			if !reflect.DeepEqual(got1, tt.words) {
				t.Errorf("GenerateModel() got1 = %v, words %v", got1, tt.words)
			}
		})
	}
}
