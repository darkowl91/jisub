package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

const (
	iniConfig = "/jisub/jisub.ini"
)

// > jisub --config "user.token RandomTokenValueStr"
// > jisub --config "jira.url "https://jira-api.com/jira/rest/api/2"
// > jisub --syb-tasks "JIRA-39106 QA:2 BE:3 FE:4"
func main() {
	flag.Func("config", "prop.key value", config)
	flag.Func("sub-tasks", "JIRA-1234 BE:2 FE:3 QA:4", subTasks)
	flag.Parse()
}

func config(arg string) error {

	items := strings.Split(arg, " ")
	// expect key value pair
	if len(items) < 2 {
		return fmt.Errorf("wrong number of arguments provided")
	}

	// load config, in case not loaded create empty
	cfg, err := ini.Load(iniConfig)
	if err != nil {
		cfg = ini.Empty()
	}

	// parse config section name, key
	sectionKey := strings.Split(items[0], ".")
	if len(sectionKey) < 2 {
		return fmt.Errorf("incorrect value format, expect: section.key")
	}

	section := cfg.Section(sectionKey[0])
	section.Key(sectionKey[1]).SetValue(items[1])

	err = cfg.SaveTo(iniConfig)
	if err != nil {
		return err
	}

	return nil
}

// string in format: EPMHRMS-39106 QA:2 BE:3 FE:4
func subTasks(arg string) error {

	jira, err := buildNewJiraFromConfig()
	if err != nil {
		return fmt.Errorf("missing configuration: %w", err)
	}

	items := strings.Split(arg, " ")

	// expecting at lest parent ticket and 1 subtask
	if len(items) < 2 {
		return fmt.Errorf("wrong number of arguments provided")
	}

	// search for parent ticket to creates subtasks for
	parent, err := jira.Issue(items[0])
	if err != nil {
		return err
	}

	// parse sub-tasks map, prefix:estimate
	spMap := make(map[string]float64)
	for i, v := range items {
		if i == 0 {
			continue
		}
		parseSubtaskToMap(v, spMap)
	}

	// post subtasks creation
	subTasks, err := jira.SubTasks(*parent, spMap)
	if err != nil {
		return err
	}

	// print result
	for _, issue := range subTasks.Issues {
		fmt.Println(issue.Key)
	}

	return nil
}

func parseSubtaskToMap(subTasks string, resultMap map[string]float64) error {

	tasksItem := strings.Split(subTasks, ":")

	if len(tasksItem) < 2 {
		return fmt.Errorf("incorrect value format, expect: PREFIX:NUM")
	}
	value, err := strconv.ParseFloat(tasksItem[1], 64)
	if err != nil {
		return err
	}
	resultMap[tasksItem[0]] = value

	return nil
}

func buildNewJiraFromConfig() (*Jira, error) {

	pwd, _ := os.Getwd()
	cfg, err := ini.Load(pwd + iniConfig)
	if err != nil {
		return nil, err
	}
	baseUrl := cfg.Section("jira").Key("url").Value()
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("missing jira.url value")
	}

	userToken := cfg.Section("user").Key("token").Value()
	if len(userToken) == 0 {
		return nil, fmt.Errorf("missing user.token value")
	}

	return NewJira(baseUrl, BearerAuth(userToken)), nil
}
