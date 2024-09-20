package agenda_jobs

import (
	"context"
	"fmt"
)

func SubscribeToAlreadyExistingAgenda(ctx context.Context, args ...interface{}) error {
	fmt.Printf("Subscribing to an already existing agenda\n")
	return nil
}
