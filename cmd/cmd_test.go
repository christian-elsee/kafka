package cmd

import (
  "testing"
)

func TestRootCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := rootCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }
}

func TestCreateCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := createCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }
}

func TestDeleteCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := deleteCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }
}

func TestListCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := listCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }  
}

func TestReadCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := readCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }  
}

func TestTopicCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := topicCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }  
}

func TestWriteCommand(t *testing.T) {
  t.Logf("Enter")
  defer t.Logf("Exit")
  
  err := writeCmd.Execute()
  if err != nil {
    t.Errorf("Failed to execute command :: %s", err)
  }  
}
