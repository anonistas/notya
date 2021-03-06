// Copyright 2021-present Anon. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package models

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// Constant values of settings.
const (
	DefaultAppName   = "notya"
	SettingsName     = ".settings.json"
	DefaultEditor    = "vi"
	DefaultLocalPath = "notya"
)

// NotyaIgnoreFiles are those files that shouldn't
// be represented as note files.
var NotyaIgnoreFiles []string = []string{
	SettingsName,
	".DS_Store", // Darwin related.
	".git",
}

// Settings is a main structure model of application settings.
//
//  Example:
// ╭────────────────────────────────────────────────────╮
// │ Name: notya                                        │
// │ Editor: vi                                         │
// │ Local Path: /User/random-user/notya/.settings.json │
// │ Firebase Project ID: notya-98tf3                   │
// │ Firebase Account Key: /User/.../notya/key.json     │
// │ Firebase Collection: notya-notes                   │
// ╰────────────────────────────────────────────────────╯
type Settings struct {
	// Development related field, shouldn't be used in production.
	ID string `json:",omitempty"`

	Name string `json:"name" default:"notya"`

	// CLI base editor of application.
	Editor string `json:"editor" default:"vi"`

	// Local folder path for notes, independently from [~/notya/] folder.
	// Does same job as [FirebaseCollection] for local env.
	// Must be given full path, like: "./User/john-doe/.../my-notya-notes/"
	LocalPath string `json:"local_path" mapstructure:"local_path" survey:"local_path"`

	// The project id of your firebase project.
	FirebaseProjectID string `json:"fire_project_id,omitempty" mapstructure:"fire_project_id,omitempty" survey:"fire_project_id"`

	// The path of key of firebase-service account file.
	// Must be given full path, like: "./User/john-doe/.../..."
	FirebaseAccountKey string `json:"fire_account_key,omitempty" mapstructure:"fire_account_key,omitempty" survey:"fire_account_key"`

	// The concrete collection of nodes.
	// Does same job as [LocalPath] but has to take just name of collection.
	FirebaseCollection string `json:"fire_collection,omitempty" mapstructure:"fire_collection,omitempty" survey:"fire_collection"`
}

// InitSettings returns default variant of settings structure model.
func InitSettings(localPath string) Settings {
	return Settings{
		Name:      DefaultAppName,
		Editor:    DefaultEditor,
		LocalPath: localPath,
	}
}

// ToByte converts settings model to JSON map,
// but returns that value as byte array.
func (s *Settings) ToByte() []byte {
	b, _ := json.Marshal(&s)

	var j map[string]interface{}
	_ = json.Unmarshal(b, &j)

	res, _ := json.Marshal(&j)

	return res
}

// ToJSON converts string structure model to map value.
func (s *Settings) ToJSON() map[string]interface{} {
	b, _ := json.Marshal(&s)

	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)

	return m
}

// FromJSON converts string(map) value to Settings structure.
func DecodeSettings(value string) Settings {
	var m map[string]interface{}
	_ = json.Unmarshal([]byte(value), &m)

	var s Settings
	mapstructure.Decode(m, &s)

	return s
}

// FirePath returns valid firebase collection name.
func (s *Settings) FirePath() string {
	if len(s.FirebaseCollection) > 0 {
		return s.FirebaseCollection
	} else if len(s.Name) > 0 {
		return s.Name
	}

	return DefaultAppName
}

// IsValid checks validness of settings structure.
func (s *Settings) IsValid() bool {
	return len(s.Name) > 0 && len(s.Editor) > 0 && len(s.LocalPath) > 0
}

func (s *Settings) IsFirebaseEnabled() bool {
	return len(s.FirebaseProjectID) > 0 || len(s.FirebaseAccountKey) > 0 || len(s.FirebaseCollection) > 0
}

func IsUpdated(old, current Settings) bool {
	return old.Name != current.Name ||
		old.Editor != current.Editor ||
		old.LocalPath != current.LocalPath ||
		old.FirebaseProjectID != current.FirebaseProjectID ||
		old.FirebaseAccountKey != current.FirebaseAccountKey ||
		old.FirebaseCollection != current.FirebaseCollection
}

// IsPathUpdated checks path difference via based on service type.
func IsPathUpdated(old, current Settings, t string) bool {
	switch t {
	case "LOCAL":
		return old.LocalPath != current.LocalPath
	case "FIREBASE":
		return old.FirebaseCollection != current.FirebaseCollection
	}

	return false
}
