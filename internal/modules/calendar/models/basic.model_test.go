package calendar_models

import (
	"fmt"
	"testing"

	"github.com/troptropcontent/what_the_tide/database"
)

func TestNewBasicAgenda(t *testing.T) {
	t.Run("It assigns the correct type to the record", func(t *testing.T) {
		if got := NewBasicAgenda(); got.Type != BasicAgendaType {
			t.Errorf("NewBasicAgenda() should set the type to %v, got %v", BasicAgendaType, got.Type)
		}
	})
}

func TestAssociations(t *testing.T) {
	t.Run("Configuration", func(t *testing.T) {
		database.MustInit()
		transaction := database.DB.Begin()
		calendar := NewBasicAgenda()
		calendar.Configuration = BasicAgendaConfiguration{
			PortID: 123,
		}
		result := transaction.Create(&calendar)
		fmt.Printf("COUCOU")
		fmt.Printf("Rows affected: %d\n", result.RowsAffected)
		t.Cleanup(func() { transaction.Rollback() })
		// if got := NewBasicAgenda(); got.Type != BasicAgendaType {
		// 	t.Errorf("NewBasicAgenda() should set the type to %v, got %v", BasicAgendaType, got.Type)
		// }
	})
}
