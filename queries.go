package ltxref

import (
	"sort"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

// packagename may be empty for the kernel commands
func (l *Ltxref) GetCommandFromPackage(commandname string, packagename string) *Command {
	var cmdlist []Command
	// Needs better implementation!

	if packagename != "" {
		for _, v := range l.packages {
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
	for _, v := range l.environments {
		for _, label := range v.Label {
			tags[label] = true
		}
	}

	for _, v := range l.packages {
		for _, label := range v.Label {
			tags[label] = true
		}
		for _, cmd := range v.Commands {
			for _, label := range cmd.Label {
				tags[label] = true
			}

		}
	}
	for _, v := range l.documentclasses {
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
func (l *Ltxref) FilterCommands(like string, tag string) []Command {
	if like == "" && tag == "" {
		return l.commands
	} else {
		like = strings.ToLower(like)
		tag = strings.ToLower(tag)
	}
	var commandsThatMatch []Command
	for _, command := range l.commands {
		if fuzzy.Match(like, command.Name) && (tag == "" || hasTag(command.Label, tag)) {
			commandsThatMatch = append(commandsThatMatch, command)
		}
	}
	return commandsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterEnvironments(like string, tag string) []Environment {
	if like == "" && tag == "" {
		return l.environments
	} else {
		like = strings.ToLower(like)
		tag = strings.ToLower(tag)
	}
	var itemsThatMatch []Environment
	for _, item := range l.environments {
		if fuzzy.Match(like, item.Name) && (tag == "" || hasTag(item.Label, tag)) {
			itemsThatMatch = append(itemsThatMatch, item)
		}
	}
	return itemsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterDocumentclasses(like string, tag string) []Documentclass {
	if like == "" && tag == "" {
		return l.documentclasses
	} else {
		like = strings.ToLower(like)
		tag = strings.ToLower(tag)
	}
	var itemsThatMatch []Documentclass
	for _, item := range l.documentclasses {
		if fuzzy.Match(like, item.Name) && (tag == "" || hasTag(item.Label, tag)) {
			itemsThatMatch = append(itemsThatMatch, item)
		}
	}
	return itemsThatMatch
}

// Case insensitive fuzzy match.
func (l *Ltxref) FilterPackages(like string, tag string) []Package {
	if like == "" && tag == "" {
		return l.packages
	} else {
		like = strings.ToLower(like)
		tag = strings.ToLower(tag)
	}
	var itemsThatMatch []Package
	for _, item := range l.packages {
		if fuzzy.Match(like, item.Name) && (tag == "" || hasTag(item.Label, tag)) {
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
