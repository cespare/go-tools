package pattern

import (
	"sync"
	"testing"
)

func TestParse(t *testing.T) {
	inputs := []string{
		`(Binding "name" _)`,
	}

	p := Parser{}
	defer p.xxx()
	for _, input := range inputs {
		if _, err := p.Parse(input); err != nil {
			t.Errorf("failed to parse %q: %s", input, err)
		}
	}
}

func TestUse(t *testing.T) {
	var p Pattern
	p.xxx()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
	wg.Wait()
}

func f() {
	_ = xxx(nil)
}
