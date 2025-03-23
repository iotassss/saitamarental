package domain

import "fmt"

type ChintaiSoubaHikakuCriteria string

func NewChintaiSoubaHikakuCriteria(criteria string) (ChintaiSoubaHikakuCriteria, error) {
	switch criteria {
	case "larger":
		return ChintaiSoubaHikakuCriteria(criteria), nil
	case "closer_to_shinjuku":
		return ChintaiSoubaHikakuCriteria(criteria), nil
	case "closer_to_tokyo":
		return ChintaiSoubaHikakuCriteria(criteria), nil
	case "larger_closer_to_shinjuku":
		return ChintaiSoubaHikakuCriteria(criteria), nil
	case "larger_closer_to_tokyo":
		return ChintaiSoubaHikakuCriteria(criteria), nil
	default:
		return ChintaiSoubaHikakuCriteria(""), fmt.Errorf("invalid criteria")
	}
}

func (c ChintaiSoubaHikakuCriteria) String() string {
	return string(c)
}
