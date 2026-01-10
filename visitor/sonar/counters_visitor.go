package sonar

import "github.com/gsixo/gocognit/visitor"

type VisitorWithCounters struct {
	complexityCounter visitor.Counter
	nestingCounter    visitor.Counter
}

func (w *VisitorWithCounters) IncComplexityCounterWithDelta(delta uint64) {
	w.complexityCounter.Inc(delta)
}

func (w *VisitorWithCounters) DecComplexityCounter() {
	w.complexityCounter.Dec()
}

func (w *VisitorWithCounters) LoadComplexityCounter() uint64 {
	return w.complexityCounter.Load()
}

func (w *VisitorWithCounters) IncNestingCounterWithDelta(delta uint64) {
	w.nestingCounter.Inc(delta)
}

func (w *VisitorWithCounters) DecNestingCounter() {
	w.nestingCounter.Dec()
}

func (w *VisitorWithCounters) LoadNestingCounter() uint64 {
	return w.nestingCounter.Load()
}
