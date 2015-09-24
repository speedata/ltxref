package ltxref

import (
	"sort"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

// Return all commands with a given tag in no special order
func (l *Ltxref) GetCommandsWithTag(tagname string) []Command {
	var commandsWithTag []Command
	for _, command := range l.commands {
		if hasTag(command.Label, tagname) {
			commandsWithTag = append(commandsWithTag, command)
		}
	}
	return commandsWithTag
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

// Return all documentclasses with a given tag in no special order
func (l *Ltxref) GetDocumentclassesWithTag(tagname string) []Documentclass {
	var classesWithTag []Documentclass
	for _, env := range l.Documentclasses {
		if hasTag(env.Label, tagname) {
			classesWithTag = append(classesWithTag, env)
		}
	}
	return classesWithTag
}

// Return all packages with a given tag in no special order
func (l *Ltxref) GetPackagesWithTag(tagname string) []Package {
	var packagesWithTag []Package
	for _, env := range l.Packages {
		if hasTag(env.Label, tagname) {
			packagesWithTag = append(packagesWithTag, env)
		}
	}
	return packagesWithTag
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
		cmdlist = l.commands
	}

	for _, v := range cmdlist {
		if v.Name == commandname {
			return &v
		}
	}
	return nil
}

// Returns all tags in alphabetical order.
func (l *Ltxref) Tags() []string {
	// Needs better implementation!
	tags := make(map[string]bool)
	for _, command := range l.commands {
		for _, label := range command.Label {
			tags[label] = true
		}
	}
	for _, v := range l.Environments {
		for _, label := range v.Label {
			tags[label] = true
		}
	}

	for _, v := range l.Packages {
		for _, label := range v.Label {
			tags[label] = true
		}
		for _, cmd := range v.Commands {
			for _, label := range cmd.Label {
				tags[label] = true
			}

		}
	}
	for _, v := range l.Documentclasses {
		for _, label := range v.Label {
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
		return l.commands
	} else {
		like = strings.ToLower(like)
	}
	var commandsThatMatch []Command
	for _, command := range l.commands {
		if fuzzy.Match(like, command.Name) {
			commandsThatMatch = append(commandsThatMatch, command)
		}
	}
	return commandsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterEnvironments(like string) []Environment {
	if like == "" {
		return l.Environments
	} else {
		like = strings.ToLower(like)
	}
	var itemsThatMatch []Environment
	for _, item := range l.Environments {
		if fuzzy.Match(like, item.Name) {
			itemsThatMatch = append(itemsThatMatch, item)
		}
	}
	return itemsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterDocumentclasses(like string) []Documentclass {
	if like == "" {
		return l.Documentclasses
	} else {
		like = strings.ToLower(like)
	}
	var itemsThatMatch []Documentclass
	for _, item := range l.Documentclasses {
		if fuzzy.Match(like, item.Name) {
			itemsThatMatch = append(itemsThatMatch, item)
		}
	}
	return itemsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterPackages(like string) []Package {
	if like == "" {
		return l.Packages
	} else {
		like = strings.ToLower(like)
	}
	var itemsThatMatch []Package
	for _, item := range l.Packages {
		if fuzzy.Match(like, item.Name) {
			itemsThatMatch = append(itemsThatMatch, item)
		}
	}
	return itemsThatMatch
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
