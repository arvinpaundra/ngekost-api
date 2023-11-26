package response

import "github.com/arvinpaundra/ngekost-api/internal/entity"

type (
	KostRule struct {
		ID          string  `json:"id"`
		KostId      string  `json:"kost_id"`
		Title       string  `json:"title"`
		Description *string `json:"description"`
		Priority    string  `json:"priority"`
	}
)

func ToResponseKostRule(rule *entity.KostRule) *KostRule {
	return &KostRule{
		ID:          rule.ID,
		KostId:      rule.KostId,
		Title:       rule.Title,
		Description: rule.Description,
		Priority:    rule.Priority,
	}
}

func ToResponseKostRules(rules []*entity.KostRule) []*KostRule {
	var res []*KostRule

	for _, rule := range rules {
		res = append(res, &KostRule{
			ID:          rule.ID,
			KostId:      rule.KostId,
			Title:       rule.Title,
			Description: rule.Description,
			Priority:    rule.Priority,
		})
	}

	return res
}
