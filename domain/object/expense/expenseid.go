package expense

import (
	"fmt"

	"github.com/google/uuid"
)

type ID string

func newID() (ID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ID(""), fmt.Errorf("failed to generate uuid: %w", err)
	}
	return ID(uid.String()), nil
}
