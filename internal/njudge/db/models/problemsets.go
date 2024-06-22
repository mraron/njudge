// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Problemset is an object representing the database table.
type Problemset struct {
	Name           string `boil:"name" json:"name" toml:"name" yaml:"name"`
	CodeVisibility string `boil:"code_visibility" json:"code_visibility" toml:"code_visibility" yaml:"code_visibility"`

	R *problemsetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L problemsetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProblemsetColumns = struct {
	Name           string
	CodeVisibility string
}{
	Name:           "name",
	CodeVisibility: "code_visibility",
}

var ProblemsetTableColumns = struct {
	Name           string
	CodeVisibility string
}{
	Name:           "problemsets.name",
	CodeVisibility: "problemsets.code_visibility",
}

// Generated where

var ProblemsetWhere = struct {
	Name           whereHelperstring
	CodeVisibility whereHelperstring
}{
	Name:           whereHelperstring{field: "\"problemsets\".\"name\""},
	CodeVisibility: whereHelperstring{field: "\"problemsets\".\"code_visibility\""},
}

// ProblemsetRels is where relationship names are stored.
var ProblemsetRels = struct {
	ProblemRels string
}{
	ProblemRels: "ProblemRels",
}

// problemsetR is where relationships are stored.
type problemsetR struct {
	ProblemRels ProblemRelSlice `boil:"ProblemRels" json:"ProblemRels" toml:"ProblemRels" yaml:"ProblemRels"`
}

// NewStruct creates a new relationship struct
func (*problemsetR) NewStruct() *problemsetR {
	return &problemsetR{}
}

func (r *problemsetR) GetProblemRels() ProblemRelSlice {
	if r == nil {
		return nil
	}
	return r.ProblemRels
}

// problemsetL is where Load methods for each relationship are stored.
type problemsetL struct{}

var (
	problemsetAllColumns            = []string{"name", "code_visibility"}
	problemsetColumnsWithoutDefault = []string{"name", "code_visibility"}
	problemsetColumnsWithDefault    = []string{}
	problemsetPrimaryKeyColumns     = []string{"name"}
	problemsetGeneratedColumns      = []string{}
)

type (
	// ProblemsetSlice is an alias for a slice of pointers to Problemset.
	// This should almost always be used instead of []Problemset.
	ProblemsetSlice []*Problemset
	// ProblemsetHook is the signature for custom Problemset hook methods
	ProblemsetHook func(context.Context, boil.ContextExecutor, *Problemset) error

	problemsetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	problemsetType                 = reflect.TypeOf(&Problemset{})
	problemsetMapping              = queries.MakeStructMapping(problemsetType)
	problemsetPrimaryKeyMapping, _ = queries.BindMapping(problemsetType, problemsetMapping, problemsetPrimaryKeyColumns)
	problemsetInsertCacheMut       sync.RWMutex
	problemsetInsertCache          = make(map[string]insertCache)
	problemsetUpdateCacheMut       sync.RWMutex
	problemsetUpdateCache          = make(map[string]updateCache)
	problemsetUpsertCacheMut       sync.RWMutex
	problemsetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var problemsetAfterSelectMu sync.Mutex
var problemsetAfterSelectHooks []ProblemsetHook

var problemsetBeforeInsertMu sync.Mutex
var problemsetBeforeInsertHooks []ProblemsetHook
var problemsetAfterInsertMu sync.Mutex
var problemsetAfterInsertHooks []ProblemsetHook

var problemsetBeforeUpdateMu sync.Mutex
var problemsetBeforeUpdateHooks []ProblemsetHook
var problemsetAfterUpdateMu sync.Mutex
var problemsetAfterUpdateHooks []ProblemsetHook

var problemsetBeforeDeleteMu sync.Mutex
var problemsetBeforeDeleteHooks []ProblemsetHook
var problemsetAfterDeleteMu sync.Mutex
var problemsetAfterDeleteHooks []ProblemsetHook

var problemsetBeforeUpsertMu sync.Mutex
var problemsetBeforeUpsertHooks []ProblemsetHook
var problemsetAfterUpsertMu sync.Mutex
var problemsetAfterUpsertHooks []ProblemsetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Problemset) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Problemset) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Problemset) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Problemset) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Problemset) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Problemset) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Problemset) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Problemset) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Problemset) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range problemsetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProblemsetHook registers your hook function for all future operations.
func AddProblemsetHook(hookPoint boil.HookPoint, problemsetHook ProblemsetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		problemsetAfterSelectMu.Lock()
		problemsetAfterSelectHooks = append(problemsetAfterSelectHooks, problemsetHook)
		problemsetAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		problemsetBeforeInsertMu.Lock()
		problemsetBeforeInsertHooks = append(problemsetBeforeInsertHooks, problemsetHook)
		problemsetBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		problemsetAfterInsertMu.Lock()
		problemsetAfterInsertHooks = append(problemsetAfterInsertHooks, problemsetHook)
		problemsetAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		problemsetBeforeUpdateMu.Lock()
		problemsetBeforeUpdateHooks = append(problemsetBeforeUpdateHooks, problemsetHook)
		problemsetBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		problemsetAfterUpdateMu.Lock()
		problemsetAfterUpdateHooks = append(problemsetAfterUpdateHooks, problemsetHook)
		problemsetAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		problemsetBeforeDeleteMu.Lock()
		problemsetBeforeDeleteHooks = append(problemsetBeforeDeleteHooks, problemsetHook)
		problemsetBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		problemsetAfterDeleteMu.Lock()
		problemsetAfterDeleteHooks = append(problemsetAfterDeleteHooks, problemsetHook)
		problemsetAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		problemsetBeforeUpsertMu.Lock()
		problemsetBeforeUpsertHooks = append(problemsetBeforeUpsertHooks, problemsetHook)
		problemsetBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		problemsetAfterUpsertMu.Lock()
		problemsetAfterUpsertHooks = append(problemsetAfterUpsertHooks, problemsetHook)
		problemsetAfterUpsertMu.Unlock()
	}
}

