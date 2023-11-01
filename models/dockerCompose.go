package models

// DockerCompose is a main object in Docker-Compose file
type DockerCompose struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
}

// Web is a service in Docker-Compose file
type Web struct {
	Build   string   `yaml:"build"`
	Ports   []string `yaml:"ports"`
	Volumes []string `yaml:"volumes"`
	Links   []string `yaml:"links"`
}

// Database is a service in Docker-Compose file
type Database struct {
	Image       string   `yaml:"image"`
	Environment []string `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
}

// Services is a services run in Docker
type Services struct {
	Web      Web      `yaml:"web"`
	Database Database `yaml:"database"`
}
