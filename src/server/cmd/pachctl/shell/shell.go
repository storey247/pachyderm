package shell

import (
	"log"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/pachyderm/pachyderm/src/server/pkg/uuid"
	"github.com/spf13/cobra"
)

const (
	completionAnnotation string = "completion"
	ldThreshold          int    = 2
)

var completions map[string]func(string) []prompt.Suggest = make(map[string]func(string) []prompt.Suggest)

func RegisterCompletionFunc(cmd *cobra.Command, completionFunc func(string) []prompt.Suggest) {
	id := uuid.NewWithoutDashes()

	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	} else if _, ok := cmd.Annotations[completionAnnotation]; ok {
		panic("duplicate completion func registration")
	}
	cmd.Annotations[completionAnnotation] = id
	completions[id] = completionFunc
}

type shell struct {
	rootCmd *cobra.Command
}

func newShell(rootCmd *cobra.Command) *shell {
	return &shell{rootCmd: rootCmd}
}

func (s *shell) executor(in string) {
	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader("pachctl " + in)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return
}

func (s *shell) suggestor(in prompt.Document) []prompt.Suggest {
	args := strings.Fields(in.Text)
	if len(strings.TrimSuffix(in.Text, " ")) < len(in.Text) {
		args = append(args, "")
	}
	cmd := s.rootCmd
	text := ""
	if len(args) > 0 {
		var err error
		cmd, _, err = s.rootCmd.Traverse(args[:len(args)-1])
		if err != nil {
			log.Fatal(err)
		}
		text = args[len(args)-1]
	}
	suggestions := cmd.SuggestionsFor(text)
	if len(suggestions) > 0 {
		var result []prompt.Suggest
		for _, suggestion := range suggestions {
			cmd, _, err := cmd.Traverse([]string{suggestion})
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, prompt.Suggest{
				Text:        suggestion,
				Description: cmd.Short,
			})
		}
		return result
	}
	if id, ok := cmd.Annotations[completionAnnotation]; ok {
		completionFunc := completions[id]
		suggests := completionFunc(text)
		var result []prompt.Suggest
		for _, s := range suggests {
			sText := s.Text
			if len(text) < len(sText) {
				sText = sText[:len(text)]
			}
			if ld(sText, text, true) < ldThreshold {
				result = append(result, s)
			}
		}
		return result
	}
	return nil
}

func (s *shell) run() {
	prompt.New(
		s.executor,
		s.suggestor,
		prompt.OptionPrefix(">>> "),
	).Run()
}

// Run runs a prompt, it does not return.
func Run(rootCmd *cobra.Command) {
	newShell(rootCmd).run()
}

// ld computes the Levenshtein Distance for two strings.
func ld(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}