// OneG returns a single problemset record from the query using the global executor.
func (q problemsetQuery) OneG(ctx context.Context) (*Problemset, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single problemset record from the query.
func (q problemsetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Problemset, error) {
	o := &Problemset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for problemsets")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Problemset records from the query using the global executor.
func (q problemsetQuery) AllG(ctx context.Context) (ProblemsetSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Problemset records from the query.
func (q problemsetQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProblemsetSlice, error) {
	var o []*Problemset

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Problemset slice")
	}

	if len(problemsetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Problemset records in the query using the global executor
func (q problemsetQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Problemset records in the query.
func (q problemsetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count problemsets rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q problemsetQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q problemsetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if problemsets exists")
	}

	return count > 0, nil
}

// ProblemRels retrieves all the problem_rel's ProblemRels with an executor.
func (o *Problemset) ProblemRels(mods ...qm.QueryMod) problemRelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"problem_rels\".\"problemset\"=?", o.Name),
	)

	return ProblemRels(queryMods...)
}

// LoadProblemRels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (problemsetL) LoadProblemRels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProblemset interface{}, mods queries.Applicator) error {
	var slice []*Problemset
	var object *Problemset

	if singular {
		var ok bool
		object, ok = maybeProblemset.(*Problemset)
		if !ok {
			object = new(Problemset)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProblemset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProblemset))
			}
		}
	} else {
		s, ok := maybeProblemset.(*[]*Problemset)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProblemset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProblemset))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &problemsetR{}
		}
		args[object.Name] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &problemsetR{}
			}
			args[obj.Name] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`problem_rels`),
		qm.WhereIn(`problem_rels.problemset in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load problem_rels")
	}

	var resultSlice []*ProblemRel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice problem_rels")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on problem_rels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for problem_rels")
	}

	if len(problemRelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProblemRels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &problemRelR{}
			}
			foreign.R.ProblemRelProblemset = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Name == foreign.Problemset {
				local.R.ProblemRels = append(local.R.ProblemRels, foreign)
				if foreign.R == nil {
					foreign.R = &problemRelR{}
				}
				foreign.R.ProblemRelProblemset = local
				break
			}
		}
	}

	return nil
}

// AddProblemRelsG adds the given related objects to the existing relationships
// of the problemset, optionally inserting them as new records.
// Appends related to o.R.ProblemRels.
// Sets related.R.ProblemRelProblemset appropriately.
// Uses the global database handle.
func (o *Problemset) AddProblemRelsG(ctx context.Context, insert bool, related ...*ProblemRel) error {
	return o.AddProblemRels(ctx, boil.GetContextDB(), insert, related...)
}

// AddProblemRels adds the given related objects to the existing relationships
// of the problemset, optionally inserting them as new records.
// Appends related to o.R.ProblemRels.
// Sets related.R.ProblemRelProblemset appropriately.
func (o *Problemset) AddProblemRels(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ProblemRel) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Problemset = o.Name
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"problem_rels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"problemset"}),
				strmangle.WhereClause("\"", "\"", 2, problemRelPrimaryKeyColumns),
			)
			values := []interface{}{o.Name, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Problemset = o.Name
		}
	}

	if o.R == nil {
		o.R = &problemsetR{
			ProblemRels: related,
		}
	} else {
		o.R.ProblemRels = append(o.R.ProblemRels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &problemRelR{
				ProblemRelProblemset: o,
			}
		} else {
			rel.R.ProblemRelProblemset = o
		}
	}
	return nil
}

// Problemsets retrieves all the records using an executor.
func Problemsets(mods ...qm.QueryMod) problemsetQuery {
	mods = append(mods, qm.From("\"problemsets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"problemsets\".*"})
	}

	return problemsetQuery{q}
}

// FindProblemsetG retrieves a single record by ID.
func FindProblemsetG(ctx context.Context, name string, selectCols ...string) (*Problemset, error) {
	return FindProblemset(ctx, boil.GetContextDB(), name, selectCols...)
}

// FindProblemset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProblemset(ctx context.Context, exec boil.ContextExecutor, name string, selectCols ...string) (*Problemset, error) {
	problemsetObj := &Problemset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"problemsets\" where \"name\"=$1", sel,
	)

	q := queries.Raw(query, name)

	err := q.Bind(ctx, exec, problemsetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from problemsets")
	}

	if err = problemsetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return problemsetObj, err
	}

	return problemsetObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Problemset) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Problemset) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no problemsets provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemsetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	problemsetInsertCacheMut.RLock()
	cache, cached := problemsetInsertCache[key]
	problemsetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			problemsetAllColumns,
			problemsetColumnsWithDefault,
			problemsetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(problemsetType, problemsetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(problemsetType, problemsetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"problemsets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"problemsets\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into problemsets")
	}

	if !cached {
		problemsetInsertCacheMut.Lock()
		problemsetInsertCache[key] = cache
		problemsetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Problemset record using the global executor.
// See Update for more documentation.
func (o *Problemset) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Problemset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Problemset) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	problemsetUpdateCacheMut.RLock()
	cache, cached := problemsetUpdateCache[key]
	problemsetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			problemsetAllColumns,
			problemsetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update problemsets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"problemsets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, problemsetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(problemsetType, problemsetMapping, append(wl, problemsetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update problemsets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for problemsets")
	}

	if !cached {
		problemsetUpdateCacheMut.Lock()
		problemsetUpdateCache[key] = cache
		problemsetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q problemsetQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q problemsetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for problemsets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for problemsets")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProblemsetSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProblemsetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"problemsets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, problemsetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in problemset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all problemset")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Problemset) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Problemset) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no problemsets provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemsetColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	problemsetUpsertCacheMut.RLock()
	cache, cached := problemsetUpsertCache[key]
	problemsetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			problemsetAllColumns,
			problemsetColumnsWithDefault,
			problemsetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			problemsetAllColumns,
			problemsetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert problemsets, could not build update column list")
		}

		ret := strmangle.SetComplement(problemsetAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(problemsetPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert problemsets, could not build conflict column list")
			}

			conflict = make([]string, len(problemsetPrimaryKeyColumns))
			copy(conflict, problemsetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"problemsets\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(problemsetType, problemsetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(problemsetType, problemsetMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert problemsets")
	}

	if !cached {
		problemsetUpsertCacheMut.Lock()
		problemsetUpsertCache[key] = cache
		problemsetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Problemset record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Problemset) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Problemset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Problemset) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Problemset provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), problemsetPrimaryKeyMapping)
	sql := "DELETE FROM \"problemsets\" WHERE \"name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from problemsets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for problemsets")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q problemsetQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q problemsetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no problemsetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problemsets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problemsets")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProblemsetSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProblemsetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(problemsetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"problemsets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemsetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problemset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problemsets")
	}

	if len(problemsetAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Problemset) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Problemset provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Problemset) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProblemset(ctx, exec, o.Name)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemsetSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ProblemsetSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemsetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProblemsetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemsetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"problemsets\".* FROM \"problemsets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemsetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProblemsetSlice")
	}

	*o = slice

	return nil
}

// ProblemsetExistsG checks if the Problemset row exists.
func ProblemsetExistsG(ctx context.Context, name string) (bool, error) {
	return ProblemsetExists(ctx, boil.GetContextDB(), name)
}

// ProblemsetExists checks if the Problemset row exists.
func ProblemsetExists(ctx context.Context, exec boil.ContextExecutor, name string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"problemsets\" where \"name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, name)
	}
	row := exec.QueryRowContext(ctx, sql, name)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if problemsets exists")
	}

	return exists, nil
}

// Exists checks if the Problemset row exists.
func (o *Problemset) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProblemsetExists(ctx, exec, o.Name)
}
