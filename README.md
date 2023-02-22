# 📋 jisub

> CLI tool to simplify jira tickets interaction

## Install

If you are using Linux, macOS or WSL you can use the following command

```bash
curl -fsSL https://raw.github.com/darkowl91/jisub/main/tools/install.sh | sudo bash
```

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
