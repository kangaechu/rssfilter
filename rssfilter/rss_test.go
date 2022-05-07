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
			words:     &[]string{"メロス", "激怒", "邪智", "暴虐", "決意", "政治", "牧人", "邪悪", "人一倍", "敏感", "domain:www.aozora.gr.jp", "category:青空文庫", "category:王様"},
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

func TestRSSEntry_GenerateLearnData(t *testing.T) {
	tests := []struct {
		name      string
		fields    RSSEntry
		wantClass string
		wantWords *[]string
		wantErr   bool
	}{
		{
			name: "通常ケース",
			fields: RSSEntry{
				Title:       "2022年5月末に利用できなくなる「ID/パスワードのみのGoogleアカウントへのログイン」とは何か",
				Description: "これは何かという話です。 いきなりまとめ Googleへのログインの話 ではない カレンダーやメールなど古からのプロトコルでは パスワード認証を前提としたものがある それが OAuthに置き換わる 感じ PIM系プロトコルとパスワード認証(認可？) このネタどこかに書いた気がするな...ってので振り返ると、メールについてこの...",
				Link:        "https://zenn.dev/ritou/articles/170b2ef9acbe59",
				Categories:  []string{"テクノロジー", "OAuth", "google", "あとで読む", "セキュリティ", "security", "重要"},
				Reputation:  "Good",
			},
			wantClass: "Good",
			wantWords: &[]string{"月末", "利用", "ID", "パスワード", "Google", "アカウント", "ログイン", "これ", "カレンダー", "メール", "プロトコル", "認証", "前提", "もの", "それ", "OAuth", "感じ", "PIM", "認可", "？)", "ネタ", "どこ", "domain:zenn.dev", "category:テクノロジー", "category:OAuth", "category:google", "category:あとで読む", "category:セキュリティ", "category:security", "category:重要"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := RSSEntry{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Link:        tt.fields.Link,
				Categories:  tt.fields.Categories,
				Reputation:  tt.fields.Reputation,
			}
			gotClass, gotWords, err := e.GenerateLearnData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateLearnData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotClass != tt.wantClass {
				t.Errorf("GenerateLearnData() gotClass = %v, wantClass %v", gotClass, tt.wantClass)
			}
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("GenerateLearnData() gotWords = %v, wantClass %v", gotWords, tt.wantWords)
			}
		})
	}
}
