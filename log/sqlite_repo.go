package log

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/johnweldon/go/util"
)

type RelDB struct {
	dbpath  string
	ensured bool
	db      *sql.DB
}

func NewRelDB(path string) *RelDB {
	return &RelDB{dbpath: path, ensured: false, db: nil}
}

func (d *RelDB) Open() func() {
	if d.db != nil {
		err := d.db.Ping()
		if err == nil {
			return func() {}
		}
		log.Print(err)
		d.db = nil
	}
	db, err := sql.Open("sqlite3", d.dbpath)
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
	if !d.ensured {
		err := ensureDb(d.db)
		if err != nil {
			log.Fatal(err)
		}
		d.ensured = true
	}
	return func() {
		d.db.Close()
		d.db = nil
	}
}

func (d *RelDB) GetRecords() []TimeRecord {
	defer d.Open()()

	result := []TimeRecord{}

	rows, err := d.db.Query(sqlSelectRecords)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var begin string
		var duration int64
		var durationstring string
		var project string
		var notes string
		var tagslist string

		err := rows.Scan(&id, &begin, &duration, &durationstring, &project, &notes, &tagslist)
		if err != nil {
			log.Fatal(err)
		}

		btime, err := time.Parse(time.RFC3339, begin)
		if err != nil {
			log.Fatal(err)
		}

		tags := strings.Split(tagslist, ",")
		result = append(result, TimeRecord{
			ID:             util.UUID(id),
			Begin:          btime,
			Duration:       time.Duration(duration),
			DurationString: durationstring,
			Project:        project,
			Notes:          notes,
			Tags:           tags})
	}

	return result
}

func (d *RelDB) SaveRecords(records []TimeRecord) error {
	defer d.Open()()

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	insProj, err := tx.Prepare(sqlInsertProject)
	if err != nil {
		return err
	}
	defer insProj.Close()

	insRec, err := tx.Prepare(sqlInsertRecord)
	if err != nil {
		return err
	}
	defer insRec.Close()

	insTag, err := tx.Prepare(sqlInsertTag)
	if err != nil {
		return err
	}
	defer insTag.Close()

	insRecTag, err := tx.Prepare(sqlInsertRecordTag)
	if err != nil {
		return err
	}
	defer insRecTag.Close()

	for _, record := range records {
		_, err = insProj.Exec(record.Project)
		if err != nil {
			return err
		}
		_, err = insRec.Exec(string(record.ID), record.Begin.Format(time.RFC3339), record.Duration, record.DurationString, record.Project, record.Notes)
		if err != nil {
			return err
		}
		for _, tag := range record.Tags {
			_, err = insTag.Exec(tag)
			if err != nil {
				return err
			}
			_, err = insRecTag.Exec(string(record.ID), tag)
			if err != nil {
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func ensureDb(db *sql.DB) error {
	for _, stmt := range []string{sqlSetupDb, sqlCreateTags, sqlCreateProjects, sqlCreateRecords, sqlCreateRecordsTags} {
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	sqlSetupDb    string = "PRAGMA foreign_keys = ON"
	sqlCreateTags string = `
	    CREATE TABLE IF NOT EXISTS tags (
	        Name TEXT
	            CONSTRAINT pk_tag_name PRIMARY KEY
	            ON CONFLICT ROLLBACK
	    ) `
	sqlCreateProjects string = `
	    CREATE TABLE IF NOT EXISTS projects (
	        Name TEXT
	            CONSTRAINT pk_project_name PRIMARY KEY
	            ON CONFLICT ROLLBACK
        ) `
	sqlCreateRecords string = `
	    CREATE TABLE IF NOT EXISTS records (
            ID TEXT
                CONSTRAINT pk_record_id PRIMARY KEY
                ON CONFLICT ROLLBACK,
            Begin TEXT,
            Duration INTEGER ,
            DurationString TEXT ,
            Project TEXT
                CONSTRAINT fk_records_project_name REFERENCES projects ( Name )
                ON DELETE CASCADE ON UPDATE CASCADE ,
            Notes TEXT
        ) `
	sqlCreateRecordsTags string = `
	    CREATE TABLE IF NOT EXISTS records_tags (
            RecordID TEXT
                CONSTRAINT fk_records_tags_record_id REFERENCES records ( ID )
                ON DELETE CASCADE ON UPDATE CASCADE ,
            TagName TEXT
                CONSTRAINT fk_records_tags_tag_name REFERENCES tags ( Name )
                ON DELETE CASCADE ON UPDATE CASCADE ,
            CONSTRAINT pk_records_tags_recordid_tagname PRIMARY KEY ( RecordID, TagName )
                ON CONFLICT ROLLBACK
        ) `

	sqlInsertTag     string = "INSERT OR IGNORE INTO tags ( Name ) VALUES ( ? )"
	sqlInsertProject string = "INSERT OR IGNORE INTO projects ( Name ) VALUES ( ? )"
	sqlInsertRecord  string = `
        INSERT OR IGNORE INTO records
        ( ID, Begin, Duration, DurationString, Project, Notes ) VALUES
        (  ?,     ?,        ?,              ?,       ?,     ? )
    `
	sqlInsertRecordTag string = `
        INSERT OR IGNORE INTO records_tags
        ( RecordID, TagName ) VALUES
        (        ?,       ? )
    `

	sqlSelectRecords string = `
	    SELECT
	        R.ID ,
	        R.Begin ,
	        R.Duration ,
	        R.DurationString ,
	        R.Project ,
	        R.Notes ,
	        group_concat(RT.TagName) AS Tags
	    FROM records AS R
	    LEFT OUTER JOIN projects AS P ON R.Project = P.Name
	    LEFT OUTER JOIN records_tags AS RT ON RT.RecordID = R.ID
	    GROUP BY R.ID
    `
)
