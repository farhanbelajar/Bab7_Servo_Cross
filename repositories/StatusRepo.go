package repositories

import (
	"database/sql"
	"farhan_s/entities"
)

func UtkString(db *sql.DB) (err error) {
	sql := "INSERT INTO Bab7Servo(Code, StatusServo) values(1,0)"
	_, err = db.Query(sql)
	return err
}

func LihatStatus(db *sql.DB) (result []entities.Status, err error) {
	sql := "SELECT * FROM Bab7Servo"
	rows, err := db.Query(sql)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Status
		err = rows.Scan(&data.Code, &data.ServoStatus)
		if err != nil {
			return
		}
		result = append(result, data)
	}
	return
}

func UbahStatus(db *sql.DB, status entities.Status) (err error) {
	sql := "UPDATE Bab7Servo SET StatusServo = $1 WHERE Code = 1"
	_, err = db.Exec(sql, status.ServoStatus)
	return
}
