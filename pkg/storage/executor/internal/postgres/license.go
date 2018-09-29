package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/pkg/errors"
)

// NewLicenseContext TODO issue#docs
func NewLicenseContext(ctx context.Context, conn *sql.Conn) licenseManager {
	return licenseManager{ctx, conn}
}

type licenseManager struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#docs
func (scope licenseManager) Create(token *repository.Token, data query.CreateLicense) (repository.License, error) {
	entity := repository.License{Contract: data.Contract}
	before, encodeErr := json.Marshal(domain.Contract{})
	if encodeErr != nil {
		return entity, errors.Wrapf(encodeErr,
			"user %q of account %q with token %q tried to encode to JSON empty license contract",
			token.UserID, token.User.AccountID, token.UserID)
	}
	after, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, errors.Wrapf(encodeErr,
			"user %q of account %q with token %q tried to encode to JSON license contract %+v",
			token.UserID, token.User.AccountID, token.UserID, entity.Contract)
	}
	q := `INSERT INTO "license" ("id", "account_id", "contract")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, token.User.AccountID, after)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Wrapf(scanErr,
			"user %q of account %q with token %q tried to create new license with contract %s",
			token.UserID, token.User.AccountID, token.UserID, after)
	}
	{
		audit := `INSERT INTO "license_audit" ("license_id", "contract", "what", "when", "who", "with")
		          VALUES ($1, $2, $3, $4, $5, $6)`
		if _, execErr := scope.conn.ExecContext(scope.ctx, audit, entity.ID, before,
			repository.Create, entity.CreatedAt, token.UserID, token.ID); execErr != nil {
			return entity, errors.Wrapf(execErr,
				"audit: user %q of account %q with token %q tried to create new license with contract %s",
				token.UserID, token.User.AccountID, token.UserID, after)
		}
	}
	return entity, nil
}

// Read TODO issue#docs
func (scope licenseManager) Read(token *repository.Token, data query.ReadLicense) (repository.License, error) {
	entity, encoded := repository.License{ID: data.ID}, []byte(nil)
	q := `SELECT "contract", "created_at", "updated_at", "deleted_at"
	        FROM "license"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, token.User.AccountID)
	if scanErr := row.Scan(&encoded, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); scanErr != nil {
		return entity, errors.Wrapf(scanErr,
			"user %q of account %q with token %q tried to read license %q",
			token.UserID, token.User.AccountID, token.UserID, entity.ID)
	}
	if decodeErr := json.Unmarshal(encoded, &entity.Contract); decodeErr != nil {
		return entity, errors.Wrapf(decodeErr,
			"user %q of account %q with token %q tried to decode contract %s of license %q from JSON",
			token.UserID, token.User.AccountID, token.UserID, encoded, entity.ID)
	}
	return entity, nil
}

// Update TODO issue#docs
func (scope licenseManager) Update(token *repository.Token, data query.UpdateLicense) (repository.License, error) {
	entity, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return entity, errors.Wrapf(readErr, "while updating")
	}
	before, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, errors.Wrapf(encodeErr,
			"user %q of account %q with token %q tried to encode to JSON current contract %+v of license %q",
			token.UserID, token.User.AccountID, token.UserID, entity.Contract, entity.ID)
	}
	after, encodeErr := json.Marshal(data.Contract)
	if encodeErr != nil {
		return entity, errors.Wrapf(encodeErr,
			"user %q of account %q with token %q tried to encode to JSON new contract %+v of license %q",
			token.UserID, token.User.AccountID, token.UserID, entity.Contract, entity.ID)
	}
	q := `UPDATE "license"
	         SET "contract" = $1
	       WHERE "id" = $2
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, after, entity.ID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Wrapf(scanErr,
			"user %q of account %q with token %q tried to update license %q with new contract %s",
			token.UserID, token.User.AccountID, token.UserID, entity.ID, after)
	}
	{
		audit := `INSERT INTO "license_audit" ("license_id", "contract", "what", "when", "who", "with")
		          VALUES ($1, $2, $3, $4, $5, $6)`
		if _, execErr := scope.conn.ExecContext(scope.ctx, audit, entity.ID, before,
			repository.Update, entity.UpdatedAt, token.UserID, token.ID); execErr != nil {
			return entity, errors.Wrapf(execErr,
				"audit: user %q of account %q with token %q tried to update license %q with new contract %s",
				token.UserID, token.User.AccountID, token.UserID, entity.ID, after)
		}
	}
	return entity, nil
}

// Delete TODO issue#docs
func (scope licenseManager) Delete(token *repository.Token, data query.DeleteLicense) (repository.License, error) {
	entity, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return entity, errors.Wrapf(readErr, "while deleting")
	}
	before, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, errors.Wrapf(encodeErr,
			"user %q of account %q with token %q tried to encode to JSON current contract %+v of license %q",
			token.UserID, token.User.AccountID, token.UserID, entity.Contract, entity.ID)
	}
	q := `UPDATE "license"
	         SET "deleted_at" = now()
	       WHERE "id" = $1
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Wrapf(scanErr,
			"user %q of account %q with token %q tried to delete license %q",
			token.UserID, token.User.AccountID, token.UserID, entity.ID)
	}
	{
		audit := `INSERT INTO "license_audit" ("license_id", "contract", "what", "when", "who", "with")
		          VALUES ($1, $2, $3, $4, $5, $6)`
		if _, execErr := scope.conn.ExecContext(scope.ctx, audit, entity.ID, before,
			repository.Delete, entity.DeletedAt, token.UserID, token.ID); execErr != nil {
			return entity, errors.Wrapf(execErr,
				"audit: user %q of account %q with token %q tried to delete license %q",
				token.UserID, token.User.AccountID, token.UserID, entity.ID)
		}
	}
	return entity, nil
}
