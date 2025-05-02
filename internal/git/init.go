package git

import (
	"fmt"
	"os"
)

func Init() (isCreated bool) {
	// check if .git directory exists
	_, err := os.Stat(".git")
	if err == nil {
		fmt.Println("Error: .git directory already exists")
		return false
	}

	os.Mkdir(".git", 0755)

	// info
	infoText := []byte("# git ls-files --others --exclude-from=.git/info/exclude\n# Lines that start with '#' are comments.\n# For a project mostly in C, the following would be a good set of\n# exclude patterns (uncomment them if you want to use them):\n# *.[oa]\n# *~")
	os.Mkdir(".git/info", 0755)
	os.WriteFile(".git/info/exclude", infoText, 0755)

	// objects
	os.Mkdir(".git/objects", 0755)
	os.Mkdir(".git/objects/info", 0755)
	os.Mkdir(".git/objects/pack", 0755)

	// refs
	os.Mkdir(".git/refs", 0755)
	os.Mkdir(".git/refs/heads", 0755)
	os.Mkdir(".git/refs/tags", 0755)

	// HEAD
	os.WriteFile(".git/HEAD", []byte("ref: refs/heads/main"), 0755)

	// config
	configText := []byte("[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n\tbare = false\n\tlogallrefupdates = true\n\tignorecase = true\n\tprecomposeunicode = true")
	os.WriteFile(".git/config", configText, 0755)

	return true
}
