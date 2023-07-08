package validators

import (
	"strconv"

	"github.com/google/uuid"
)

type AdminActionValidator struct{}

func (v *AdminActionValidator) PriceValidator(userid string, price string) error {
	_, err := uuid.Parse(userid)
	if err != nil {
		return err
	}

	p, err := strconv.Atoi(price)
	if err != nil || p <= 0 {
		return err
	}

	return nil
}

func (v *AdminActionValidator) VerifyTwoUUID(userid string, targetid string) error {
	_, err := uuid.Parse(userid)
	if err != nil {
		return err
	}

	_, err = uuid.Parse(targetid)
	if err != nil {
		return err
	}

	return nil
}
