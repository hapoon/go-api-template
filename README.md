# go-api-template

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hapoon/go-api-template)
![CircleCI](https://img.shields.io/circleci/build/github/hapoon/go-api-template/main)
![GitHub](https://img.shields.io/github/license/hapoon/go-api-template)

API server template created by Go.

# Guide

## Installation

1. [Use this template](https://github.com/hapoon/go-api-template/generate)よりリポジトリを作成する。

2. 必要なモジュールをダウンロード

```
% go mod download
```

3. [mockgen](https://github.com/golang/mock)をインストール

```
% go install github.com/golang/mock/mockgen@v1.5.0
```

## Run locally

```
% make run
```

## Run unit test with coverage

```
% make test
```

## Build docker image

```
% make docker-build
```

## Run docker image

```
% make docker-run
```

## Generate mock

```
% make mockgen
```

# License

[MIT](LICENSE)