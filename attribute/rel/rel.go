package rel

import "github.com/doojin/gohtml/attribute"

// Rel represents rel attribute value
type Rel struct {
	attribute.Attribute
}

// New returns a new instance of Rel
func New(value string) Rel {
	rel := Rel{}
	rel.Value = value
	return rel
}

var (
	Answer = New("answer")
	Chapter = New("chapter")
	CoWorker = New("co-worker")
	Colleague = New("colleague")
	Contact = New("contact")
	Details = New("details")
	Edit = New("edit")
	Friend = New("friend")
	Question = New("question")
	Archives = New("archives")
	Author = New("author")
	Bookmark = New("bookmark")
	First = New("first")
	Help = New("help")
	Index = New("index")
	Last = New("last")
	License = New("license")
	Me = New("me")
	Next = New("next")
	NoFollow = New("nofollow")
	NoReferrer = New("noreferrer")
	PreFetch = New("prefetch")
	Prev = New("prev")
	Search = New("search")
	Sidebar = New("sidebar")
	Tag = New("tag")
	Up = New("up")
)

