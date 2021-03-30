// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Member is an object representing the database table.
type Member struct {
	ID     int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	Role   int64            `boil:"role" json:"role" toml:"role" yaml:"role"`
	PubKey roomdb.DBFeedRef `boil:"pub_key" json:"pub_key" toml:"pub_key" yaml:"pub_key"`

	R *memberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L memberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MemberColumns = struct {
	ID     string
	Role   string
	PubKey string
}{
	ID:     "id",
	Role:   "role",
	PubKey: "pub_key",
}

// Generated where

var MemberWhere = struct {
	ID     whereHelperint64
	Role   whereHelperint64
	PubKey whereHelperroomdb_DBFeedRef
}{
	ID:     whereHelperint64{field: "\"members\".\"id\""},
	Role:   whereHelperint64{field: "\"members\".\"role\""},
	PubKey: whereHelperroomdb_DBFeedRef{field: "\"members\".\"pub_key\""},
}

// MemberRels is where relationship names are stored.
var MemberRels = struct {
	SIWSSBSessions    string
	Aliases           string
	FallbackPasswords string
	CreatedByInvites  string
}{
	SIWSSBSessions:    "SIWSSBSessions",
	Aliases:           "Aliases",
	FallbackPasswords: "FallbackPasswords",
	CreatedByInvites:  "CreatedByInvites",
}

// memberR is where relationships are stored.
type memberR struct {
	SIWSSBSessions    SIWSSBSessionSlice    `boil:"SIWSSBSessions" json:"SIWSSBSessions" toml:"SIWSSBSessions" yaml:"SIWSSBSessions"`
	Aliases           AliasSlice            `boil:"Aliases" json:"Aliases" toml:"Aliases" yaml:"Aliases"`
	FallbackPasswords FallbackPasswordSlice `boil:"FallbackPasswords" json:"FallbackPasswords" toml:"FallbackPasswords" yaml:"FallbackPasswords"`
	CreatedByInvites  InviteSlice           `boil:"CreatedByInvites" json:"CreatedByInvites" toml:"CreatedByInvites" yaml:"CreatedByInvites"`
}

// NewStruct creates a new relationship struct
func (*memberR) NewStruct() *memberR {
	return &memberR{}
}

// memberL is where Load methods for each relationship are stored.
type memberL struct{}

var (
	memberAllColumns            = []string{"id", "role", "pub_key"}
	memberColumnsWithoutDefault = []string{}
	memberColumnsWithDefault    = []string{"id", "role", "pub_key"}
	memberPrimaryKeyColumns     = []string{"id"}
)

type (
	// MemberSlice is an alias for a slice of pointers to Member.
	// This should generally be used opposed to []Member.
	MemberSlice []*Member
	// MemberHook is the signature for custom Member hook methods
	MemberHook func(context.Context, boil.ContextExecutor, *Member) error

	memberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	memberType                 = reflect.TypeOf(&Member{})
	memberMapping              = queries.MakeStructMapping(memberType)
	memberPrimaryKeyMapping, _ = queries.BindMapping(memberType, memberMapping, memberPrimaryKeyColumns)
	memberInsertCacheMut       sync.RWMutex
	memberInsertCache          = make(map[string]insertCache)
	memberUpdateCacheMut       sync.RWMutex
	memberUpdateCache          = make(map[string]updateCache)
	memberUpsertCacheMut       sync.RWMutex
	memberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var memberBeforeInsertHooks []MemberHook
var memberBeforeUpdateHooks []MemberHook
var memberBeforeDeleteHooks []MemberHook
var memberBeforeUpsertHooks []MemberHook

var memberAfterInsertHooks []MemberHook
var memberAfterSelectHooks []MemberHook
var memberAfterUpdateHooks []MemberHook
var memberAfterDeleteHooks []MemberHook
var memberAfterUpsertHooks []MemberHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Member) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Member) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Member) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Member) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Member) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Member) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Member) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Member) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Member) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range memberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMemberHook registers your hook function for all future operations.
func AddMemberHook(hookPoint boil.HookPoint, memberHook MemberHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		memberBeforeInsertHooks = append(memberBeforeInsertHooks, memberHook)
	case boil.BeforeUpdateHook:
		memberBeforeUpdateHooks = append(memberBeforeUpdateHooks, memberHook)
	case boil.BeforeDeleteHook:
		memberBeforeDeleteHooks = append(memberBeforeDeleteHooks, memberHook)
	case boil.BeforeUpsertHook:
		memberBeforeUpsertHooks = append(memberBeforeUpsertHooks, memberHook)
	case boil.AfterInsertHook:
		memberAfterInsertHooks = append(memberAfterInsertHooks, memberHook)
	case boil.AfterSelectHook:
		memberAfterSelectHooks = append(memberAfterSelectHooks, memberHook)
	case boil.AfterUpdateHook:
		memberAfterUpdateHooks = append(memberAfterUpdateHooks, memberHook)
	case boil.AfterDeleteHook:
		memberAfterDeleteHooks = append(memberAfterDeleteHooks, memberHook)
	case boil.AfterUpsertHook:
		memberAfterUpsertHooks = append(memberAfterUpsertHooks, memberHook)
	}
}

