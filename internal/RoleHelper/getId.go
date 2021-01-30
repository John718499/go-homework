package RoleHelper

func SetRoleSkillID(roles *[]Role, u *Role) {

	for _, roleValue := range *roles {
		if u.ID <= roleValue.ID {
			u.ID = roleValue.ID + 1
		}

		if len(u.Skills) > 0 {
			for _, skillValue := range roleValue.Skills {
				if u.Skills[0].ID <= skillValue.ID {
					u.Skills[0].ID = skillValue.ID + 1
				}
			}
		}
	}

	for index := 1; index < len(u.Skills); index++ {
		u.Skills[index].ID = u.Skills[index-1].ID + 1
	}

	return
}
