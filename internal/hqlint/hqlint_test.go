package hqlint

// Structural lint for the soul-hq headquarters layout.
//
// Enforces the invariants promised in 00-GENESIS/how-we-work.md: the five
// areas exist with their READMEs, research topics carry a component and legal
// non-terminal states (no graduated/abandoned folder lingers), journey
// episodes are contiguously numbered from 0001 with a legal component tag and
// indexed, every episode records its reversal condition, every component has
// its design folder with an index, and relative markdown links inside the
// live areas resolve. The frozen 99-ARCHIVE subtree is permitted and excluded
// from the link check — it is superseded material, not the live design.
// Links that leave the repository (sibling component checkouts such as
// ../soulstream/specs/...) are treated like external URLs and not checked.

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var (
	areas        = []string{"00-GENESIS", "01-RESEARCH", "02-DESIGN", "03-IMPLEMENTATION", "04-JOURNEY", "99-ARCHIVE"}
	genesisFiles = []string{"README.md", "vision.md", "constitution.md", "how-we-work.md", "rationale.md"}
	components   = []string{"soulstream", "soulstream-core", "soulstream-workloads", "soulstream-identity", "soulstream-idp", "soulstream-shell", "soulstream-inference"}
	legalTags    = map[string]bool{"soulstream": true, "soulrealm": true, "soulidentity": true, "soulnode": true, "soulfold": true, "soulhelm": true, "ecosystem": true, "core": true, "workloads": true, "identity": true, "idp": true, "shell": true, "mcp": true, "cli": true, "inference": true}
	legalStates  = map[string]bool{"active": true, "graduated": true, "abandoned": true}
	terminal     = map[string]bool{"graduated": true, "abandoned": true}
	nonEpisode   = map[string]bool{"README.md": true, "TEMPLATE.md": true}

	episodeRe   = regexp.MustCompile(`^(\d{4})-([a-z]+)-[a-z0-9-]+\.md$`)
	stateRe     = regexp.MustCompile(`(?m)^\*\*State:\*\* *(\S+)`)
	componentRe = regexp.MustCompile(`(?m)^\*\*Component:\*\* *(\S+)`)
	linkRe      = regexp.MustCompile(`\[[^\]]*\]\(([^)\s]+)\)`)
)

// repoRoot resolves the repository root from this test file's own location
// (<root>/internal/hqlint/hqlint_test.go), so the lint works from any working
// directory, locally and in CI.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot determine caller file")
	}
	return filepath.Join(filepath.Dir(file), "..", "..")
}

func mustFile(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		t.Errorf("missing required file: %s", path)
	}
}

// episodes lists the episode files in 04-JOURNEY (everything that is not a
// README/TEMPLATE), sorted by name.
func episodes(t *testing.T, root string) []string {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "04-JOURNEY"))
	if err != nil {
		t.Fatalf("reading 04-JOURNEY: %v", err)
	}
	var eps []string
	for _, e := range entries {
		if e.IsDir() || nonEpisode[e.Name()] {
			continue
		}
		eps = append(eps, e.Name())
	}
	sort.Strings(eps)
	return eps
}

func TestAreasExistWithReadmes(t *testing.T) {
	root := repoRoot(t)
	mustFile(t, filepath.Join(root, "README.md"))
	for _, area := range areas {
		mustFile(t, filepath.Join(root, area, "README.md"))
	}
	for _, name := range genesisFiles {
		mustFile(t, filepath.Join(root, "00-GENESIS", name))
	}
	mustFile(t, filepath.Join(root, "01-RESEARCH", "TEMPLATE.md"))
	mustFile(t, filepath.Join(root, "04-JOURNEY", "TEMPLATE.md"))
	for _, c := range components {
		mustFile(t, filepath.Join(root, "02-DESIGN", c, "README.md"))
	}
}

func TestConstitutionIsTheMergedOne(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(repoRoot(t), "00-GENESIS", "constitution.md"))
	if err != nil {
		t.Fatalf("reading constitution: %v", err)
	}
	if !strings.Contains(string(data), "# The Soul Constitution") {
		t.Error("constitution missing its '# The Soul Constitution' heading")
	}
}

