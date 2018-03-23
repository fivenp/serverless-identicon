# serverless-identicon
[![serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)
[![PyPI - License](https://img.shields.io/pypi/l/Django.svg)](LICENSE)


[go-identicon](https://github.com/fivenp/go-identicon) for AWS Lambda via Serverless Framework.

## Prerequisite

* [AWS Account](https://console.aws.amazon.com/)
* [Node.js](https://nodejs.org) v6.5.0 or later
* [Serverless](https://serverless.com) CLI v1.9.0 or later `npm install -g serverless`
* [dep](https://github.com/golang/dep)

A very good starting point is the [Serverless tutorial](https://serverless.com/blog/framework-example-golang-lambda-support/) for Golang on Lambda.

Make sure to also read the [first time deployment guide](https://serverless.com/blog/anatomy-of-a-serverless-app/#setup)!

## Build

```zsh
make
```

## Deploy

```zsh
sls deploy
```
*You can also pass `--aws-profile <PROFILE>` if you have multiple AWS accounts*

## Invoke

```zsh
sls invoke -f identicon
```
*You can also pass `--aws-profile <PROFILE>` if you have multiple AWS accounts*
