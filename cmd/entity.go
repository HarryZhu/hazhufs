package cmd

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/mattn/go-sqlite3"
)

type Entity struct {
	Db                 *bolt.DB
	Key                string
	Meta               []byte
	Data               []byte
	beforeSaveHandlers []func()
	afterSaveHandlers  []func()
}

func (e *Entity) BeforeSave(fn func()) {
	e.beforeSaveHandlers = append(e.beforeSaveHandlers, fn)
}

func (e *Entity) AfterSave(fn func()) {
	e.afterSaveHandlers = append(e.afterSaveHandlers, fn)
}

func (e *Entity) MetaExists() bool {
	Logger.Debug("MetaExists")
	var b_return bool = false
	var k string
	stmt, err := DBMETA.Prepare("SELECT key FROM data WHERE key = ? LIMIT 1")
	if err != nil {
		Logger.Error(err)
	}

	err = stmt.QueryRow(string(e.Key)).Scan(&k)

	if err != nil {
		return false
	}
	if len(k) > 0 {
		b_return = true
	}
	return b_return
}

func (e *Entity) DataExists() bool {
	Logger.Debug("DataExists")
	var b_return bool = false
	db := e.Db
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("meta"))
		if bucket != nil {
			m := bucket.Get([]byte(e.Key))
			if m != nil {
				b_return = true
			} else {
				b_return = false
			}
		}
		return errors.New("cannot get the bucket")
	})
	return b_return
}

func (e *Entity) MetaUpdate() error {
	Logger.Debug("MetaUpdate")
	var isExisting bool = false
	isExisting = e.MetaExists()
	if isExisting == true {
		return nil
	}

	stmt, err := DBMETA.Prepare("INSERT INTO data(key,size,dbid,node,synced,created) values(?,?,?,?,?,?)")
	if err != nil {
		Logger.Error(err)
		return errors.New("failed.")
	}

	node := strings.Join([]string{CFG.Volume.Ip, CFG.Volume.Port}, ":")
	_, err = stmt.Exec(e.Key, len(e.Data), (e.Key)[0:1], node, 0, time.Now().Unix())
	if err != nil {
		Logger.Error(err)
		return errors.New("failed.")
	}
	return nil
}

func (e *Entity) Get() *Entity {
	Logger.Debug("Get")
	var r_meta []byte
	var r_data []byte
	var r_entity = &Entity{}

	db := e.Db
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("meta"))
		if bucket != nil {
			r_meta = bucket.Get([]byte(e.Key))
			if r_meta == nil {
				return errors.New("cannot get the meta")
			}
		}
		return errors.New("cannot get the bucket meta")
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("data"))
		if bucket != nil {
			r_data = bucket.Get([]byte(e.Key))
			if r_data == nil {
				return errors.New("cannot get the data")
			}
		}
		return errors.New("cannot get the bucket data")
	})
	//Logger.Warn("#####mmm", r_meta)
	//Logger.Warn("#####ddd", r_data)

	if r_meta != nil && r_data != nil {
		r_entity.Db = e.Db
		r_entity.Key = e.Key
		r_entity.Meta = r_meta
		r_entity.Data = r_data
		return r_entity
	}

	return nil

}

func (e *Entity) Save() error {
	if e.DataExists() == true {
		Logger.Debug("Existed, will ignore.")
		return nil
	}
	for _, fn_before := range e.beforeSaveHandlers {
		fn_before()
	}
	Logger.Debug("Save")

	db := e.Db

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt_meta, err_meta := tx.CreateBucketIfNotExists([]byte("meta"))
	if err_meta != nil {
		return err_meta
	}

	bkt_data, err_data := tx.CreateBucketIfNotExists([]byte("data"))
	if err_data != nil {
		return err_data
	}

	put_err_meta := bkt_meta.Put([]byte(e.Key), e.Meta)
	if put_err_meta != nil {
		return put_err_meta
	}
	put_err_data := bkt_data.Put([]byte(e.Key), e.Data)
	if put_err_data != nil {
		return put_err_data
	}

	err_meta_update := e.MetaUpdate()
	if err_meta_update != nil {
		return err_meta_update
	}

	err_commit := tx.Commit()
	if err_commit != nil {
		Logger.Error("SaveFile: cannot commit transaction.")
		return err_commit
	} else {
		Logger.Debug("SaveFile, Commit OK: ", e.Key)
	}

	for _, fn_after := range e.afterSaveHandlers {
		fn_after()
	}

	return nil
}

func SaveFile(db *bolt.DB, key string, meta []byte, data []byte) error {
	if len(key) != 32 {
		return errors.New("key length should be 32.")
	}

	entity := &Entity{Db: db, Key: strings.ToLower(key), Meta: meta, Data: data}
	//entity.BeforeSave()
	entity.Save()
	//entity.AfterSave()

	return nil
}

func GetFile(db *bolt.DB, key string) *Entity {
	if len(key) != 32 {
		Logger.Warn("key length should be 32.")
		return nil
	}
	KEY := strings.ToLower(key)

	entity := &Entity{Db: db, Key: KEY}
	e := entity.Get()
	if e != nil {
		//		Logger.Info("entity.Get.ERROR.")
		return e
	}
	Logger.Error("entity.Get.ERROR.")
	return nil
}

func GenerateSyncNeededListFromMeta() error {
	var arr_keys []string
	var key string
	var rowsCount int = 0
	var str_keys string = ""
	var file_synclist string = strings.Join([]string{"data/synclist_", CFG.Volume.Ip, "_", CFG.Volume.Port, ".txt"}, "")

	Logger.Debug("GenerateSyncNeededListFromMeta...")

	rows, err := DBMETA.Query("SELECT key FROM data where synced = ? limit 100", "0")

	if err != nil {
		Logger.Debug("Error: ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&key)
		if err != nil {
			Logger.Error("Error: ", key)
		} else {
			rowsCount++
			arr_keys = append(arr_keys, key)
		}
	}

	if rowsCount > 0 {
		str_keys = strings.Join(arr_keys, ";")
		Logger.Debug("OK: ", str_keys)
	}

	f, err := os.OpenFile(file_synclist, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		Logger.Fatal(err)
	}
	f.Write([]byte(str_keys))
	f.Close()

	return nil
}
