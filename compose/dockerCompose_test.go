package compose

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_OfDockerComposeYamlParser(t *testing.T) {
	path, err := filepath.Abs("./docker-compose.yml")
	if err != nil {
		fmt.Println("hogee")
		t.Errorf("docker-compose.yml not read; %s", err.Error())
	}
	containers, err := readYaml(path)
	if len(containers) != 2 {
		t.Errorf("data volume settings incorrect; 2 containers needed but %d found", len(containers))
	}

	volumeSettings := map[string][]string{"dbdata": []string{"/var/lib/mysql"}, "wpdata": []string{"/var/www/html"}}
	for _, container := range containers {
		for loadedVolume := range container.Volumes {
			for correctVolume := range volumeSettings[container.Name] {
				if loadedVolume != correctVolume {
					t.Errorf("'%s' container has incorrect volume '%s'", container.Name, loadedVolume)
				}
			}
		}
	}

	psCommands := makeComposePsCommandOptions("hoge", "/path/config.yml")
	correctCommands := []string{"--file=/path/config.yml", "ps", "-q", "hoge"}
	for i, psCommand := range psCommands {
		correctCommand := correctCommands[i]
		if psCommand != correctCommand {
			t.Error("makeComposePsCommand() makes invalid command: %v", psCommands)
		}
	}
}
