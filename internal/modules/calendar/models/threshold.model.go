package calendar_models

const ThresholdCalendarType = "treshold"

type ThresholdCalendarConfiguration struct {
	PortID     int
	Maximum    int
	Minimum    int
	CalendarID uint
}

type ThresholdCalendar struct {
	Base
	Configuration ThresholdCalendarConfiguration
}

func NewThresholdCalendar() *ThresholdCalendar {
	return &ThresholdCalendar{Base: Base{Type: ThresholdCalendarType}}
}
