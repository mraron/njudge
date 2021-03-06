// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProblemRel is an object representing the database table.
type ProblemRel struct {
	Problemset string   `boil:"problemset" json:"problemset" toml:"problemset" yaml:"problemset"`
	Problem    string   `boil:"problem" json:"problem" toml:"problem" yaml:"problem"`
	ID         int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CategoryID null.Int `boil:"category_id" json:"category_id,omitempty" toml:"category_id" yaml:"category_id,omitempty"`

	R *problemRelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L problemRelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProblemRelColumns = struct {
	Problemset string
	Problem    string
	ID         string
	CategoryID string
}{
	Problemset: "problemset",
	Problem:    "problem",
	ID:         "id",
	CategoryID: "category_id",
}

// Generated where

var ProblemRelWhere = struct {
	Problemset whereHelperstring
	Problem    whereHelperstring
	ID         whereHelperint
	CategoryID whereHelpernull_Int
}{
	Problemset: whereHelperstring{field: "\"problem_rels\".\"problemset\""},
	Problem:    whereHelperstring{field: "\"problem_rels\".\"problem\""},
	ID:         whereHelperint{field: "\"problem_rels\".\"id\""},
	CategoryID: whereHelpernull_Int{field: "\"problem_rels\".\"category_id\""},
}

// ProblemRelRels is where relationship names are stored.
var ProblemRelRels = struct {
	Category string
}{
	Category: "Category",
}

// problemRelR is where relationships are stored.
type problemRelR struct {
	Category *ProblemCategory `boil:"Category" json:"Category" toml:"Category" yaml:"Category"`
}

// NewStruct creates a new relationship struct
func (*problemRelR) NewStruct() *problemRelR {
	return &problemRelR{}
}

// problemRelL is where Load methods for each relationship are stored.
type problemRelL struct{}

var (
	problemRelAllColumns            = []string{"problemset", "problem", "id", "category_id"}
	problemRelColumnsWithoutDefault = []string{"problemset", "problem", "category_id"}
	problemRelColumnsWithDefault    = []string{"id"}
	problemRelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProblemRelSlice is an alias for a slice of pointers to ProblemRel.
	// This should generally be used opposed to []ProblemRel.
	ProblemRelSlice []*ProblemRel
	// ProblemRelHook is the signature for custom ProblemRel hook methods
	ProblemRelHook func(boil.Executor, *ProblemRel) error

	problemRelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	problemRelType                 = reflect.TypeOf(&ProblemRel{})
	problemRelMapping              = queries.MakeStructMapping(problemRelType)
	problemRelPrimaryKeyMapping, _ = queries.BindMapping(problemRelType, problemRelMapping, problemRelPrimaryKeyColumns)
	problemRelInsertCacheMut       sync.RWMutex
	problemRelInsertCache          = make(map[string]insertCache)
	problemRelUpdateCacheMut       sync.RWMutex
	problemRelUpdateCache          = make(map[string]updateCache)
	problemRelUpsertCacheMut       sync.RWMutex
	problemRelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var problemRelBeforeInsertHooks []ProblemRelHook
var problemRelBeforeUpdateHooks []ProblemRelHook
var problemRelBeforeDeleteHooks []ProblemRelHook
var problemRelBeforeUpsertHooks []ProblemRelHook

var problemRelAfterInsertHooks []ProblemRelHook
var problemRelAfterSelectHooks []ProblemRelHook
var problemRelAfterUpdateHooks []ProblemRelHook
var problemRelAfterDeleteHooks []ProblemRelHook
var problemRelAfterUpsertHooks []ProblemRelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProblemRel) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProblemRel) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProblemRel) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProblemRel) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProblemRel) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProblemRel) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProblemRel) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProblemRel) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProblemRel) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range problemRelAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProblemRelHook registers your hook function for all future operations.
func AddProblemRelHook(hookPoint boil.HookPoint, problemRelHook ProblemRelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		problemRelBeforeInsertHooks = append(problemRelBeforeInsertHooks, problemRelHook)
	case boil.BeforeUpdateHook:
		problemRelBeforeUpdateHooks = append(problemRelBeforeUpdateHooks, problemRelHook)
	case boil.BeforeDeleteHook:
		problemRelBeforeDeleteHooks = append(problemRelBeforeDeleteHooks, problemRelHook)
	case boil.BeforeUpsertHook:
		problemRelBeforeUpsertHooks = append(problemRelBeforeUpsertHooks, problemRelHook)
	case boil.AfterInsertHook:
		problemRelAfterInsertHooks = append(problemRelAfterInsertHooks, problemRelHook)
	case boil.AfterSelectHook:
		problemRelAfterSelectHooks = append(problemRelAfterSelectHooks, problemRelHook)
	case boil.AfterUpdateHook:
		problemRelAfterUpdateHooks = append(problemRelAfterUpdateHooks, problemRelHook)
	case boil.AfterDeleteHook:
		problemRelAfterDeleteHooks = append(problemRelAfterDeleteHooks, problemRelHook)
	case boil.AfterUpsertHook:
		problemRelAfterUpsertHooks = append(problemRelAfterUpsertHooks, problemRelHook)
	}
}

