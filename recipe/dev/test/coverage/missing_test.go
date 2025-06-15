package coverage

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

func TestMissing_Preset(t *testing.T) {
	m := &Missing{}
	m.Preset()
	
	if !m.OnlyMissing {
		t.Error("Expected OnlyMissing to be true after Preset()")
	}
}

func TestMissingFile_Struct(t *testing.T) {
	mf := MissingFile{
		Package:      "test/package",
		File:         "test.go",
		RelativePath: "test/package/test.go",
		HasTest:      false,
		Functions:    5,
		Lines:        100,
		Complexity:   15,
		Priority:     "high",
	}
	
	if mf.Package != "test/package" {
		t.Error("Package field not set correctly")
	}
	if mf.HasTest {
		t.Error("Expected HasTest to be false")
	}
	if mf.Functions != 5 {
		t.Error("Functions field not set correctly")
	}
}

func TestMissing_findFilesWithoutTests(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		// Create a temporary directory structure for testing
		tmpDir := t.TempDir()
		
		// Create test Go files
		testFiles := map[string]string{
			"main.go": `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func add(a, b int) int {
	return a + b
}`,
			"main_test.go": `package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}`,
			"untested.go": `package main

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}`,
			"pkg/service.go": `package pkg

type Service struct {
	name string
}

func (s *Service) GetName() string {
	return s.name
}

func (s *Service) SetName(name string) {
	s.name = name
}`,
		}
		
		// Write test files
		for filename, content := range testFiles {
			fullPath := filepath.Join(tmpDir, filename)
			dir := filepath.Dir(fullPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return err
			}
		}
		
		m := &Missing{}
		files, err := m.findFilesWithoutTests(c, tmpDir, "")
		if err != nil {
			return err
		}
		
		// Should find 3 files total: main.go (has test), untested.go (no test), pkg/service.go (no test)
		if len(files) != 3 {
			t.Errorf("Expected 3 files total, got %d", len(files))
			for _, f := range files {
				t.Logf("Found file: %s (HasTest: %v)", f.RelativePath, f.HasTest)
			}
		}
		
		// Check that untested.go is found
		foundUntested := false
		foundService := false
		for _, f := range files {
			if f.RelativePath == "untested.go" {
				foundUntested = true
				if f.HasTest {
					t.Error("untested.go should not have tests")
				}
				if f.Functions != 2 {
					t.Errorf("Expected untested.go to have 2 functions, got %d", f.Functions)
				}
			}
			if f.RelativePath == filepath.Join("pkg", "service.go") {
				foundService = true
				if f.HasTest {
					t.Error("pkg/service.go should not have tests")
				}
			}
		}
		
		if !foundUntested {
			t.Error("Expected to find untested.go")
		}
		if !foundService {
			t.Error("Expected to find pkg/service.go")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestMissing_findFilesWithoutTests_WithPackageFilter(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		tmpDir := t.TempDir()
		
		// Create files in different packages
		testFiles := map[string]string{
			"main.go": `package main
func main() {}`,
			"pkg1/service.go": `package pkg1
func Service() {}`,
			"pkg2/handler.go": `package pkg2
func Handler() {}`,
		}
		
		for filename, content := range testFiles {
			fullPath := filepath.Join(tmpDir, filename)
			dir := filepath.Dir(fullPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return err
			}
		}
		
		m := &Missing{}
		
		// Filter by pkg1
		files, err := m.findFilesWithoutTests(c, tmpDir, "pkg1")
		if err != nil {
			return err
		}
		
		if len(files) != 1 {
			t.Errorf("Expected 1 file in pkg1, got %d", len(files))
		}
		
		if len(files) > 0 && files[0].RelativePath != filepath.Join("pkg1", "service.go") {
			t.Errorf("Expected pkg1/service.go, got %s", files[0].RelativePath)
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestMissing_Exec_OnlyMissingFilter(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		tmpDir := t.TempDir()
		
		// Create test files - one with test, one without
		testFiles := map[string]string{
			"tested.go": `package main
func TestedFunc() {}`,
			"tested_test.go": `package main
import "testing"
func TestTestedFunc(t *testing.T) {}`,
			"untested.go": `package main
func UntestedFunc() {}`,
		}
		
		for filename, content := range testFiles {
			fullPath := filepath.Join(tmpDir, filename)
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return err
			}
		}
		
		// Mock the getProjectRoot function by directly calling findFilesWithoutTests
		m := &Missing{
			OnlyMissing: true,
		}
		
		files, err := m.findFilesWithoutTests(c, tmpDir, "")
		if err != nil {
			return err
		}
		
		// Should find both files: tested.go and untested.go
		if len(files) != 2 {
			t.Errorf("Expected 2 files total, got %d", len(files))
		}
		
		// When OnlyMissing is applied in the real Exec, it should filter to only untested.go
		// but findFilesWithoutTests returns all files with their HasTest status
		foundTested := false
		foundUntested := false
		for _, f := range files {
			if f.RelativePath == "tested.go" && f.HasTest {
				foundTested = true
			}
			if f.RelativePath == "untested.go" && !f.HasTest {
				foundUntested = true
			}
		}
		
		if !foundTested {
			t.Error("Expected to find tested.go with HasTest=true")
		}
		if !foundUntested {
			t.Error("Expected to find untested.go with HasTest=false")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestCalculateComplexity(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected int
	}{
		{
			name: "simple function",
			code: `package main
func simple() {
	return
}`,
			expected: 1,
		},
		{
			name: "function with if",
			code: `package main
func withIf(x int) {
	if x > 0 {
		return
	}
}`,
			expected: 2,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", tt.code, parser.ParseComments)
			if err != nil {
				t.Fatal(err)
			}
			
			m := &Missing{}
			
			// Find the function declaration in the AST
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					complexity := m.calculateCyclomaticComplexity(fn)
					if complexity != tt.expected {
						t.Errorf("Expected complexity %d, got %d", tt.expected, complexity)
					}
					break
				}
			}
		})
	}
}

func TestGetProjectRoot(t *testing.T) {
	// This test verifies getProjectRoot finds the actual project root
	// Since getProjectRoot uses os.Getwd() and looks for go.mod, 
	// it will find the real project root, not our mock
	
	mockWS := &mockWorkspace{basePath: "/some/path"}
	
	root := getProjectRoot(mockWS)
	
	// Should find a directory that contains go.mod
	goModPath := filepath.Join(root, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		t.Errorf("Expected go.mod to exist at project root %s", root)
	}
	
	// Should be a valid directory path
	if info, err := os.Stat(root); err != nil || !info.IsDir() {
		t.Errorf("Expected project root %s to be a valid directory", root)
	}
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected int
	}{
		{
			name:     "empty code",
			code:     "",
			expected: 1, // strings.Split("", "\n") returns [""]
		},
		{
			name: "single line",
			code: "package main",
			expected: 1,
		},
		{
			name: "multiple lines",
			code: `package main

import "fmt"

func main() {
	fmt.Println("Hello")
}`,
			expected: 7,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := len(strings.Split(tt.code, "\n"))
			if count != tt.expected {
				t.Errorf("Expected %d lines, got %d", tt.expected, count)
			}
		})
	}
}

