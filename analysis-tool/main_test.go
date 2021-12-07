package main

import (
  "testing"
  "os"
)

// Test GetFileContent
func TestGetFileContent(t *testing.T) {
  // Change working directory
  os.Chdir("../")

  // Execute function
  content := GetFileContent("12-07-2021.txt")

  // Validate content
  if content == "error" {
    t.Error(content)
  }

  if len(content) == 0 {
    t.Error(content)
  }
}
