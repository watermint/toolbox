package uc_team_migration

const (
	PeerNameActorTeamAAdmin01 = "test-migration-a01"
	PeerNameActorTeamBAdmin01 = "test-migration-b01"
	PeerNameActorIndividual01 = "test-migration-i01"
)

type Actors struct {
	TeamAAdmin01  string `json:"team_a_admin_01"`
	TeamAMember02 string `json:"team_a_member_02"`
	TeamAMember03 string `json:"team_a_member_03"`
	TeamAMember04 string `json:"team_a_member_04"`
	TeamBAdmin01  string `json:"team_b_admin_01"`
	Individual01  string `json:"individual_01"`
}