// OneG returns a single problemRel record from the query using the global executor.
func (q problemRelQuery) OneG() (*ProblemRel, error) {
	return q.One(boil.GetDB())
}

// One returns a single problemRel record from the query.
func (q problemRelQuery) One(exec boil.Executor) (*ProblemRel, error) {
	o := &ProblemRel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for problem_rels")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProblemRel records from the query using the global executor.
func (q problemRelQuery) AllG() (ProblemRelSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all ProblemRel records from the query.
func (q problemRelQuery) All(exec boil.Executor) (ProblemRelSlice, error) {
	var o []*ProblemRel

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProblemRel slice")
	}

	if len(problemRelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProblemRel records in the query, and panics on error.
func (q problemRelQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all ProblemRel records in the query.
func (q problemRelQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count problem_rels rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q problemRelQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q problemRelQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if problem_rels exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *ProblemRel) Category(mods ...qm.QueryMod) problemCategoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	query := ProblemCategories(queryMods...)
	queries.SetFrom(query.Query, "\"problem_categories\"")

	return query
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (problemRelL) LoadCategory(e boil.Executor, singular bool, maybeProblemRel interface{}, mods queries.Applicator) error {
	var slice []*ProblemRel
	var object *ProblemRel

	if singular {
		object = maybeProblemRel.(*ProblemRel)
	} else {
		slice = *maybeProblemRel.(*[]*ProblemRel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &problemRelR{}
		}
		if !queries.IsNil(object.CategoryID) {
			args = append(args, object.CategoryID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &problemRelR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CategoryID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.CategoryID) {
				args = append(args, obj.CategoryID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`problem_categories`),
		qm.WhereIn(`problem_categories.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ProblemCategory")
	}

	var resultSlice []*ProblemCategory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ProblemCategory")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for problem_categories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for problem_categories")
	}

	if len(problemRelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &problemCategoryR{}
		}
		foreign.R.CategoryProblemRels = append(foreign.R.CategoryProblemRels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.CategoryID, foreign.ID) {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &problemCategoryR{}
				}
				foreign.R.CategoryProblemRels = append(foreign.R.CategoryProblemRels, local)
				break
			}
		}
	}

	return nil
}

// SetCategoryG of the problemRel to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.CategoryProblemRels.
// Uses the global database handle.
func (o *ProblemRel) SetCategoryG(insert bool, related *ProblemCategory) error {
	return o.SetCategory(boil.GetDB(), insert, related)
}

// SetCategory of the problemRel to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.CategoryProblemRels.
func (o *ProblemRel) SetCategory(exec boil.Executor, insert bool, related *ProblemCategory) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"problem_rels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
		strmangle.WhereClause("\"", "\"", 2, problemRelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.CategoryID, related.ID)
	if o.R == nil {
		o.R = &problemRelR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &problemCategoryR{
			CategoryProblemRels: ProblemRelSlice{o},
		}
	} else {
		related.R.CategoryProblemRels = append(related.R.CategoryProblemRels, o)
	}

	return nil
}

// RemoveCategoryG relationship.
// Sets o.R.Category to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *ProblemRel) RemoveCategoryG(related *ProblemCategory) error {
	return o.RemoveCategory(boil.GetDB(), related)
}

// RemoveCategory relationship.
// Sets o.R.Category to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *ProblemRel) RemoveCategory(exec boil.Executor, related *ProblemCategory) error {
	var err error

	queries.SetScanner(&o.CategoryID, nil)
	if _, err = o.Update(exec, boil.Whitelist("category_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Category = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.CategoryProblemRels {
		if queries.Equal(o.CategoryID, ri.CategoryID) {
			continue
		}

		ln := len(related.R.CategoryProblemRels)
		if ln > 1 && i < ln-1 {
			related.R.CategoryProblemRels[i] = related.R.CategoryProblemRels[ln-1]
		}
		related.R.CategoryProblemRels = related.R.CategoryProblemRels[:ln-1]
		break
	}
	return nil
}

// ProblemRels retrieves all the records using an executor.
func ProblemRels(mods ...qm.QueryMod) problemRelQuery {
	mods = append(mods, qm.From("\"problem_rels\""))
	return problemRelQuery{NewQuery(mods...)}
}

// FindProblemRelG retrieves a single record by ID.
func FindProblemRelG(iD int, selectCols ...string) (*ProblemRel, error) {
	return FindProblemRel(boil.GetDB(), iD, selectCols...)
}

// FindProblemRel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProblemRel(exec boil.Executor, iD int, selectCols ...string) (*ProblemRel, error) {
	problemRelObj := &ProblemRel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"problem_rels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, problemRelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from problem_rels")
	}

	return problemRelObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProblemRel) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProblemRel) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no problem_rels provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemRelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	problemRelInsertCacheMut.RLock()
	cache, cached := problemRelInsertCache[key]
	problemRelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			problemRelAllColumns,
			problemRelColumnsWithDefault,
			problemRelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(problemRelType, problemRelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(problemRelType, problemRelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"problem_rels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"problem_rels\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into problem_rels")
	}

	if !cached {
		problemRelInsertCacheMut.Lock()
		problemRelInsertCache[key] = cache
		problemRelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single ProblemRel record using the global executor.
// See Update for more documentation.
func (o *ProblemRel) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the ProblemRel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProblemRel) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	problemRelUpdateCacheMut.RLock()
	cache, cached := problemRelUpdateCache[key]
	problemRelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			problemRelAllColumns,
			problemRelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update problem_rels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"problem_rels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, problemRelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(problemRelType, problemRelMapping, append(wl, problemRelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update problem_rels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for problem_rels")
	}

	if !cached {
		problemRelUpdateCacheMut.Lock()
		problemRelUpdateCache[key] = cache
		problemRelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q problemRelQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q problemRelQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for problem_rels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for problem_rels")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProblemRelSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProblemRelSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemRelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"problem_rels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, problemRelPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in problemRel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all problemRel")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProblemRel) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProblemRel) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no problem_rels provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(problemRelColumnsWithDefault, o)

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

	problemRelUpsertCacheMut.RLock()
	cache, cached := problemRelUpsertCache[key]
	problemRelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			problemRelAllColumns,
			problemRelColumnsWithDefault,
			problemRelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			problemRelAllColumns,
			problemRelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert problem_rels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(problemRelPrimaryKeyColumns))
			copy(conflict, problemRelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"problem_rels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(problemRelType, problemRelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(problemRelType, problemRelMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert problem_rels")
	}

	if !cached {
		problemRelUpsertCacheMut.Lock()
		problemRelUpsertCache[key] = cache
		problemRelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single ProblemRel record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProblemRel) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single ProblemRel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProblemRel) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProblemRel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), problemRelPrimaryKeyMapping)
	sql := "DELETE FROM \"problem_rels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from problem_rels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for problem_rels")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q problemRelQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q problemRelQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no problemRelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problem_rels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problem_rels")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProblemRelSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProblemRelSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(problemRelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemRelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"problem_rels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemRelPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from problemRel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for problem_rels")
	}

	if len(problemRelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProblemRel) ReloadG() error {
	if o == nil {
		return errors.New("models: no ProblemRel provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProblemRel) Reload(exec boil.Executor) error {
	ret, err := FindProblemRel(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemRelSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ProblemRelSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProblemRelSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProblemRelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), problemRelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"problem_rels\".* FROM \"problem_rels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, problemRelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProblemRelSlice")
	}

	*o = slice

	return nil
}

// ProblemRelExistsG checks if the ProblemRel row exists.
func ProblemRelExistsG(iD int) (bool, error) {
	return ProblemRelExists(boil.GetDB(), iD)
}

// ProblemRelExists checks if the ProblemRel row exists.
func ProblemRelExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"problem_rels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if problem_rels exists")
	}

	return exists, nil
}