// One returns a single member record from the query.
func (q memberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Member, error) {
	o := &Member{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for members")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Member records from the query.
func (q memberQuery) All(ctx context.Context, exec boil.ContextExecutor) (MemberSlice, error) {
	var o []*Member

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Member slice")
	}

	if len(memberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Member records in the query.
func (q memberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count members rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q memberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if members exists")
	}

	return count > 0, nil
}

// SIWSSBSessions retrieves all the SIWSSB_session's SIWSSBSessions with an executor.
func (o *Member) SIWSSBSessions(mods ...qm.QueryMod) sIWSSBSessionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"SIWSSB_sessions\".\"member_id\"=?", o.ID),
	)

	query := SIWSSBSessions(queryMods...)
	queries.SetFrom(query.Query, "\"SIWSSB_sessions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"SIWSSB_sessions\".*"})
	}

	return query
}

// Aliases retrieves all the alias's Aliases with an executor.
func (o *Member) Aliases(mods ...qm.QueryMod) aliasQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"aliases\".\"member_id\"=?", o.ID),
	)

	query := Aliases(queryMods...)
	queries.SetFrom(query.Query, "\"aliases\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"aliases\".*"})
	}

	return query
}

// FallbackPasswords retrieves all the fallback_password's FallbackPasswords with an executor.
func (o *Member) FallbackPasswords(mods ...qm.QueryMod) fallbackPasswordQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"fallback_passwords\".\"member_id\"=?", o.ID),
	)

	query := FallbackPasswords(queryMods...)
	queries.SetFrom(query.Query, "\"fallback_passwords\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"fallback_passwords\".*"})
	}

	return query
}

// CreatedByInvites retrieves all the invite's Invites with an executor via created_by column.
func (o *Member) CreatedByInvites(mods ...qm.QueryMod) inviteQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"invites\".\"created_by\"=?", o.ID),
	)

	query := Invites(queryMods...)
	queries.SetFrom(query.Query, "\"invites\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"invites\".*"})
	}

	return query
}

