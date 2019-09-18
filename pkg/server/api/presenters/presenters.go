/* Copyright (C) 2019 Monomax Software Pty Ltd
 *
 * This file is part of Dnote.
 *
 * Dnote is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Dnote is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Dnote.  If not, see <https://www.gnu.org/licenses/>.
 */

package presenters

import (
	"time"

	"github.com/dnote/dnote/pkg/server/database"
)

// FormatTS rounds up the given timestamp to the microsecond
// so as to make the times in the responses consistent
func FormatTS(ts time.Time) time.Time {
	return ts.UTC().Round(time.Microsecond)
}

// Book is a result of PresentBooks
type Book struct {
	UUID      string    `json:"uuid"`
	USN       int       `json:"usn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
}

// PresentBook presents a book
func PresentBook(book database.Book) Book {
	return Book{
		UUID:      book.UUID,
		USN:       book.USN,
		CreatedAt: FormatTS(book.CreatedAt),
		UpdatedAt: FormatTS(book.UpdatedAt),
		Label:     book.Label,
	}
}

// PresentBooks presents books
func PresentBooks(books []database.Book) []Book {
	ret := []Book{}

	for _, book := range books {
		p := PresentBook(book)
		ret = append(ret, p)
	}

	return ret
}

// Note is a result of PresentNote
type Note struct {
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"content"`
	AddedOn   int64     `json:"added_on"`
	Public    bool      `json:"public"`
	USN       int       `json:"usn"`
	Book      NoteBook  `json:"book"`
	User      NoteUser  `json:"user"`
}

// NoteBook is a nested book for PresentNotesResult
type NoteBook struct {
	UUID  string `json:"uuid"`
	Label string `json:"label"`
}

// NoteUser is a nested book for PresentNotesResult
type NoteUser struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// PresentNote presents note
func PresentNote(note database.Note) Note {
	ret := Note{
		UUID:      note.UUID,
		CreatedAt: FormatTS(note.CreatedAt),
		UpdatedAt: FormatTS(note.UpdatedAt),
		Body:      note.Body,
		AddedOn:   note.AddedOn,
		Public:    note.Public,
		USN:       note.USN,
		Book: NoteBook{
			UUID:  note.Book.UUID,
			Label: note.Book.Label,
		},
		User: NoteUser{
			Name: note.User.Name,
			UUID: note.User.UUID,
		},
	}

	return ret
}

// PresentNotes presents notes
func PresentNotes(notes []database.Note) []Note {
	ret := []Note{}

	for _, note := range notes {
		p := PresentNote(note)
		ret = append(ret, p)
	}

	return ret
}

// Digest is a presented digest
type Digest struct {
	UUID      string    `json:"uuid"`
	Notes     []Note    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PresentDigests presetns digests
func PresentDigests(digests []database.Digest) []Digest {
	ret := []Digest{}

	for _, digest := range digests {
		p := Digest{
			UUID:      digest.UUID,
			CreatedAt: digest.CreatedAt,
			UpdatedAt: digest.UpdatedAt,
		}

		ret = append(ret, p)
	}

	return ret
}

// PresentDigest presents a digest
func PresentDigest(digest database.Digest) Digest {
	ret := Digest{
		UUID:  digest.UUID,
		Notes: PresentNotes(digest.Notes),
	}

	return ret
}

// EmailPreference is a presented email digest
type EmailPreference struct {
	DigestWeekly bool      `json:"digest_weekly"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PresentEmailPreference presents a digest
func PresentEmailPreference(p database.EmailPreference) EmailPreference {
	ret := EmailPreference{
		DigestWeekly: p.DigestWeekly,
		CreatedAt:    FormatTS(p.CreatedAt),
		UpdatedAt:    FormatTS(p.UpdatedAt),
	}

	return ret
}
