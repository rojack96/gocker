package command

import "testing"

func TestUpAttachDependencies(t *testing.T) {
	command := NewUp("docker compose up").AttachDependencies().ServiceNames().GetCommand()
	expected := "docker compose up --attach-dependencies"

	if command != expected {
		t.Fatalf("expected %q, got %q", expected, command)
	}
}

func TestUpScale(t *testing.T) {
	command := NewUp("docker compose up").Scale("web", 2).ServiceNames().GetCommand()
	expected := "docker compose up --scale web=2"

	if command != expected {
		t.Fatalf("expected %q, got %q", expected, command)
	}
}
