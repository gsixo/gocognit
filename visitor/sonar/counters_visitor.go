package sonar

import "github.com/gsixo/gocognit/visitor"

type VisitorCounters struct {
	complexityCounter visitor.Counter
	nestingCounter    visitor.Counter
}

func NewVisitorCounters() visitor.VisitorCounters {
	return &VisitorCounters{
		complexityCounter: NewComplexityCounter(),
		nestingCounter:    NewNestingCounter(),
	}
}

func (w *VisitorCounters) IncComplexityCounterWithDelta(delta uint64) {
	w.complexityCounter.Inc(delta)
}

func (w *VisitorCounters) DecComplexityCounter() {
	w.complexityCounter.Dec()
}

func (w *VisitorCounters) LoadComplexityCounter() uint64 {
	return w.complexityCounter.Load()
}

func (w *VisitorCounters) IncNestingCounterWithDelta(delta uint64) {
	w.nestingCounter.Inc(delta)
}

func (w *VisitorCounters) DecNestingCounter() {
	w.nestingCounter.Dec()
}

func (w *VisitorCounters) LoadNestingCounter() uint64 {
	return w.nestingCounter.Load()
}

func (w *VisitorCounters) IncComplexityCounterWithPlusNestingCounterValue(delta uint64) {
	w.complexityCounter.Inc(w.nestingCounter.Load() + delta)
}

func (w *VisitorCounters) IncDecNestingCounterWithFnBetween(fn func()) {
	w.nestingCounter.Inc(1)
	fn()
	w.nestingCounter.Dec()
}
