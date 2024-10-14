package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// Graph structure to represent the directed graph
type Graph struct {
	nodes map[string][]string
}

// NewGraph creates a new directed graph
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// AddEdge adds an edge from function `from` to function `to`
func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
}

// PrintGraph prints the directed graph
func (g *Graph) PrintGraph() {
	for from, toList := range g.nodes {
		fmt.Printf("%s -> %v\n", from, toList)
	}
}

// ParseGoFiles parses Go files in a directory or a single Go file and returns a function call graph
func ParseGoFiles(path string) (*Graph, error) {
	fset := token.NewFileSet()

	// Check if the provided path is a directory or a single file
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		// Parse all Go files in the directory
		pkgs, err := parser.ParseDir(fset, path, func(info os.FileInfo) bool {
			return filepath.Ext(info.Name()) == ".go"
		}, parser.ParseComments)

		if err != nil {
			return nil, err
		}

		// Process the package and generate the graph
		graph := processPackages(pkgs)
		return graph, nil
	}

	// If the path is a file, parse just that file
	if filepath.Ext(path) == ".go" {
		// Parse the Go file
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		// Create and process the graph for a single file
		graph := NewGraph()
		processFile(node, graph)
		return graph, nil
	}

	return nil, fmt.Errorf("unsupported path: %s", path)
}

// processPackages processes parsed Go packages and updates the graph
func processPackages(pkgs map[string]*ast.Package) *Graph {
	graph := NewGraph()

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			processFile(file, graph)
		}
	}

	return graph
}

// processFile processes a single Go file and updates the graph
func processFile(file *ast.File, graph *Graph) {
	var currentFunc string

	// Walk the AST and inspect each node
	ast.Inspect(file, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			// Capture the current function name
			currentFunc = funcDecl.Name.Name
			if _, exists := graph.nodes[currentFunc]; !exists {
				graph.nodes[currentFunc] = []string{}
			}
		}

		if callExpr, ok := n.(*ast.CallExpr); ok {
			if fun, ok := callExpr.Fun.(*ast.Ident); ok {
				// Add the called function as an edge in the graph
				graph.AddEdge(currentFunc, fun.Name)
			}
		}

		return true
	})
}

func main() {
	// Provide the path to the Go source directory or file
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to a Go source directory or file")
		return
	}
	path := os.Args[1]

	// Parse the Go files and generate the function call graph
	graph, err := ParseGoFiles(path)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}

	// Print the function call graph
	graph.PrintGraph()
}
