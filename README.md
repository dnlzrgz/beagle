# Beagle

[![Go Report Card](https://goreportcard.com/badge/github.com/danielkvist/beagle)](https://goreportcard.com/report/github.com/danielkvist/beagle)
[![GoDoc](https://godoc.org/github.com/danielkvist/beagle?status.svg)](https://godoc.org/github.com/danielkvist/beagle)
[![Docker Pulls](https://img.shields.io/docker/pulls/danielkvist/beagle.svg?maxAge=604800)](https://hub.docker.com/r/danielkvist/beagle/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

```text
            __
 \,--------/_/'--o  	Use beagle with
 /_    ___    /~"   	responsability.
  /_/_/  /_/_/
^^^^^^^^^^^^^^^^^^
```

beagle is simple Go cli to search for an especific username accross the Internet.

> beagle is a project inspired by [Sherlock](https://github.com/sherlock-project/sherlock).

## Example

```bash
beagle -g 4 -v -u mario
```

## Install

### Go

```bash
go get github.com/danielkvist/beagle
```

### Cloning the repository

```bash
# First, clone the repository
git clone https://github.com/danielkvist/beagle

# Then navigate into the beagle directory
cd beagle

# Run
go run main.go

# Or install
go install
```

### Docker

```bash
docker image pull danielkvist/beagle

# And
docker container run danielkvist/beagle -g 4 -v -u mario
```

> Note that the image danielkvist/beagle uses the urls.csv file from this repository. So it is not a valid option if you want to customize the URLs that beagle is gonna to use.

## Building the Docker image

```bash
# Inside the beagle directory, after cloning the git repository
docker image build -t beagle .
```

## Options

```text
$ beagle --help
Usage:
  beagle [flags]

Flags:
  -a, --agent string       user agent (default "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
      --csv string         .csv file with the URLs to parse and check (default "./urls.csv")
      --debug              prints error messages
      --disclaimer         disables disclaimer (default true)
  -g, --goroutines int     number of goroutines (default 1)
  -h, --help               help for beagle
  -p, --proxy string
  -t, --timeout duration   max time to wait for a response (default 3s)
  -u, --user string        username you want to search for (default "me")
  -v, --verbose            enables verbose mode
```

## URLs .csv file

The [urls.csv](https://github.com/danielkvist/beagle/blob/master/urls.csv) file of this repository contains more than 500 sites. Yet it's still possible that some sites may still be missing. If you have any suggestions please let me know by opening an issue.

The format of the ```.csv``` file, if you do not want to use the one provided by this repository must have the following structure:

```csv
name, url
```

The URL must contain a ```$``` where the username should go, for example:

```csv
instagram,https://instagram.com/$
devianart,https://$.devianart.com
```

## False positives

Some sites return an HTTP status code ```200 OK``` even if the user had not been found or doesn't exist. This causes beagle to report that the user has been found even though this has not been the case.

> In my tests these false positives are not too common to be a serious problem.

## Use beagle with responsability

beagle is a tool whose use I am not responsible for. And that has been built for the sole purpose of learning more about Go.

## Help is always welcome

If you have any problems or there is something you would like to improve please let me know.
