package compose

import (
	"bytes"
	"fmt"
	"github.com/tottokotkd/dockervlm/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
)

func GetComposeConfig(configPath string) ([]common.Container, error) {
	containers, err := readYaml(configPath)
	if err != nil {
		return containers, err
	}
	for _, container := range containers {
		container.Id, _ = getContainerId(container.Name, configPath)
	}

	return containers, nil
}

func readYaml(configPath string) ([]common.Container, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]map[string]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}

	containers := []common.Container{}
	for name, settings := range m {
		if volumes, ok := settings["volumes"].([]interface{}); ok {
			container := common.Container{Name: name}
			for _, volume := range volumes {
				if v, ok := volume.(string); ok {
					container.Volumes = append(container.Volumes, v)
				}
			}
			containers = append(containers, container)
		}
	}
	return containers, nil
}

func getContainerId(name string, configFile string) (string, error) {

	cmd := exec.Command("docker-compose", makeComposePsCommandOptions(name, configFile)...)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return stdout.String(), nil
}

func makeComposePsCommandOptions(name string, configFile string) []string {
	return []string{fmt.Sprintf("--file=%s", configFile), "ps", "-q", name}
}
