package ltxref

import (
	"sort"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

// Return all commands with a given tag in no special order
func (l *Ltxref) GetCommandsWithTag(tagname string) []Command {
	var commandsWithTag []Command
	for _, command := range l.Commands {
		if hasTag(command.Label, tagname) {
			commandsWithTag = append(commandsWithTag, command)
		}
	}
	return commandsWithTag
}

// packagename may be empty for the kernel commands
func (l *Ltxref) GetCommandFromPackage(commandname string, packagename string) *Command {
	var cmdlist []Command
	// Needs better implementation!

	if packagename != "" {
		for _, v := range l.Packages {
			if v.Name == packagename {
				cmdlist = v.Commands
				break
			}
		}
	} else {
		cmdlist = l.Commands
	}

	for _, v := range cmdlist {
		if v.Name == commandname {
			return &v
		}
	}
	return nil
}

// Return all environments with a given tag in no special order
func (l *Ltxref) GetEnvironmentsWithTag(tagname string) []Environment {
	var environmentsWithTag []Environment
	for _, env := range l.Environments {
		if hasTag(env.Label, tagname) {
			environmentsWithTag = append(environmentsWithTag, env)
		}
	}
	return environmentsWithTag
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

// Case insensitive fuzzy match.
func (l *Ltxref) FilterCommands(like string) []Command {
	if like == "" {
		return l.Commands
	} else {
		like = strings.ToLower(like)
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
func hasTag(labels []string, label string) bool {
	for _, v := range labels {
		if v == label {
			return true
		}
	}
	return false
}