func TestResearchTopicsHaveComponentAndLegalNonterminalStates(t *testing.T) {
	research := filepath.Join(repoRoot(t), "01-RESEARCH")
	entries, err := os.ReadDir(research)
	if err != nil {
		t.Fatalf("reading 01-RESEARCH: %v", err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(research, e.Name(), "README.md"))
		if err != nil {
			t.Errorf("%s: research topic without README.md", e.Name())
			continue
		}
		text := string(data)
		if !strings.HasPrefix(strings.TrimLeft(text, " \n\t"), "# ") {
			t.Errorf("%s: README lacks a title", e.Name())
		}
		if !strings.Contains(text, "## Abstract") {
			t.Errorf("%s: README lacks an Abstract section", e.Name())
		}
		if m := componentRe.FindStringSubmatch(text); m == nil {
			t.Errorf("%s: README lacks a '**Component:** ...' line", e.Name())
		} else if !legalTags[m[1]] {
			t.Errorf("%s: illegal component %q", e.Name(), m[1])
		}
		m := stateRe.FindStringSubmatch(text)
		if m == nil {
			t.Errorf("%s: README lacks a '**State:** ...' line", e.Name())
			continue
		}
		state := m[1]
		switch {
		case !legalStates[state]:
			t.Errorf("%s: illegal state %q", e.Name(), state)
		case terminal[state]:
			t.Errorf("%s: state %q is terminal but the folder lingers — "+
				"/research-graduate removes the topic folder on every outcome", e.Name(), state)
		}
	}
}

func TestJourneyEpisodesNumberedContiguouslyWithLegalComponents(t *testing.T) {
	eps := episodes(t, repoRoot(t))
	var nums []int
	for _, name := range eps {
		m := episodeRe.FindStringSubmatch(name)
		if m == nil {
			t.Errorf("file in 04-JOURNEY that is not an NNNN-<component>-<slug>.md episode: %s", name)
			continue
		}
		if !legalTags[m[2]] {
			t.Errorf("%s: illegal component tag %q", name, m[2])
		}
		n, _ := strconv.Atoi(m[1])
		nums = append(nums, n)
	}
	seen := map[int]bool{}
	for _, n := range nums {
		if seen[n] {
			t.Errorf("duplicate episode number: %04d", n)
		}
		seen[n] = true
	}
	sort.Ints(nums)
	for i, n := range nums {
		if n != i+1 {
			t.Errorf("episode numbers not contiguous from 0001: %v", nums)
			break
		}
	}
}

func TestJourneyEpisodesAreIndexed(t *testing.T) {
	root := repoRoot(t)
	data, err := os.ReadFile(filepath.Join(root, "04-JOURNEY", "README.md"))
	if err != nil {
		t.Fatalf("reading 04-JOURNEY/README.md: %v", err)
	}
	index := string(data)
	for _, name := range episodes(t, root) {
		if !strings.Contains(index, name) {
			t.Errorf("episode missing from the 04-JOURNEY/README.md index: %s", name)
		}
	}
}

func TestJourneyEpisodesRecordReversalCondition(t *testing.T) {
	root := repoRoot(t)
	dir := filepath.Join(root, "04-JOURNEY")
	for _, name := range episodes(t, root) {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Errorf("%s: %v", name, err)
			continue
		}
		if !strings.Contains(string(data), "Reversal condition:") {
			t.Errorf("episode without the required 'Reversal condition:' line: %s "+
				"(see 04-JOURNEY/TEMPLATE.md)", name)
		}
	}
}

func TestRelativeLinksResolve(t *testing.T) {
	root, err := filepath.Abs(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	var broken []string
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			switch d.Name() {
			case "99-ARCHIVE", ".git", ".claude", "internal":
				if path != root {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, m := range linkRe.FindAllStringSubmatch(string(data), -1) {
			target := m[1]
			if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") ||
				strings.HasPrefix(target, "mailto:") || strings.HasPrefix(target, "#") {
				continue
			}
			p, _, _ := strings.Cut(target, "#")
			if p == "" {
				continue
			}
			abs := filepath.Clean(filepath.Join(filepath.Dir(path), p))
			if !strings.HasPrefix(abs, root+string(filepath.Separator)) {
				// leaves the repository: a sibling component checkout — not ours to check
				continue
			}
			if _, statErr := os.Stat(abs); statErr != nil {
				rel, _ := filepath.Rel(root, path)
				broken = append(broken, rel+" -> "+target)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking the repository: %v", err)
	}
	if len(broken) > 0 {
		t.Errorf("broken relative markdown links:\n%s", strings.Join(broken, "\n"))
	}
}
