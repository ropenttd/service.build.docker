package models

import (
	"github.com/satori/go.uuid"
	"time"
)

// FIXME could this be better implemented as a custom type?
const (
	GameVersionTrainTag_UNKNOWN           = "unknown"
	GameVersionTrainTag_NIGHTLY           = "nightly"
	GameVersionTrainTag_ALPHA             = "alpha"
	GameVersionTrainTag_BETA              = "beta"
	GameVersionTrainTag_RELEASE_CANDIDATE = "rc"
	GameVersionTrainTag_STABLE            = "stable"
	GameVersionTrainTag_SPECIAL           = "special"
)

type OpenttdGameVersion struct {
	ID uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();not null"`

	Version string `json:"version" gorm:"index;not null"`

	Released  time.Time  `json:"released" gorm:"index;not null"`
	FirstSeen time.Time  `json:"seen_first" gorm:"not null"`
	LastSeen  *time.Time `json:"seen_last"`
	// Trains map as above.
	Train string `json:"train" gorm:"index;not null"`

	CreatedAt time.Time  `json:"created" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted,omitempty"`
}

func ListLatestVersions() (versions []*OpenttdGameVersion) {
	// SELECT DISTINCT ON (train) * FROM "openttd_game_versions" WHERE "openttd_game_versions"."deleted_at" IS NULL ORDER BY train, released desc
	err := GetDB().Order("train, released desc").Select("DISTINCT ON (train) *").Find(&versions).Error
	if err != nil {
		return nil
	}
	return versions
}

func ListVersions(search *OpenttdGameVersion) (versions []*OpenttdGameVersion) {
	err := GetDB().Where(search).Order("released desc").Find(&versions).Error
	if err != nil {
		return nil
	}
	return versions
}
