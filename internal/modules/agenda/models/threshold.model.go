package agenda_models

const ThresholdAgendaType = "treshold"

type ThresholdAgendaConfiguration struct {
	PortID  int
	Maximum int
	Minimum int
}

type ThresholdAgenda struct {
	Base
	Configuration ThresholdAgendaConfiguration
}

func NewThresholdAgenda() *ThresholdAgenda {
	return &ThresholdAgenda{Base: Base{Type: ThresholdAgendaType}}
}
