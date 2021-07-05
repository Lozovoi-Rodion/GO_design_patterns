package DIP

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	//
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

// Low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		from:         parent,
		relationship: Parent,
		to:           child,
	})
}

// High-level module
type Research struct {
	// break DIP
	//relationships Relationships
	browser RelationshipBrowser
}

// break DIP
//func (r *Research) Investigate() {
//	relations := r.relationships.relations
//	for _, rel := range relations {
//		if rel.from.name == "John" && rel.relationship == Parent {
//			fmt.Println("John has a child called ", rel.to.name)
//		}
//	}
//}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}
