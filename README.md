# 📋 jisub

![ci-build](https://github.com/darkowl91/jisub/actions/workflows/ci-branch.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/darkowl91/jisub)](https://goreportcard.com/report/github.com/darkowl91/jisub)
[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/darkowl91/sys-dia-log/blob/master/LICENSE)

> CLI tool to simplify jira tickets interaction

## Install

Download [latest release](https://github.com/darkowl91/jisub/releases/latest) version. Extract to user home dir.
Add jisub executable to PATH `export PATH=$PATH:~/jisub`

## Config

+ Configure jira instance API path:

```bash
    jisub --config "jira.url https://jira-api.com/jira/rest/api/2"
```

+ Obtain jira token from profile and add it to configuration

```bash
    jisub --config "user.token JIRA_TOKEN"
```

## Usage

+ Create required sub tasks with estimates for the parent ticket:

```bash
    jisub --syb-tasks "QA:2 BE:3 FE:4" --fields "storypoints:4 dealsize:2,3,4" JIRA-39106 
```

+ Shorten version:

```bash
    jisub -st "QA:2 BE:3 FE:4" -f "storypoints:4 dealsize:2,3,4" JIRA-39106 
```

### Support

[![bmc](https://www.buymeacoffee.com/assets/img/guidelines/download-assets-sm-1.svg)](https://www.buymeacoffee.com/darkowl91)
