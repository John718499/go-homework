package RoleHelper

type Role struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	Skills  []Skill `json:"skills"`
}
