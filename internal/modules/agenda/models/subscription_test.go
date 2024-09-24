package agenda_models

import (
	"reflect"
	"testing"
)

func TestNewSubscription(t *testing.T) {
	t.Run("It instantiates a  new Subscription with relevant fields", func(t *testing.T) {
		var email string = "toto@gmail.com"
		var agendaId uint = 3
		expected := Subscription{
			AgendaId: agendaId,
			Email:    email,
		}
		if got := NewSubscription(agendaId, email); !reflect.DeepEqual(*got, expected) {
			t.Errorf("NewSubscription() = %v, want %v", *got, expected)
		}
	})
}
