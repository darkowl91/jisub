# 📋 jisub

> CLI tool to simplify jira tickets interaction

## Install

Download latest release version. Extract to user home dir.
Add to jisub executable to PATH `~/jisub`

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
    jisub --syb-tasks "JIRA-39106 QA:2 BE:3 FE:4"
```

Output:
> There are will be 3 sub-tasks created with the summary of parent task `JIRA-39106` eg:

+ `JIRA-1`
+ `JIRA-2`
+ `JIRA-3`

## TODO

+ Support multi command execution eg.:

``` bash
jisub --issue "JIRA-39106" --sub-tasks "QA:2 BE:3 FE:4" --fields "storypoints:4 dealsize:2,3,4"
```
