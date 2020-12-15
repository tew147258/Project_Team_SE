// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/tanapon395/playlist-video/ent/playlist"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/resolution"
	"github.com/tanapon395/playlist-video/ent/video"
)

// Playlist_Video is the model entity for the Playlist_Video schema.
type Playlist_Video struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AddedTime holds the value of the "added_time" field.
	AddedTime time.Time `json:"added_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Playlist_VideoQuery when eager-loading is set.
	Edges         Playlist_VideoEdges `json:"edges"`
	playlist_id   *int
	resolution_id *int
	video_id      *int
}

// Playlist_VideoEdges holds the relations/edges for other nodes in the graph.
type Playlist_VideoEdges struct {
	// Playlist holds the value of the playlist edge.
	Playlist *Playlist
	// Video holds the value of the video edge.
	Video *Video
	// Resolution holds the value of the resolution edge.
	Resolution *Resolution
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PlaylistOrErr returns the Playlist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Playlist_VideoEdges) PlaylistOrErr() (*Playlist, error) {
	if e.loadedTypes[0] {
		if e.Playlist == nil {
			// The edge playlist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: playlist.Label}
		}
		return e.Playlist, nil
	}
	return nil, &NotLoadedError{edge: "playlist"}
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Playlist_VideoEdges) VideoOrErr() (*Video, error) {
	if e.loadedTypes[1] {
		if e.Video == nil {
			// The edge video was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: video.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// ResolutionOrErr returns the Resolution value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Playlist_VideoEdges) ResolutionOrErr() (*Resolution, error) {
	if e.loadedTypes[2] {
		if e.Resolution == nil {
			// The edge resolution was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resolution.Label}
		}
		return e.Resolution, nil
	}
	return nil, &NotLoadedError{edge: "resolution"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Playlist_Video) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // added_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Playlist_Video) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // playlist_id
		&sql.NullInt64{}, // resolution_id
		&sql.NullInt64{}, // video_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Playlist_Video fields.
func (pv *Playlist_Video) assignValues(values ...interface{}) error {
	if m, n := len(values), len(playlist_video.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pv.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field added_time", values[0])
	} else if value.Valid {
		pv.AddedTime = value.Time
	}
	values = values[1:]
	if len(values) == len(playlist_video.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field playlist_id", value)
		} else if value.Valid {
			pv.playlist_id = new(int)
			*pv.playlist_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field resolution_id", value)
		} else if value.Valid {
			pv.resolution_id = new(int)
			*pv.resolution_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field video_id", value)
		} else if value.Valid {
			pv.video_id = new(int)
			*pv.video_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPlaylist queries the playlist edge of the Playlist_Video.
func (pv *Playlist_Video) QueryPlaylist() *PlaylistQuery {
	return (&Playlist_VideoClient{config: pv.config}).QueryPlaylist(pv)
}

// QueryVideo queries the video edge of the Playlist_Video.
func (pv *Playlist_Video) QueryVideo() *VideoQuery {
	return (&Playlist_VideoClient{config: pv.config}).QueryVideo(pv)
}

// QueryResolution queries the resolution edge of the Playlist_Video.
func (pv *Playlist_Video) QueryResolution() *ResolutionQuery {
	return (&Playlist_VideoClient{config: pv.config}).QueryResolution(pv)
}

// Update returns a builder for updating this Playlist_Video.
// Note that, you need to call Playlist_Video.Unwrap() before calling this method, if this Playlist_Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *Playlist_Video) Update() *PlaylistVideoUpdateOne {
	return (&Playlist_VideoClient{config: pv.config}).UpdateOne(pv)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pv *Playlist_Video) Unwrap() *Playlist_Video {
	tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: Playlist_Video is not a transactional entity")
	}
	pv.config.driver = tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *Playlist_Video) String() string {
	var builder strings.Builder
	builder.WriteString("Playlist_Video(")
	builder.WriteString(fmt.Sprintf("id=%v", pv.ID))
	builder.WriteString(", added_time=")
	builder.WriteString(pv.AddedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Playlist_Videos is a parsable slice of Playlist_Video.
type Playlist_Videos []*Playlist_Video

func (pv Playlist_Videos) config(cfg config) {
	for _i := range pv {
		pv[_i].config = cfg
	}
}
