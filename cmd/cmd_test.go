/*
Copyright © 2023 Yonatan Sasson <yonatanxd72@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"testing"
)

func TestAddinsCmd(t *testing.T) {
	// Create a command buffer to capture the output
	var cmdOutput bytes.Buffer

	// Replace stdout with our buffer
	rootCmd.SetOutput(&cmdOutput)

	// Create a Cobra command with the desired arguments
	rootCmd.SetArgs([]string{"addins", "example.com", "--daemon-file", "../example.json", "--port", "8080"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Command execution failed: %v", err)
	}

	// Check if the output matches the expected output
	expectedOutput := ""
	if cmdOutput.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s\n", expectedOutput, cmdOutput.String())
	}
}

func TestRminsCmd(t *testing.T) {
	// Create a command buffer to capture the output
	var cmdOutput bytes.Buffer

	// Replace stdout with our buffer
	rootCmd.SetOutput(&cmdOutput)

	// Create a Cobra command with the desired arguments
	rootCmd.SetArgs([]string{"rmins", "example.com", "--daemon-file", "../example.json", "--port", "8080"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Command execution failed: %v", err)
	}

	// Check if the output matches the expected output
	expectedOutput := ""
	if cmdOutput.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s\n", expectedOutput, cmdOutput.String())
	}
}
