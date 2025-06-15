package coverage

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Missing struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	Package mo_string.OptionalString `name:"package" desc:"Package to analyze (optional, defaults to entire project)"`
	OnlyMissing bool `name:"only-missing" desc:"Show only files without any tests"`
}

type MissingFile struct {
	Package     string `json:"package"`
	File        string `json:"file"`
	RelativePath string `json:"relative_path"`
	HasTest     bool   `json:"has_test"`
	Functions   int    `json:"functions"`
	Lines       int    `json:"lines"`
	Complexity  int    `json:"complexity"`
	Priority    string `json:"priority"`
}

func (z *Missing) Preset() {
	z.OnlyMissing = true
}

func (z *Missing) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	
	projectRoot := getProjectRoot(c.Workspace())
	l.Debug("Project root", esl.String("path", projectRoot))
	
	var packagePath string
	if z.Package.IsExists() {
		packagePath = z.Package.Value()
	}
	
	files, err := z.findFilesWithoutTests(c, projectRoot, packagePath)
	if err != nil {
		return err
	}
	
	// Sort by priority (complexity * lines)
	sort.Slice(files, func(i, j int) bool {
		scoreI := files[i].Complexity * files[i].Lines
		scoreJ := files[j].Complexity * files[j].Lines
		return scoreI > scoreJ
	})
	
	// Filter if only missing requested
	if z.OnlyMissing {
		filtered := make([]MissingFile, 0)
		for _, f := range files {
			if !f.HasTest {
				filtered = append(filtered, f)
			}
		}
		files = filtered
	}
	
	l.Info("Analysis complete", esl.Int("total_files", len(files)))
	
	// Display results
	ui.Info(app_msg.Raw(fmt.Sprintf("Files without tests (%d total):", len(files))))
	ui.Info(app_msg.Raw(strings.Repeat("=", 80)))
	ui.Info(app_msg.Raw(fmt.Sprintf("%-60s %8s %8s %8s %8s", "File", "Funcs", "Lines", "Complexity", "Priority")))
	ui.Info(app_msg.Raw(strings.Repeat("-", 80)))
	
	maxDisplay := 20
	if len(files) < maxDisplay {
		maxDisplay = len(files)
	}
	
	for i := 0; i < maxDisplay; i++ {
		file := files[i]
		ui.Info(app_msg.Raw(fmt.Sprintf("%-60s %8d %8d %8d %8s",
			file.RelativePath,
			file.Functions,
			file.Lines,
			file.Complexity,
			file.Priority,
		)))
	}
	
	return nil
}

func (z *Missing) findFilesWithoutTests(c app_control.Control, projectRoot, packageFilter string) ([]MissingFile, error) {
	l := c.Log()
	results := make([]MissingFile, 0)
	
	err := filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip non-Go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		
		// Skip test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}
		
		// Skip vendor and build directories
		if strings.Contains(path, "/vendor/") || strings.Contains(path, "/build/") {
			return nil
		}
		
		// Get relative path from project root
		relPath, err := filepath.Rel(projectRoot, path)
		if err != nil {
			return err
		}
		
		// Determine package path
		packagePath := filepath.Dir(relPath)
		if packagePath == "." {
			packagePath = ""
		}
		packagePath = strings.ReplaceAll(packagePath, string(filepath.Separator), "/")
		
		// Filter by package if specified
		if packageFilter != "" && !strings.Contains(packagePath, packageFilter) {
			return nil
		}
		
		// Check if test file exists
		testFile := strings.TrimSuffix(path, ".go") + "_test.go"
		_, statErr := os.Stat(testFile)
		hasTest := !os.IsNotExist(statErr)
		
		// Analyze the file
		analysis, err := z.analyzeFile(path)
		if err != nil {
			l.Warn("Failed to analyze file", esl.String("path", path), esl.Error(err))
			return nil // Continue processing other files
		}
		
		priority := z.calculatePriority(analysis.Functions, analysis.Lines, analysis.Complexity)
		
		results = append(results, MissingFile{
			Package:      packagePath,
			File:         filepath.Base(path),
			RelativePath: relPath,
			HasTest:      hasTest,
			Functions:    analysis.Functions,
			Lines:        analysis.Lines,
			Complexity:   analysis.Complexity,
			Priority:     priority,
		})
		
		return nil
	})
	
	return results, err
}

type FileAnalysis struct {
	Functions  int
	Lines      int
	Complexity int
}

func (z *Missing) analyzeFile(filePath string) (*FileAnalysis, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	
	analysis := &FileAnalysis{}
	
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Body != nil { // Only count functions with bodies
				analysis.Functions++
				analysis.Complexity += z.calculateCyclomaticComplexity(x)
			}
		}
		return true
	})
	
	// Count lines by reading the file
	content, err := os.ReadFile(filePath)
	if err == nil {
		analysis.Lines = len(strings.Split(string(content), "\n"))
	}
	
	return analysis, nil
}

func (z *Missing) calculateCyclomaticComplexity(fn *ast.FuncDecl) int {
	complexity := 1 // Base complexity
	
	ast.Inspect(fn, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
			complexity++
		case *ast.CaseClause:
			complexity++
		}
		return true
	})
	
	return complexity
}

func (z *Missing) calculatePriority(functions, lines, complexity int) string {
	score := complexity * lines + functions*10
	
	switch {
	case score > 1000:
		return "high"
	case score > 300:
		return "medium"
	default:
		return "low"
	}
}

func (z *Missing) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Missing{}, func(r rc_recipe.Recipe) {
		m := r.(*Missing)
		m.OnlyMissing = true
	})
}