// LoadSIWSSBSessions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (memberL) LoadSIWSSBSessions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMember interface{}, mods queries.Applicator) error {
	var slice []*Member
	var object *Member

	if singular {
		object = maybeMember.(*Member)
	} else {
		slice = *maybeMember.(*[]*Member)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &memberR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &memberR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`SIWSSB_sessions`),
		qm.WhereIn(`SIWSSB_sessions.member_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SIWSSB_sessions")
	}

	var resultSlice []*SIWSSBSession
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SIWSSB_sessions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on SIWSSB_sessions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for SIWSSB_sessions")
	}

	if len(sIWSSBSessionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SIWSSBSessions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sIWSSBSessionR{}
			}
			foreign.R.Member = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MemberID {
				local.R.SIWSSBSessions = append(local.R.SIWSSBSessions, foreign)
				if foreign.R == nil {
					foreign.R = &sIWSSBSessionR{}
				}
				foreign.R.Member = local
				break
			}
		}
	}

	return nil
}

// LoadAliases allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (memberL) LoadAliases(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMember interface{}, mods queries.Applicator) error {
	var slice []*Member
	var object *Member

	if singular {
		object = maybeMember.(*Member)
	} else {
		slice = *maybeMember.(*[]*Member)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &memberR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &memberR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`aliases`),
		qm.WhereIn(`aliases.member_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load aliases")
	}

	var resultSlice []*Alias
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice aliases")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on aliases")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for aliases")
	}

	if len(aliasAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Aliases = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &aliasR{}
			}
			foreign.R.Member = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MemberID {
				local.R.Aliases = append(local.R.Aliases, foreign)
				if foreign.R == nil {
					foreign.R = &aliasR{}
				}
				foreign.R.Member = local
				break
			}
		}
	}

	return nil
}

// LoadFallbackPasswords allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (memberL) LoadFallbackPasswords(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMember interface{}, mods queries.Applicator) error {
	var slice []*Member
	var object *Member

	if singular {
		object = maybeMember.(*Member)
	} else {
		slice = *maybeMember.(*[]*Member)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &memberR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &memberR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`fallback_passwords`),
		qm.WhereIn(`fallback_passwords.member_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load fallback_passwords")
	}

	var resultSlice []*FallbackPassword
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice fallback_passwords")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on fallback_passwords")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for fallback_passwords")
	}

	if len(fallbackPasswordAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FallbackPasswords = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &fallbackPasswordR{}
			}
			foreign.R.Member = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MemberID {
				local.R.FallbackPasswords = append(local.R.FallbackPasswords, foreign)
				if foreign.R == nil {
					foreign.R = &fallbackPasswordR{}
				}
				foreign.R.Member = local
				break
			}
		}
	}

	return nil
}

// LoadCreatedByInvites allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (memberL) LoadCreatedByInvites(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMember interface{}, mods queries.Applicator) error {
	var slice []*Member
	var object *Member

	if singular {
		object = maybeMember.(*Member)
	} else {
		slice = *maybeMember.(*[]*Member)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &memberR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &memberR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`invites`),
		qm.WhereIn(`invites.created_by in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load invites")
	}

	var resultSlice []*Invite
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice invites")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on invites")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for invites")
	}

	if len(inviteAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CreatedByInvites = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &inviteR{}
			}
			foreign.R.CreatedByMember = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CreatedBy {
				local.R.CreatedByInvites = append(local.R.CreatedByInvites, foreign)
				if foreign.R == nil {
					foreign.R = &inviteR{}
				}
				foreign.R.CreatedByMember = local
				break
			}
		}
	}

	return nil
}

// AddSIWSSBSessions adds the given related objects to the existing relationships
// of the member, optionally inserting them as new records.
// Appends related to o.R.SIWSSBSessions.
// Sets related.R.Member appropriately.
func (o *Member) AddSIWSSBSessions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SIWSSBSession) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MemberID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"SIWSSB_sessions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"member_id"}),
				strmangle.WhereClause("\"", "\"", 0, sIWSSBSessionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MemberID = o.ID
		}
	}

	if o.R == nil {
		o.R = &memberR{
			SIWSSBSessions: related,
		}
	} else {
		o.R.SIWSSBSessions = append(o.R.SIWSSBSessions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sIWSSBSessionR{
				Member: o,
			}
		} else {
			rel.R.Member = o
		}
	}
	return nil
}

