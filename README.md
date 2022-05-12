# RssFilter

RSSFilter filters your favorite RSS feeds by [naive Bayes classifier](https://en.wikipedia.org/wiki/Naive_Bayes_classifier).

# Installation

Download from GitHub releases.

# Usage

As an example, we use the Hacker News RSS feed.
https://news.ycombinator.com/rss

## fetch

The `rssfilter fetch` command fetches RSS feeds and converts them to JSON format.

```shell
rssfilter fetch -u "https://news.ycombinator.com/rss" -f hacker_news.json
```

The following JSON will be output to `hacker_news.json`.

```json
{
  "title": "Hacker News",
  "entries": [
    {
      "title": "Very impressive and useful article",
      "description": "Very impressive and useful article",
      "link": "https://example.com/very-impressive-and-useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
    },
    {
      "title": "Very bored and useless article",
      "description": "bored and useless article",
      "link": "https://example.com/bored-and-useless-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
    }
    (snip)
  ]
}
```

Running `rssfilter fetch` at intervals fetches RSS from the specified URL and append it to JSON.

## train

Training data is required for classification. Before create model, you must create the training data yourself.
Open `hacker_news.json` with your favorite editor and add `reputation: "Good"` or `reputation: "Bad"` to each articles.

```json
{
  "title": "Hacker News",
  "entries": [
    {
      "title": "Very impressive and useful article",
      "description": "Very impressive and useful article",
      "link": "https://example.com/very-impressive-and-useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
      "reputation": "Good"
    },
    {
      "title": "Very bored and useless article",
      "description": "bored and useless article",
      "link": "https://example.com/bored-and-useless-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00",
      "reputation": "Bad"
    }
    (snip)
  ]
}
```

After adding reputation, `rssfilter train` creates a model from the reputation you entered.

```shell
rssfilter train -f hacker_news.json -m hacker_news.model
```

The `hacker_news.model` file has been created. You cannot open via text editor.

## classify

`rssfilter classify` automatically add a reputation for each newly added articles based on the reputation you have entered.

Before running `rssfilter classify`, add new article using `rssfilter fetch`.

```shell
rssfilter fetch -u "https://news.ycombinator.com/rss" -f hacker_news.json
```

Added one new entry.

```json
{
  "title": "Hacker News",
  "entries": [
    {
      "title": "Very impressive and useful article",
      "description": "Very impressive and useful article",
      "link": "https://example.com/very-impressive-and-useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
      "reputation": "Good"
    },
    {
      "title": "Very bored and useless article",
      "description": "bored and useless article",
      "link": "https://example.com/bored-and-useless-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00",
      "reputation": "Bad"
    },
    {
      "title": "useful article",
      "description": "useful article",
      "link": "https://example.com/useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
    }
    (snip)
  ]
}
```

After that, run following command.

```shell
rssfilter classify -f hacker_news.json -m hacker_news.model
```

Then reputation is added to new entry. 

```json
{
  "title": "Hacker News",
  "entries": [
    {
      "title": "Very impressive and useful article",
      "description": "Very impressive and useful article",
      "link": "https://example.com/very-impressive-and-useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00"
      "reputation": "Good"
    },
    {
      "title": "Very bored and useless article",
      "description": "bored and useless article",
      "link": "https://example.com/bored-and-useless-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00",
      "reputation": "Bad"
    },
    {
      "title": "useful article",
      "description": "useful article",
      "link": "https://example.com/useful-article",
      "published": "0001-01-01T00:00:00Z",
      "retrieved": "2022-05-08T22:15:53.185503+09:00",
      "reputation": "Good"
    }
    (snip)
  ]
}
```

## export

`rssfilter export` exports entries that has good reputation to RSS. 

```shell
rssfilter export -f hacker_news.json -r hacker_news.xml
```

## status

`rssfilter status` shows the status of model.

```shell
rssfilter status -m hacker_news.model
```

```json
{
  "learned_count": 276,
  "words_by_classes": [
    {
      "class_name": "Good",
      "word_scores": [
        {
          "word": "impressive",
          "score": 0.04413102820746133
        },
        {
          "word": "useful",
          "score": 0.034576888080072796
        },
      (snip.)
      ]
    },
    {
      "class_name": "Bad",
      "word_scores": [
        {
          "word": "bored",
          "score": 0.06645056726094004
        },
        {
          "word": "useless",
          "score": 0.05186385737439222
        },
        (snip.)
      ]
    }
  ]
}
```


# License

"rssfilter" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).

