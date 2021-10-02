/*
Geometric Shapes
- Triangles
- Rectangles
- Squares
Animals
- Pets
- Dogs
- Rabbits
And these templates:
Name: dog in a square
Tags: Squares, Dogs
Name: rabbit in a rectangle
Tags: Rabbits, Rectangles
*/

package main

import (
	"log"
	"strings"
)

//tag related structures for faceted search
type tag struct {
	// tagName string being used as key
	ancestor       string   // one parent
	descendants    []string //multiple childen
	templates_list []*template
}

type template struct {
	name string
	tags []string
}

var tagMap map[string]tag
var inputTemplates []template
var sortedTagNames []string

//trie related structures for free text search
//num of alphabets
const A2Z = 26

var tagTrie *trie

type trieNode struct {
	child     [A2Z]*trieNode
	endOfWord bool
}

type trie struct {
	root *trieNode
}

func NewTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.child[index] == nil {
			current.child[index] = &trieNode{}
		}
		current = current.child[index]
	}
	current.endOfWord = true
}

func (t *trie) prefix(word string) []string {
	tagList := make([]string, 0)
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.child[index] == nil {
			return tagList
		}
		current = current.child[index]
	}
	//now that we have the node i.e where the prefix points to, we now do a regular DFS to build the rest of the string
	var dfs func(curNode *trieNode, slate string)
	dfs = func(curNode *trieNode, slate string) {
		if curNode.endOfWord {
			tagList = append(tagList, strings.ToTitle(slate))
			return
		}
		if curNode == nil {
			return
		}
		for i := 0; i < A2Z; i++ {
			if curNode.child[i] != nil {
				dfs(curNode.child[i], slate+string(byte(i+'a')))
			}
		}
	}
	dfs(current, word)
	return tagList
}

///End of trie related stuff

//cannot find a good way to sort the templates so storing them as an array of pointers ,
//if this operation gets too expensive then it can always be changed into a dictionary
func appendIfNotPresent(list_ptr []*template, template_ptr *template) []*template {
	found := false
	for _, tmp_ptr := range list_ptr {
		if tmp_ptr == template_ptr {
			found = true
			break
		}
	}
	if !found {
		list_ptr = append(list_ptr, template_ptr)
	}
	return list_ptr
}

func initialize() {
	//my structure is a hash map with tag as the key and the value being a struct that contains the ancestor and list of descendants
	inputTemplates = make([]template, 2)
	tagMap = make(map[string]tag)
	tagMap["Triangles"] = tag{}
	tagMap["Rectangles"] = tag{descendants: []string{"Squares"}}
	tagMap["Squares"] = tag{ancestor: "Rectangles"}
	tagMap["Pets"] = tag{descendants: []string{"Rabbits", "Dogs"}}
	tagMap["Dogs"] = tag{ancestor: "Pets"}
	tagMap["Rabbits"] = tag{ancestor: "Pets"}
	//the input templates are stored as a slice
	inputTemplates[0] = template{name: "dog in a square", tags: []string{"Squares", "Dogs"}}
	inputTemplates[1] = template{name: "rabbit in a rectangle", tags: []string{"Rabbits", "Rectangles"}}
	for template_idx, tmpl := range inputTemplates {
		for _, tagName := range tmpl.tags {
			if tg, ok := tagMap[tagName]; ok {
				tg.templates_list = appendIfNotPresent(tagMap[tagName].templates_list, &inputTemplates[template_idx])
				tagMap[tagName] = tg
			} else {
				log.Printf("tag is invalid %s\n", tagName)
			}

		}
	}
	for key := range tagMap {
		sortedTagNames = append(sortedTagNames, strings.ToLower(key))
	}
	//this trie is used only for text search
	tagTrie = NewTrie()
	for i := 0; i < len(sortedTagNames); i++ {
		tagTrie.insert(sortedTagNames[i])
	}
}

func GetTemplates(tag string) (results []*template) {

	//This is an adaptation of breadth first search
	tag = strings.Title(strings.ToLower(tag))
	visited := make(map[string]bool)
	results = make([]*template, 0)
	queue := []string{tag}
	for len(queue) > 0 {
		itm := queue[0] //pop
		queue = queue[1:]
		if visited[itm] {
			continue
		}
		if len(tagMap[itm].descendants) > 0 {
			queue = append(queue, tagMap[itm].descendants...)
		}
		for _, val := range tagMap[itm].templates_list {
			results = appendIfNotPresent(results, val)
		}
		visited[itm] = true
	}

	return
}

func freeTextSearch(str string) (results []*template) {
	str = strings.ToLower(str)
	results = make([]*template, 0)
	for _, tagName := range tagTrie.prefix(str) {
		results := GetTemplates(tagName)
		for _, tmpl := range results {
			log.Println("freeTextSearch", tagName, "Name", tmpl.name, "Tags ", tmpl.tags)
		}
	}
	return
}

func main() {
	initialize()
	results := GetTemplates("rectangles")
	for _, tmpl := range results {
		log.Println("Name", tmpl.name, "Tags ", tmpl.tags)
	}
	log.Println("=====================================")
	results = GetTemplates("SQUARES")
	for _, tmpl := range results {
		log.Println("Name", tmpl.name, "Tags ", tmpl.tags)
	}
	log.Println("=====================================")
	results = freeTextSearch("R")
	for _, tmpl := range results {
		log.Println("Name", tmpl.name, "Tags ", tmpl.tags)
	}
	log.Println("=====================================")
	results = freeTextSearch("s")
	for _, tmpl := range results {
		log.Println("Name", tmpl.name, "Tags ", tmpl.tags)
	}

}
