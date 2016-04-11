// Copyright 2016 Danko Miocevic. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Danko Miocevic

package store

import (
	"bazil.org/fuse"
	"github.com/boltdb/bolt"
)

// GetPlaylistFilePath function should return the path for a specific
// file in a specific playlist.
// The file could be on two places, first option is that the file is
// stored in the database. In that case, the file will be stored somewhere
// else in the MuLi filesystem but that will be specified on the
// item in the database.
// On the other hand, the file could be just dropped inside the playlist
// and it will be temporary stored in a directory inside the playlists
// directory.
// The playlist name is specified on the first argument and the song
// name on the second.
// The mount path is also needed and should be specified on the third
// argument.
// This function returns a string containing the file path and an error
// that will be nil if everything is ok.
func GetPlaylistFilePath(playlist, song, mPoint string) (string, error) {
	return "", nil
}

// ListPlaylists function returns all the names of the playlists available
// in the MuLi system.
// It receives no arguments and returns a slice of Dir objects to list
// all the available playlists and the error if there is any.
func ListPlaylists() ([]fuse.Dirent, error) {
	db, err := bolt.Open(config.DbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var a []fuse.Dirent
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Playlists"))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if v == nil {
				var node fuse.Dirent
				node.Name = string(k)
				node.Type = fuse.DT_Dir
				a = append(a, node)
			}
		}
		return nil
	})
	return a, nil
}

// ListPlaylistSongs function returns all the songs inside a playlist.
// The available songs are loaded from the database and also from the
// temporary drop directory named after the playlist.
// It receives a playlist name and returns a slice with all the
// files.
func ListPlaylistSongs(playlist string) ([]fuse.Dirent, error) {
	db, err := bolt.Open(config.DbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var a []fuse.Dirent
	err = db.View(func(tx *bolt.Tx) error {
		root := tx.Bucket([]byte("Playlists"))
		if root == nil {
			return nil
		}

		b := root.Bucket([]byte(playlist))
		if b == nil {
			return nil
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if v != nil {
				var node fuse.Dirent
				node.Name = string(k)
				node.Type = fuse.DT_File
				a = append(a, node)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return a, nil

	return nil, nil
}

// CreatePlaylist function creates a playlist item in the database and
// also creates it in the filesystem.
// It receives the playlist name and returns the modified name and an
// error if something went wrong.
func CreatePlaylist(name string) (string, error) {
	return "", nil
}
