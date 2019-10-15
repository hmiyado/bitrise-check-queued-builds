# bitrise-check-queued-builds : Check how many builds are queued on bitrise

## Overview

Go command checks how many builds are queued on bitrise.

## Install

```
go get github.com/hmiyado/bitrise-check-queued-builds
```

## Example

```
// require environment variable "BITRISE_TOKEN"
$ echo $BITRISE_TOKEN
// {AppSlug} is the application identifier for your app on bitrise
$ bitrise-check-queued-builds {AppSlug}
// command returns only number
0
```