// AddAliases adds the given related objects to the existing relationships
// of the member, optionally inserting them as new records.
// Appends related to o.R.Aliases.
// Sets related.R.Member appropriately.
func (o *Member) AddAliases(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Alias) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MemberID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"aliases\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"member_id"}),
				strmangle.WhereClause("\"", "\"", 0, aliasPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MemberID = o.ID
		}
	}

	if o.R == nil {
		o.R = &memberR{
			Aliases: related,
		}
	} else {
		o.R.Aliases = append(o.R.Aliases, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &aliasR{
				Member: o,
			}
		} else {
			rel.R.Member = o
		}
	}
	return nil
}

// AddFallbackPasswords adds the given related objects to the existing relationships
// of the member, optionally inserting them as new records.
// Appends related to o.R.FallbackPasswords.
// Sets related.R.Member appropriately.
func (o *Member) AddFallbackPasswords(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*FallbackPassword) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MemberID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"fallback_passwords\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"member_id"}),
				strmangle.WhereClause("\"", "\"", 0, fallbackPasswordPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MemberID = o.ID
		}
	}

	if o.R == nil {
		o.R = &memberR{
			FallbackPasswords: related,
		}
	} else {
		o.R.FallbackPasswords = append(o.R.FallbackPasswords, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &fallbackPasswordR{
				Member: o,
			}
		} else {
			rel.R.Member = o
		}
	}
	return nil
}

// AddCreatedByInvites adds the given related objects to the existing relationships
// of the member, optionally inserting them as new records.
// Appends related to o.R.CreatedByInvites.
// Sets related.R.CreatedByMember appropriately.
func (o *Member) AddCreatedByInvites(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Invite) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CreatedBy = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"invites\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"created_by"}),
				strmangle.WhereClause("\"", "\"", 0, invitePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CreatedBy = o.ID
		}
	}

	if o.R == nil {
		o.R = &memberR{
			CreatedByInvites: related,
		}
	} else {
		o.R.CreatedByInvites = append(o.R.CreatedByInvites, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &inviteR{
				CreatedByMember: o,
			}
		} else {
			rel.R.CreatedByMember = o
		}
	}
	return nil
}

// Members retrieves all the records using an executor.
func Members(mods ...qm.QueryMod) memberQuery {
	mods = append(mods, qm.From("\"members\""))
	return memberQuery{NewQuery(mods...)}
}

// FindMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMember(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Member, error) {
	memberObj := &Member{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"members\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, memberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from members")
	}

	return memberObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Member) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no members provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(memberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	memberInsertCacheMut.RLock()
	cache, cached := memberInsertCache[key]
	memberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			memberAllColumns,
			memberColumnsWithDefault,
			memberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(memberType, memberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(memberType, memberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"members\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"members\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"members\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, memberPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into members")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == memberMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for members")
	}

CacheNoHooks:
	if !cached {
		memberInsertCacheMut.Lock()
		memberInsertCache[key] = cache
		memberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Member.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Member) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	memberUpdateCacheMut.RLock()
	cache, cached := memberUpdateCache[key]
	memberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			memberAllColumns,
			memberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update members, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"members\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, memberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(memberType, memberMapping, append(wl, memberPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update members row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for members")
	}

	if !cached {
		memberUpdateCacheMut.Lock()
		memberUpdateCache[key] = cache
		memberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q memberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for members")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MemberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in member slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all member")
	}
	return rowsAff, nil
}

// Delete deletes a single Member record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Member) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Member provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), memberPrimaryKeyMapping)
	sql := "DELETE FROM \"members\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for members")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q memberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no memberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for members")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MemberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(memberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"members\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from member slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for members")
	}

	if len(memberAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Member) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMember(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MemberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), memberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"members\".* FROM \"members\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, memberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MemberSlice")
	}

	*o = slice

	return nil
}

// MemberExists checks if the Member row exists.
func MemberExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"members\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if members exists")
	}

	return exists, nil
}
