# Imgur CLI

Simple CLI tool generated using [Cobra](https://github.com/spf13/cobra) and interacts with [Imgur API](https://apidocs.imgur.com/).

## Installation

`go get github.com/davidPatters0n/imgurcli`

### Top level commands
```
Usage:
  imgur-cli [command]

Available Commands:
  account     Interacts with the Account API
  authorize   Authorize your account
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  image       Interacts with the Image API
```

### Account Commands

```
Interacts with the Account API

Usage:
  imgur-cli account [flags]
  imgur-cli account [command]

Available Commands:
  current     Return standard information for the currently authenticated account
  settings    Returns the account settings for the currently authenticated user
  update      Update account settings for the currently authenticated user
  verified    Checks to see if the currently authenticated user has verified their email address
  verify      Sends a verification email to the currently authenticated user
```

### Image Commands

```
Interacts with the Image API

Usage:
  imgur-cli image [flags]
  imgur-cli image [command]

Available Commands:
  delete      Delete an image given an image ID
  favourite   Favorite an image with the given ID.
  fetch       Supply an image ID and get information about about it.
  upload      Supply a path to your image and upload it to Imgur
```

## Usage

As a prerequiste using this package assumes you already have a imgur account. If you don't have one please register through imgur. Once you've registered you'll need to authenticate yourself with this library by running: `imgurcli authorize` follow the instructions presented. Once complete you can run the following commands documented above.
