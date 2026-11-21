// Command mutter prints a PR opener such as "Joe The Meatbag notates...".
// With "-" as the argument, it prints the opener above the text from stdin.
package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var nouns = []string{
	"Human", "Hooman", "Meatbag", "H2O-Sack", "Carbon Unit", "Primate",
	"Biped", "Mammal", "Flesh Golem", "Hairless Ape", "Homo Sapiens",
	"Mortal", "Earthling", "Bag of Bones", "Protein Engine", "Wetware",
	"Organism", "Sack of Water", "Air Breather", "Thumb Haver",
}

var verbs = []string{
	"writes", "notates", "eloquates", "mutters", "scribbles", "proclaims",
	"decrees", "pontificates", "types", "ponders", "bloviates", "opines",
	"declares", "scrawls", "utters", "murmurs", "expounds", "jots",
	"chronicles", "ruminates", "grumbles", "posits", "beseeches",
	"transcribes", "waxes poetic",
}

func pick(words []string) string {
	return words[rand.IntN(len(words))]
}

func defaultName() string {
	if name := strings.TrimSpace(os.Getenv("MUTTER_NAME")); name != "" {
		return name
	}
	out, err := exec.Command("git", "config", "user.name").Output()
	if err == nil {
		if fields := strings.Fields(string(out)); len(fields) > 0 {
			return fields[0]
		}
	}
	if name := gecosName(); name != "" {
		return name
	}
	return "Human"
}

// gecosName returns the first word of the full name in the GECOS field of the
// current user. On Windows, it uses the display name.
func gecosName() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	// By convention, "&" in the GECOS field stands for the capitalized login name.
	login := u.Username
	if login != "" {
		login = strings.ToUpper(login[:1]) + login[1:]
	}
	fields := strings.Fields(strings.ReplaceAll(u.Name, "&", login))
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

func main() {
	name := flag.String("name", "", "name at the start of the opener (default $MUTTER_NAME, then the git user.name, then the GECOS full name)")
	flag.Parse()
	if *name == "" {
		*name = defaultName()
	}

	opener := fmt.Sprintf("%s The %s %s...", *name, pick(nouns), pick(verbs))

	if flag.Arg(0) != "-" {
		fmt.Println(opener)
		return
	}

	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mutter:", err)
		os.Exit(1)
	}
	if len(body) == 0 {
		fmt.Println(opener)
		return
	}
	fmt.Printf("%s\n\n%s", opener, body)
}
