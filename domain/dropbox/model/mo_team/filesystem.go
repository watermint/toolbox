package mo_team

type TeamFileSystemType string

func (z TeamFileSystemType) Version() FileSystemVersion {
	switch z {
	case TeamFileSystem2016TeamFolder:
		return FileSystemVersion{
			Version:                              string(z),
			ReleaseYear:                          2016,
			HasDistinctMemberHomes:               false,
			HasTeamSharedDropbox:                 false,
			IsTeamFolderApiSupported:             true,
			IsPathRootRequiredToAccessTeamFolder: false,
		}
	case TeamFileSystem2017TeamSpace:
		return FileSystemVersion{
			Version:                              string(z),
			ReleaseYear:                          2017,
			HasDistinctMemberHomes:               false,
			HasTeamSharedDropbox:                 true,
			IsTeamFolderApiSupported:             false,
			IsPathRootRequiredToAccessTeamFolder: true,
		}
	case TeamFileSystem2023TeamSpace:
		return FileSystemVersion{
			Version:                              string(z),
			ReleaseYear:                          2023,
			HasDistinctMemberHomes:               true,
			HasTeamSharedDropbox:                 false,
			IsTeamFolderApiSupported:             true,
			IsPathRootRequiredToAccessTeamFolder: true,
		}
	default:
		return FileSystemVersion{
			Version:                              string(z),
			ReleaseYear:                          0,
			HasDistinctMemberHomes:               false,
			HasTeamSharedDropbox:                 false,
			IsTeamFolderApiSupported:             false,
			IsPathRootRequiredToAccessTeamFolder: false,
		}
	}
}

const (
	// TeamFileSystem2016TeamFolder Team folder
	// Release year: 2016, https://blog.dropbox.com/topics/company/announcing-adminx
	TeamFileSystem2016TeamFolder TeamFileSystemType = "2016_team_folder"

	// TeamFileSystem2017TeamSpace Team space
	// Release year: 2017, https://github.com/dropbox/dropbox-api-spec/commit/6194bea0f324d3894e91c5c637e5df9fd9392140
	TeamFileSystem2017TeamSpace TeamFileSystemType = "2017_team_space"

	// TeamFileSystem2023TeamSpace Updated Team space
	// Release year: 2023, https://dropbox.tech/developers/api-updates-to-better-support-team-spaces
	TeamFileSystem2023TeamSpace TeamFileSystemType = "2023_updated_team_space"

	TeamFileSystemUnknown TeamFileSystemType = "unknown"
)

func IdentifyFileSystemType(f *Feature) TeamFileSystemType {
	switch {
	case f.HasDistinctMemberHomes && !f.HasTeamSharedDropbox:
		return TeamFileSystem2023TeamSpace
	case f.HasDistinctMemberHomes && f.HasTeamSharedDropbox:
		return TeamFileSystem2017TeamSpace
	case !f.HasDistinctMemberHomes && !f.HasTeamSharedDropbox:
		return TeamFileSystem2016TeamFolder
	default:
		return TeamFileSystemUnknown
	}
}

type FileSystemVersion struct {
	Version                              string `json:"version"`
	ReleaseYear                          int    `json:"release_year"`
	HasDistinctMemberHomes               bool   `json:"has_distinct_member_homes"`
	HasTeamSharedDropbox                 bool   `json:"has_team_shared_dropbox"`
	IsTeamFolderApiSupported             bool   `json:"is_team_folder_api_supported"`
	IsPathRootRequiredToAccessTeamFolder bool   `json:"is_path_root_required_to_access_team_folder"`
}
