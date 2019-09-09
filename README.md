# Beagle

[![Go Report Card](https://goreportcard.com/badge/github.com/danielkvist/beagle)](https://goreportcard.com/report/github.com/danielkvist/beagle)
[![CircleCI](https://circleci.com/gh/danielkvist/beagle.svg?style=svg)](https://circleci.com/gh/danielkvist/beagle)
[![GoDoc](https://godoc.org/github.com/danielkvist/beagle?status.svg)](https://godoc.org/github.com/danielkvist/beagle)
[![Docker Pulls](https://img.shields.io/docker/pulls/danielkvist/beagle.svg?maxAge=604800)](https://hub.docker.com/r/danielkvist/beagle/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

```text
            __
 \,--------/_/'--o    Use beagle with
 /_    ___    /~"     responsibility.
  /_/_/  /_/_/
^^^^^^^^^^^^^^^^^^
```

Beagle is a CLI written in Go to search for an specific username across the Internet.

> Beagle is a project inspired by [Sherlock](https://github.com/sherlock-project/sherlock).

## Example

```bash
beagle -g 10 -t 1s -u me -v
```

## Install

### Go

```bash
go install github.com/danielkvist/beagle
```

### Docker

```bash
docker image pull danielkvist/beagle
```

> Note that the image danielkvist/beagle uses the urls.csv file from this repository. So it is not a valid option if you want to customize the URLs that beagle is gonna to use.

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

## Building the Docker image

```bash
# Inside the beagle directory, after cloning the git repository:
docker image build -t beagle .
```

## Options

```text
$ beagle --help
Beagle is a CLI written in Go to search for an specific username across the Internet.

Usage:
  beagle [flags]

Examples:
beagle -g 10 -t 1s -u me -v

Flags:
  -a, --agent string       user agent (default "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
      --debug              prints errors messages
  -f, --file string        .csv file with the URLs to check (default "./urls.csv")
  -g, --goroutines int     number of goroutines (default 1)
  -h, --help               help for beagle
  -p, --proxy string       proxy URL
  -t, --timeout duration   max time to wait for a response from a site (default 3s)
  -u, --user string        username you want to search for (default "me")
  -v, --verbose            prints all the results
```

## URLs .csv file

The [urls.csv](https://github.com/danielkvist/beagle/blob/master/urls.csv) file of this repository contains more than 500 sites. Yet it's still possible that some sites may still be missing. If you have any suggestions please let me know by opening an issue.

The format of the ```.csv``` file, if you do not want to use the one provided by this repository must have the following structure:

```csv
name, mainURL, userURL
```

The URLs must contain a ```$``` where the username should go, for example:

```csv
instagram,https://instagram.com/$,https://instagrma.com/$
devianart,https://$.devianart.com,https://$.devianart.com
```

## Use Beagle with responsability

Beagle is a tool whose use I am not responsible for. And that has been built for the sole purpose of learning more about Go.

## Help is always welcome

If you have any problems or there is something you would like to improve please let me know.
