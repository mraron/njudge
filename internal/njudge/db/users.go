package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/njudge/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (us *Users) toNjudge(ctx context.Context, u *models.User) (*njudge.User, error) {
	res := njudge.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: njudge.HashedPassword(u.Password),
		Email:    u.Email,
		ActivationInfo: njudge.UserActivationInfo{
			Activated: !u.ActivationKey.Valid,
			Key:       u.ActivationKey.String,
		},
		Role:   u.Role,
		Points: u.Points.Float32,
		Settings: njudge.UserSettings{
			ShowUnsolvedTags: u.ShowUnsolvedTags,
		},
	}

	if key := u.R.ForgottenPasswordKey; key != nil {
		res.ForgottenPasswordKey = &njudge.ForgottenPasswordKey{
			ID:         key.ID,
			UserID:     u.ID,
			Key:        key.Key,
			ValidUntil: key.Valid,
		}
	}

	return &res, nil
}

// converts to *models.User, ignoring forgotten password fkey
func (us *Users) toModel(u njudge.User) *models.User {
	res := &models.User{
		ID:               u.ID,
		Name:             u.Name,
		Password:         string(u.Password),
		Email:            u.Email,
		ActivationKey:    null.NewString(u.ActivationInfo.Key, !u.ActivationInfo.Activated),
		Role:             u.Role,
		Points:           null.Float32From(u.Points),
		ShowUnsolvedTags: u.Settings.ShowUnsolvedTags,
	}
	return res
}

func (us *Users) get(ctx context.Context, mods ...qm.QueryMod) (*njudge.User, error) {
	dbobj, err := models.Users(append(mods, qm.Load(models.UserRels.ForgottenPasswordKey))...).One(ctx, us.db)
	if err != nil {
		return nil, MaskNotFoundError(err, njudge.ErrorUserNotFound)
	}

	res, err := us.toNjudge(ctx, dbobj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *Users) Get(ctx context.Context, ID int) (*njudge.User, error) {
	return us.get(ctx, models.UserWhere.ID.EQ(ID))
}

func (us *Users) GetByName(ctx context.Context, name string) (*njudge.User, error) {
	return us.get(ctx, models.UserWhere.Name.EQ(name))
}

func (us *Users) GetByEmail(ctx context.Context, email string) (*njudge.User, error) {
	return us.get(ctx, models.UserWhere.Email.EQ(email))
}

// inserts a njudge.User to the database with a possible forgotten password key
func (us *Users) Insert(ctx context.Context, u njudge.User) (*njudge.User, error) {
	dbobj := us.toModel(u)
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := dbobj.Insert(ctx, tx, boil.Infer()); err != nil {
		pgerr := &pq.Error{}
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				if pgerr.Constraint == "users_name_unique" {
					return nil, errors.Join(njudge.ErrorSameName, err, tx.Rollback())
				}
				if pgerr.Constraint == "users_email_unique" {
					return nil, errors.Join(njudge.ErrorSameEmail, err, tx.Rollback())
				}
			}
		}
		return nil, errors.Join(err, tx.Rollback())
	}

	if u.ForgottenPasswordKey != nil {
		key := &models.ForgottenPasswordKey{
			Key:   u.ForgottenPasswordKey.Key,
			Valid: u.ForgottenPasswordKey.ValidUntil,
		}

		if err := dbobj.SetForgottenPasswordKey(ctx, tx, true, key); err != nil {
			return nil, errors.Join(err, tx.Rollback())
		}
	}

	return &u, tx.Commit()
}

func (us *Users) Delete(ctx context.Context, ID int) error {
	_, err := models.Users(models.UserWhere.ID.EQ(ID)).DeleteAll(ctx, us.db)
	return err
}

// updates the user's given fields
func (us *Users) Update(ctx context.Context, u *njudge.User, fields []string) error {
	dbobj := us.toModel(*u)
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	whitelist := make([]string, 0, len(fields))
	for ind := range fields {
		switch fields[ind] {
		case njudge.UserFields.Name:
			whitelist = append(whitelist, models.UserColumns.Name)
		case njudge.UserFields.Password:
			whitelist = append(whitelist, models.UserColumns.Password)
		case njudge.UserFields.Email:
			whitelist = append(whitelist, models.UserColumns.Email)
		case njudge.UserFields.ActivationInfo:
			whitelist = append(whitelist, models.UserColumns.ActivationKey)
		case njudge.UserFields.Role:
			whitelist = append(whitelist, models.UserColumns.Role)
		case njudge.UserFields.Points:
			whitelist = append(whitelist, models.UserColumns.Points)
		case njudge.UserFields.Settings:
			whitelist = append(whitelist, models.UserColumns.ShowUnsolvedTags)
		}
	}

	if len(whitelist) > 0 {
		if _, err := dbobj.Update(ctx, tx, boil.Whitelist(whitelist...)); err != nil {
			return errors.Join(err, tx.Rollback())
		}
	}

	if u.ForgottenPasswordKey != nil {
		if u.ForgottenPasswordKey.ID != 0 {
			key := &models.ForgottenPasswordKey{
				Key:   u.ForgottenPasswordKey.Key,
				Valid: u.ForgottenPasswordKey.ValidUntil,
			}

			if err := dbobj.SetForgottenPasswordKey(ctx, tx, false, key); err != nil {
				return errors.Join(err, tx.Rollback())
			}
		} else {
			key := &models.ForgottenPasswordKey{
				Key:   u.ForgottenPasswordKey.Key,
				Valid: u.ForgottenPasswordKey.ValidUntil,
			}

			if err := dbobj.SetForgottenPasswordKey(ctx, tx, true, key); err != nil {
				return errors.Join(err, tx.Rollback())
			}

			u.ForgottenPasswordKey.ID = key.ID
		}
	} else {
		_, err = models.ForgottenPasswordKeys(models.ForgottenPasswordKeyWhere.UserID.EQ(u.ID)).DeleteAll(ctx, tx)
		if err != nil {
			return errors.Join(err, tx.Rollback())
		}
	}

	return tx.Commit()
}
