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

//tag related structures for free text search
const (
	//ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
)

var tagTrie *trie

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) prefix(word string) []string {
	tagList := make([]string, 0)
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return tagList
		}
		current = current.childrens[index]
	}
	//now that we have the node i.e where the prefix points to, we now do a regular DFS to build the rest of the string
	var dfs func(curNode *trieNode, slate string)
	dfs = func(curNode *trieNode, slate string) {
		if curNode.isWordEnd {
			tagList = append(tagList, strings.ToTitle(slate))
			return
		}
		if curNode == nil {
			return
		}
		for i := 0; i < ALBHABET_SIZE; i++ {
			if curNode.childrens[i] != nil {
				dfs(curNode.childrens[i], slate+string(byte(i+'a')))
			}
		}
	}
	dfs(current, word)
	return tagList
}

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

func buildHashGraph() {
	//my structure is a hash map with tag as the key and the value being a struct that contains the ancestor and list of descendants
	inputTemplates = make([]template, 2)
	tagMap = make(map[string]tag)
	tagMap["Triangles"] = tag{}
	tagMap["Rectangles"] = tag{descendants: []string{"Squares"}}
	tagMap["Squares"] = tag{ancestor: "Rectangles"}
	tagMap["Pets"] = tag{descendants: []string{"Rabbits", "Dogs"}}
	tagMap["Dogs"] = tag{ancestor: "Pets"}
	tagMap["Rabbits"] = tag{ancestor: "Pets"}
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

	tagTrie = initTrie()
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
		// if len(tagMap[itm].ancestor) > 0 {
		// 	queue = append(queue, tagMap[itm].ancestor)
		// }
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
			log.Println("TAG", tagName, "Name", tmpl.name, "Tags ", tmpl.tags)
		}
	}
	return
}

func main() {
	buildHashGraph()
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
