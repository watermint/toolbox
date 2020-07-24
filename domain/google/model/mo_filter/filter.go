package mo_filter

import "encoding/json"

type Filter struct {
	Raw                  json.RawMessage
	Id                   string `json:"id" path:"id"`
	CriteriaFrom         string `json:"criteria_from" path:"criteria.from"`
	CriteriaTo           string `json:"criteria_to" path:"criteria.to"`
	CriteriaSubject      string `json:"criteria_subject" path:"criteria.subject"`
	CriteriaQuery        string `json:"criteria_query" path:"criteria.query"`
	CriteriaNegatedQuery string `json:"criteria_negated_query" path:"criteria.negatedQuery"`
}