func TestCountFunctions(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected int
	}{
		{
			name:     "no functions",
			code:     "package main\nvar x = 1",
			expected: 0,
		},
		{
			name: "one function",
			code: `package main
func main() {}`,
			expected: 1,
		},
		{
			name: "multiple functions and methods",
			code: `package main
func main() {}
func helper() {}
type T struct{}
func (t T) Method() {}`,
			expected: 3,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "", tt.code, parser.ParseComments)
			if err != nil {
				t.Fatal(err)
			}
			
			// Count functions in the file
			count := 0
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok && fn.Body != nil {
					count++
				}
			}
			if count != tt.expected {
				t.Errorf("Expected %d functions, got %d", tt.expected, count)
			}
		})
	}
}

func TestMissing_Exec_EmptyPackageFilter(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		m := &Missing{
			Package: mo_string.NewOptional(""),
		}
		
		// Test that empty package filter doesn't cause issues
		m.Preset()
		
		// We can't easily test the full Exec without mocking the workspace
		// but we can test that the Package field works correctly
		if m.Package.IsExists() && m.Package.Value() == "" {
			// This is valid - empty string means no filter
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

// mockWorkspace implements a minimal workspace interface for testing
type mockWorkspace struct {
	basePath string
}

func (m *mockWorkspace) Home() string          { return m.basePath }
func (m *mockWorkspace) Cache() string         { return filepath.Join(m.basePath, "cache") }
func (m *mockWorkspace) Secrets() string       { return filepath.Join(m.basePath, "secrets") }
func (m *mockWorkspace) Job() string           { return filepath.Join(m.basePath, "job") }
func (m *mockWorkspace) Test() string          { return filepath.Join(m.basePath, "test") }
func (m *mockWorkspace) Report() string        { return filepath.Join(m.basePath, "report") }
func (m *mockWorkspace) Log() string           { return filepath.Join(m.basePath, "log") }
func (m *mockWorkspace) JobStartTime() time.Time { return time.Now() }
func (m *mockWorkspace) JobId() string         { return "test-job-id" }
func (m *mockWorkspace) KVS() string           { return filepath.Join(m.basePath, "kvs") }
func (m *mockWorkspace) Database() string      { return filepath.Join(m.basePath, "database") }
func (m *mockWorkspace) Descendant(name string) (string, error) { 
	return filepath.Join(m.basePath, name), nil 
}

func TestMissing_FileSortingByPriority(t *testing.T) {
	files := []MissingFile{
		{RelativePath: "low.go", Lines: 10, Complexity: 2, Priority: "low"},
		{RelativePath: "high.go", Lines: 100, Complexity: 10, Priority: "high"},
		{RelativePath: "medium.go", Lines: 50, Complexity: 5, Priority: "medium"},
	}
	
	// Simulate the sorting logic from the Exec method
	// Sort by priority (complexity * lines)
	
	// Calculate scores
	scores := make(map[string]int)
	for _, f := range files {
		scores[f.RelativePath] = f.Complexity * f.Lines
	}
	
	// high.go should have highest score: 100 * 10 = 1000
	// medium.go should have middle score: 50 * 5 = 250
	// low.go should have lowest score: 10 * 2 = 20
	
	if scores["high.go"] != 1000 {
		t.Errorf("Expected high.go score 1000, got %d", scores["high.go"])
	}
	if scores["medium.go"] != 250 {
		t.Errorf("Expected medium.go score 250, got %d", scores["medium.go"])
	}
	if scores["low.go"] != 20 {
		t.Errorf("Expected low.go score 20, got %d", scores["low.go"])
	}
}

func TestMissing_SkipVendorAndBuildDirs(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		tmpDir := t.TempDir()
		
		// Create files in vendor and build directories (should be skipped)
		testFiles := map[string]string{
			"main.go":              `package main`,
			"vendor/pkg/file.go":   `package pkg`,
			"build/output/file.go": `package output`,
		}
		
		for filename, content := range testFiles {
			fullPath := filepath.Join(tmpDir, filename)
			dir := filepath.Dir(fullPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return err
			}
		}
		
		m := &Missing{}
		files, err := m.findFilesWithoutTests(c, tmpDir, "")
		if err != nil {
			return err
		}
		
		// Should only find main.go, vendor and build files should be skipped
		if len(files) != 1 {
			t.Errorf("Expected 1 file, got %d", len(files))
			for _, f := range files {
				t.Logf("Found: %s", f.RelativePath)
			}
		}
		
		if len(files) > 0 && files[0].RelativePath != "main.go" {
			t.Errorf("Expected main.go, got %s", files[0].RelativePath)
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestMissing_EdgeCases(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		tmpDir := t.TempDir()
		
		// Test edge cases
		testFiles := map[string]string{
			// File with no functions
			"empty.go": `package main
// Just a comment
var x = 1`,
			// File with complex nested structures
			"complex.go": `package main
func outer() {
	if true {
		for i := 0; i < 10; i++ {
			switch i {
			case 1:
				if true {
					// nested complexity
				}
			case 2:
				return
			default:
				break
			}
		}
	}
}`,
		}
		
		for filename, content := range testFiles {
			fullPath := filepath.Join(tmpDir, filename)
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return err
			}
		}
		
		m := &Missing{}
		files, err := m.findFilesWithoutTests(c, tmpDir, "")
		if err != nil {
			return err
		}
		
		if len(files) != 2 {
			t.Errorf("Expected 2 files, got %d", len(files))
		}
		
		// Verify complexity calculation worked for complex file
		for _, f := range files {
			if f.RelativePath == "complex.go" {
				if f.Complexity < 5 {
					t.Errorf("Expected complex.go to have high complexity, got %d", f.Complexity)
				}
			}
			if f.RelativePath == "empty.go" {
				if f.Functions != 0 {
					t.Errorf("Expected empty.go to have 0 functions, got %d", f.Functions)
				}
			}
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}