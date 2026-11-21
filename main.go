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
	if name := userName(); name != "" {
		return name
	}
	return "Human"
}

// userName returns the first word of the full name in the GECOS field of the
// current user, or else the login name that whoami prints. On Windows, it uses
// the display name, or else the account name without the domain.
func userName() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	login := u.Username
	if i := strings.LastIndex(login, `\`); i >= 0 {
		login = login[i+1:]
	}
	// By convention, "&" in the GECOS field stands for the capitalized login name.
	capitalized := login
	if capitalized != "" {
		capitalized = strings.ToUpper(capitalized[:1]) + capitalized[1:]
	}
	if fields := strings.Fields(strings.ReplaceAll(u.Name, "&", capitalized)); len(fields) > 0 {
		return fields[0]
	}
	return login
}

func main() {
	name := flag.String("name", "", "name at the start of the opener (default $MUTTER_NAME, then the git user.name, then the GECOS full name, then the login name)")
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
