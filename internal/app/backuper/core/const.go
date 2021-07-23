package core

type Config struct {
	SPACES_KEY      string `json:"SPACES_KEY"`
	SPACES_SECRET   string `json:"SPACES_SECRET"`
	SPACES_ENDPOINT string `json:"SPACES_ENDPOINT"`
	SPACES_BUCKET   string `json:"SPACES_BUCKET"`
}

var DataDir = ""
