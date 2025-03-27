package domain

import "fmt"

type ChintaiSoubaHikakuCriteria string

const (
	CriteriaLarger                 ChintaiSoubaHikakuCriteria = "larger"
	CriteriaCloserToShinjuku       ChintaiSoubaHikakuCriteria = "closer_to_shinjuku"
	CriteriaCloserToTokyo          ChintaiSoubaHikakuCriteria = "closer_to_tokyo"
	CriteriaLargerCloserToShinjuku ChintaiSoubaHikakuCriteria = "larger_closer_to_shinjuku"
	CriteriaLargerCloserToTokyo    ChintaiSoubaHikakuCriteria = "larger_closer_to_tokyo"
)

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
