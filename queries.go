package ltxref

import (
	"sort"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

// Return all commands with a given tag in no special order
func (l *Ltxref) GetCommandsWithTag(tagname string) []Command {
	var commandsWithTag []Command
	for _, command := range l.Commands {
		if hasTag(command, tagname) {
			commandsWithTag = append(commandsWithTag, command)
		}
	}
	return commandsWithTag
}

// Returns all tags in alphabetical order.
func (l *Ltxref) Tags() []string {
	// Needs better implementation!
	tags := make(map[string]bool)
	for _, command := range l.Commands {
		for _, label := range command.Label {
			tags[label] = true
		}
	}
	mk := make([]string, len(tags))
	i := 0
	for k, _ := range tags {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	return mk
}

func (l *Ltxref) FilterCommands(like string) []Command {
	if like == "" {
		return l.Commands
	}
	var commandsThatMatch []Command
	for _, command := range l.Commands {
		if fuzzy.Match(like, command.Name) {
			commandsThatMatch = append(commandsThatMatch, command)
		}
	}

	return commandsThatMatch
}

// Return true if command c has the given label (tag)
func hasTag(c Command, label string) bool {
	for _, v := range c.Label {
		if v == label {
			return true
		}
	}
	return false
}
