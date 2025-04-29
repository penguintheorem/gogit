package cli

func PrintHelp() string {
	return `Usage: gogit [command] [options]
	
	Commands:
		help, --help show this help menu
		init, --init initialize a new git repository in the current directory
	`
}
