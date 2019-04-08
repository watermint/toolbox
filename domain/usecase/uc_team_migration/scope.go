package uc_team_migration

type ScopeOpt func(opt *scopeOpts) *scopeOpts
type scopeOpts struct {
	membersAllExceptAdmin    bool
	membersSpecifiedEmail    []string
	teamFoldersAll           bool
	teamFoldersSpecifiedName []string
	groupsOnlyRelated        bool
	keepDesktopSessions      bool
}

func MembersAllExceptAdmin() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.membersAllExceptAdmin = true
		return opt
	}
}
func MembersSpecifiedEmail(members []string) ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.membersSpecifiedEmail = members
		return opt
	}
}
func TeamFoldersAll() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.teamFoldersAll = true
		return opt
	}
}
func TeamFoldersSpecifiedName(name []string) ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.teamFoldersSpecifiedName = name
		return opt
	}
}
func GroupsOnlyRelated() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.groupsOnlyRelated = true
		return opt
	}
}
func KeepDesktopSessions() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.keepDesktopSessions = true
		return opt
	}
}
