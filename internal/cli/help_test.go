package cli

import "testing"

func TestPrintHelp(t *testing.T) {
	expected := `Usage: gogit [command] [options]
	
	Commands:
		help, --help show this help menu
		init, --init initialize a new git repository in the current directory
	`
	got := PrintHelp()

	if got != expected {
		t.Errorf("PrintHelp() = %v, want %v", got, expected)
	}
}
