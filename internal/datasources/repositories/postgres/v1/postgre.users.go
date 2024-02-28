package v1

import (
	"context"
	"encoding/hex"

	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/dmzsz/duozhuayu/internal/datasources/records"
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
	"github.com/dmzsz/duozhuayu/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	ChangeActiveUser(ctx context.Context, inDom *V1Domains.UserDomain) (err error)
	Delete(ctx context.Context, inDom *V1Domains.UserDomain) (err error)
	GetByEmail(ctx context.Context, inDom *V1Domains.UserDomain, decryptEmail bool) (outDomain V1Domains.UserDomain, err error)
	// GetById(ctx context.Context, id string) (outDomain V1Domains.UserDomain, err error)
	GetByWithRoleByField(ctx context.Context, columnName string, columnValue string) (outDomain V1Domains.UserDomain, err error)
	Store(ctx context.Context, inDom *V1Domains.UserDomain) (err error)
}

type postgreUserRepository struct {
	conn *sqlx.DB
	// config string
}

func NewUserRepository(conn *sqlx.DB) UserRepository {
	return &postgreUserRepository{
		conn: conn,
	}
}
func (r *postgreUserRepository) ChangeActiveUser(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	_, err = r.conn.NamedQueryContext(ctx, `UPDATE users SET active = :active WHERE id = :id`, userRecord)

	return
}
func (r *postgreUserRepository) Delete(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO users(id, username, email, password, active, role_id, created_at) VALUES (uuid_generate_v4(), :username, :email, :password, false, :role_id, :created_at)`, userRecord)
	if err != nil {
		return err
	}

	return nil
}

// GetById implements v1.UserRepository.
// func (r *postgreUserRepository) GetById(ctx context.Context, id string) (outDomain V1Domains.UserDomain, err error) {
// 	userRecord := records.FromUsersV1Domain(&V1Domains.UserDomain{Id: id})

// 	err = r.conn.GetContext(ctx, &userRecord, `SELECT * FROM users WHERE "id" = $1`, id)
// 	if err != nil {
// 		return V1Domains.UserDomain{}, err
// 	}

// 	return userRecord.ToV1Domain(), nil
// }

func (r *postgreUserRepository) GetByWithRoleByField(ctx context.Context, columnName string, columnValue string) (outDomain V1Domains.UserDomain, err error) {
	userRecord := records.FromUsersV1Domain(&V1Domains.UserDomain{})

	err = r.conn.GetContext(ctx, &userRecord,
		`SELECT
			u.*,
			r.*
		FROM
			(SELECT * FROM user WHERE $1 = $2) u
		LEFT JOIN
			user_to_role ur ON u.user_id = ur.user_id
		LEFT JOIN
			role r ON ur.role_id = r.role_id`,
		columnName,
		columnValue)
	if err != nil {
		return V1Domains.UserDomain{}, err
	}

	return userRecord.ToV1Domain(), nil
}

func (r *postgreUserRepository) GetByEmail(ctx context.Context, inDom *V1Domains.UserDomain, decryptEmail bool) (outDomain V1Domains.UserDomain, err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	if err = r.conn.GetContext(ctx, &userRecord, `SELECT * FROM users WHERE "email" = $1`, userRecord.Email); err == nil {
		return userRecord.ToV1Domain(), nil
	}

	// encryption at rest
	if configs.IsCipher() {
		// hash of the email in hexadecimal string format
		emailHash, err := helpers.CalcHash(
			userRecord.Email,
			configs.AppConfig.SecurityConfig.Blake2bSec,
		)
		if err != nil {
			return V1Domains.UserDomain{}, err
		}

		err = r.conn.GetContext(ctx, &userRecord, `SELECT * FROM users WHERE "email_hash" = $1`, emailHash)
		// email must be unique
		if err == nil {
			if decryptEmail {
				userRecord.Email, err = helpers.DecryptEmail(userRecord.EmailNonce, userRecord.EmailCipher, configs.AppConfig.SecurityConfig.CipherKey)
				if err != nil {
					return V1Domains.UserDomain{}, err
				}
			}

			return userRecord.ToV1Domain(), nil
		}
	}

	return V1Domains.UserDomain{}, err
}

func (r *postgreUserRepository) Store(ctx context.Context, inDom *V1Domains.UserDomain) (err error) {
	userRecord := records.FromUsersV1Domain(inDom)

	if configs.IsCipher() {
		// encrypt the email
		var cipherEmail, nonce []byte
		cipherEmail, nonce, err = helpers.EncryptChacha20poly1305(
			[]byte(configs.AppConfig.SecurityConfig.CipherKey),
			userRecord.Email,
		)
		if err != nil {
			// log.WithError(err).Error("error code: 1001.2")
			// httpResponse.Message = "internal server error"
			// httpStatusCode = http.StatusInternalServerError
			return err
		}

		// save email only in ciphertext
		userRecord.Email = ""
		userRecord.EmailCipher = hex.EncodeToString(cipherEmail)
		userRecord.EmailNonce = hex.EncodeToString(nonce)
	}

	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO users(id, username, email, password, active, role_id, created_at) VALUES (uuid_generate_v4(), :username, :email, :password, false, :role_id, :created_at)`, userRecord)
	if err != nil {
		return err
	}

	return nil
}
