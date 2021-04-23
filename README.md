# go-api-template

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hapoon/go-api-template)
![CircleCI](https://img.shields.io/circleci/build/github/hapoon/go-api-template/main)
![GitHub](https://img.shields.io/github/license/hapoon/go-api-template)

API server template created by Go.

## Description

## Usage

### Run locally

```
% make run
```

### Run unit test with coverage

```
% make test
```

### Build docker image

```
% make docker-build
```

### Run docker image

```
% make docker-run
```

### Generate mock

```
% make mockgen
```


## Install

1. [Use this template](https://github.com/hapoon/go-api-template/generate)よりリポジトリを作成する。

2. 必要なモジュールをダウンロード

```
% go mod download
```

3. [mockgen](https://github.com/golang/mock)をインストール

```
% go install github.com/golang/mock/mockgen@v1.5.0
```

## Contribution

1. Fork it ( [https://github.com/hapoon/go-api-template/fork](https://github.com/hapoon/go-api-template/fork))
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Rebase your local changes against the main branch ( `git rebase -i` )
5. Push to the branch ( `git push origin my-new-feature` )
6. Create new Pull Request

## License

[MIT](LICENSE)