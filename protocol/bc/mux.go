package bc

type Mux struct {
	body struct {
		Sources []ValueSource
		Program Program
		ExtHash Hash
	}
	witness struct {
		Destinations []ValueDestination
	}
}

const typeMux = "mux1"

func (Mux) Type() string            { return typeMux }
func (m *Mux) Body() interface{}    { return &m.body }
func (m *Mux) Witness() interface{} { return &m.witness }

func (m *Mux) Destinations() []ValueDestination {
	return m.witness.Destinations
}

func newMux(sources []ValueSource, program Program) *Mux {
	m := new(Mux)
	m.body.Sources = sources
	m.body.Program = program
	return m
}
