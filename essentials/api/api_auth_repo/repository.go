package api_auth_repo

import (
	"database/sql"
	"encoding/base64"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"os"
	"path/filepath"
	"time"
)

// NewPersistent creates new Repository that persist to the file
func NewPersistent(path string) (r api_auth.Repository, err error) {
	l := esl.Default().With(esl.String("path", path))
	info, err := os.Lstat(path)
	if err == nil && info.IsDir() {
		l.Debug("Create auth database with default database name")
		return NewPersistent(filepath.Join(path, app_definitions2.AuthDatabaseDefaultName))
	}

	if os.IsNotExist(err) {
		l.Debug("Create new database")
		tmpDb, err := sql.Open("sqlite3", path)
		if err != nil {
			l.Debug("Unable to create the database", esl.Error(err))
			return nil, err
		}
		if _, err = tmpDb.Exec(`CREATE TABLE app (version TEXT, timestamp DATETIME)`); err != nil {
			l.Debug("Unable to create app version table", esl.Error(err))
			return nil, err
		}
		if _, err = tmpDb.Exec("INSERT INTO app (version, timestamp) VALUES(?, ?)", app_definitions2.Version.String(), time.Now()); err != nil {
			l.Debug("Unable to record app version data", esl.Error(err))
			return nil, err
		}
		if err := tmpDb.Close(); err != nil {
			l.Debug("Unable to close the database", esl.Error(err))
			return nil, err
		}
		l.Debug("Protect the database")
		if err := os.Chmod(path, 0600); err != nil {
			l.Debug("Unable to protect the database, try remove it")
			rmErr := os.Remove(path)
			l.Debug("Remove operation finished", esl.Error(rmErr))
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return newWithDb(db)
}

// NewInMemory creates new in-memory repository
func NewInMemory() (r api_auth.Repository, err error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	return newWithDb(db)
}

func newWithDb(db *sql.DB) (r api_auth.Repository, err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS repository (
	build_stream TEXT NOT NULL,
	key_name     TEXT NOT NULL,
	scopes       TEXT NOT NULL,
	peer_name    TEXT NOT NULL,
	credential   TEXT,
	description  TEXT,
	entity_ts    TEXT,
	PRIMARY KEY(build_stream, key_name, scopes, peer_name)
	)`)
	if err != nil {
		return nil, err
	}

	sp, err := db.Prepare(`INSERT INTO repository (
                        build_stream, key_name, scopes, peer_name, credential, description, entity_ts
                        ) VALUES (?, ?, ?, ?, ?, ?, ?)
						ON CONFLICT(build_stream, key_name, scopes, peer_name)
						DO UPDATE SET credential = ?, description = ?, entity_ts = ?
                        `)
	if err != nil {
		return nil, err
	}

	sg, err := db.Prepare(`SELECT credential, description, entity_ts FROM repository WHERE build_stream = ? AND key_name = ? AND scopes = ? AND peer_name = ?`)
	if err != nil {
		return nil, err
	}

	sd, err := db.Prepare(`DELETE FROM repository WHERE build_stream = ? AND key_name = ? AND scopes = ? AND peer_name = ?`)
	if err != nil {
		return nil, err
	}

	sl, err := db.Prepare(`SELECT peer_name, credential, description, entity_ts FROM repository WHERE build_stream = ? AND key_name = ? AND scopes = ? ORDER BY peer_name`)
	if err != nil {
		return nil, err
	}

	sa, err := db.Prepare(`SELECT key_name, scopes, peer_name, credential, description, entity_ts FROM repository WHERE build_stream = ? ORDER BY key_name, peer_name, scopes`)
	if err != nil {
		return nil, err
	}

	return &sqlRepo{
		stmtPut:  sp,
		stmtGet:  sg,
		stmtDel:  sd,
		stmtList: sl,
		stmtAll:  sa,
		db:       db,
	}, nil
}

type sqlRepo struct {
	stmtPut  *sql.Stmt
	stmtGet  *sql.Stmt
	stmtDel  *sql.Stmt
	stmtList *sql.Stmt
	stmtAll  *sql.Stmt
	db       *sql.DB
}

func (z sqlRepo) Close() {
	l := esl.Default()
	if err := z.db.Close(); err != nil {
		l.Debug("Unable to close", esl.Error(err))
	}
}

func (z sqlRepo) Put(entity api_auth.Entity) {
	l := esl.Default()
	cred, err := sc_obfuscate.Obfuscate(l, sc_obfuscate.XapKey(), []byte(entity.Credential))
	if err != nil {
		l.Debug("Unable to obfuscate the credential", esl.Error(err))
		return
	}
	encCred := base64.StdEncoding.EncodeToString(cred)

	_, err = z.stmtPut.Exec(
		// insert into
		sc_obfuscate.BuildStream(),
		entity.KeyName,
		entity.Scope,
		entity.PeerName,
		encCred,
		entity.Description,
		entity.Timestamp,

		// update set
		encCred,
		entity.Description,
		entity.Timestamp,
	)
	if err != nil {
		l.Debug("Unable to insert/update data", esl.Error(err))
	}
}

func (z sqlRepo) decodeCredential(credObf string) (cred string, found bool) {
	l := esl.Default()
	credDec, err := base64.StdEncoding.DecodeString(credObf)
	if err != nil {
		l.Debug("Unable to decode credential", esl.Error(err))
		return "", false
	}
	credRaw, err := sc_obfuscate.Deobfuscate(l, sc_obfuscate.XapKey(), credDec)
	if err != nil {
		l.Debug("Unable to de-obfuscate", esl.Error(err))
		return "", false
	}
	return string(credRaw), true
}

func (z sqlRepo) Get(keyName, scope, peerName string) (entity api_auth.Entity, found bool) {
	l := esl.Default()
	r, err := z.stmtGet.Query(sc_obfuscate.BuildStream(), keyName, scope, peerName)
	if err != nil {
		l.Debug("Query failure", esl.Error(err))
		return entity, false
	}
	if !r.Next() {
		l.Debug("Not found")
		return entity, false
	}

	var credObf, desc, entityTs string
	if err := r.Scan(&credObf, &desc, &entityTs); err != nil {
		l.Debug("Cannot retrieve", esl.Error(err))
		return entity, false
	}
	credRaw, found := z.decodeCredential(credObf)
	if !found {
		return entity, false
	}
	if err := r.Close(); err != nil {
		l.Debug("Unable to close the result", esl.Error(err))
	}

	return api_auth.Entity{
		KeyName:     keyName,
		Scope:       scope,
		PeerName:    peerName,
		Credential:  credRaw,
		Description: desc,
		Timestamp:   entityTs,
	}, true
}

func (z sqlRepo) Delete(keyName, scope, peerName string) {
	l := esl.Default()
	_, err := z.stmtDel.Exec(sc_obfuscate.BuildStream(), keyName, scope, peerName)
	if err != nil {
		l.Debug("Unable to delete the entry", esl.Error(err))
	}
}

func (z sqlRepo) List(keyName, scope string) (entities []api_auth.Entity) {
	l := esl.Default()
	entities = make([]api_auth.Entity, 0)

	r, err := z.stmtList.Query(sc_obfuscate.BuildStream(), keyName, scope)
	if err != nil {
		l.Debug("Query failure", esl.Error(err))
		return entities
	}

	for r.Next() {
		var peerName, credObf, desc, entityTs string
		if err := r.Scan(&peerName, &credObf, &desc, &entityTs); err != nil {
			l.Debug("Cannot retrieve, skip", esl.Error(err))
			continue
		}
		credRaw, found := z.decodeCredential(credObf)
		if !found {
			continue
		}

		entities = append(entities, api_auth.Entity{
			KeyName:     keyName,
			Scope:       scope,
			PeerName:    peerName,
			Credential:  credRaw,
			Description: desc,
			Timestamp:   entityTs,
		})
	}
	if err := r.Close(); err != nil {
		l.Debug("Unable to close the result", esl.Error(err))
	}

	return entities
}

func (z sqlRepo) All() (entities []api_auth.Entity) {
	l := esl.Default()
	entities = make([]api_auth.Entity, 0)

	r, err := z.stmtAll.Query(sc_obfuscate.BuildStream())
	if err != nil {
		l.Debug("Query failure", esl.Error(err))
		return entities
	}

	for r.Next() {
		var keyName, scopes, peerName, credObf, desc, entityTs string
		if err := r.Scan(&keyName, &scopes, &peerName, &credObf, &desc, &entityTs); err != nil {
			l.Debug("Cannot retrieve, skip", esl.Error(err))
			continue
		}
		credRaw, found := z.decodeCredential(credObf)
		if !found {
			continue
		}

		entities = append(entities, api_auth.Entity{
			KeyName:     keyName,
			Scope:       scopes,
			PeerName:    peerName,
			Credential:  credRaw,
			Description: desc,
			Timestamp:   entityTs,
		})
	}
	if err := r.Close(); err != nil {
		l.Debug("Unable to close the result", esl.Error(err))
	}

	return entities
}
