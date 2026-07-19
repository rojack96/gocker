package gocker

import "testing"

func TestComposeUpCommand(t *testing.T) {
	command := Compose().FileName("compose.yml").Up().Build().ServiceNames("hello-world").GetCommand()
	expected := "docker compose --file compose.yml up --build hello-world"

	if command != expected {
		t.Fatalf("expected %q, got %q", expected, command)
	}
}
