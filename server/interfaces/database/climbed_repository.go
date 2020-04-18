package database

import (
	"fmt"

	"github.com/shiki-tak/shiki-web/server/domain"
)

type ClimbedMountainRepository struct {
	SqlHandler
}

const (
	fields        = "name, height, climbed_date, weather, number, description, image_url"
	uFieldsFormat = "name=?, height=?, climbed_date=?, weather=?, number=?, description=?, image_url=?"

	values = "VALUES "
	v7     = values + "(?, ?, ?, ?, ?, ?, ?)"

	tableName = "climbed_mountains"

	insertC = "INSERT INTO"
	updateC = "UPDATE"
	selectC = "SELECT"
	deleteC = "DELETE"

	from  = "FROM"
	where = "WHERE"
	set   = "SET"
)

func (repo *ClimbedMountainRepository) Store(m domain.ClimbedMountain) error {
	var err error
	err = repo.TxBegin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = repo.TxRollback()
		}
	}()

	insertStatement := fmt.Sprintf("%s %s (%s) %s", insertC, tableName, fields, v7)
	result, err := repo.Execute(insertStatement, m.Name, m.Height, m.ClimbedDate, m.Weather, m.Number, m.Description, m.ImageURL)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	if err = repo.TxCommit(); err != nil {
		return err
	}

	return nil
}

func (repo *ClimbedMountainRepository) Update(m domain.ClimbedMountain) error {
	var err error
	err = repo.TxBegin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = repo.TxRollback()
		}
	}()

	updateStatement := fmt.Sprintf("%s %s %s %s %s id=%d", updateC, tableName, set, uFieldsFormat, where, m.ID)
	_, err = repo.Execute(updateStatement, m.Name, m.Height, m.ClimbedDate, m.Weather, m.Number, m.Description, m.ImageURL)
	if err != nil {
		return err
	}

	if err = repo.TxCommit(); err != nil {
		return err
	}

	return nil
}

func (repo *ClimbedMountainRepository) FindById(id int) (domain.ClimbedMountain, error) {
	var climbedMountain domain.ClimbedMountain
	getStatement := fmt.Sprintf("%s %s %s %s %s id=?", selectC, fields, from, tableName, where)

	row, err := repo.Query(getStatement, id)
	if err != nil {
		return climbedMountain, err
	}
	defer row.Close()

	row.Next()
	var (
		name        string
		height      int
		climbedDate string
		weather     string
		number      int
		description string
		imageURL    string
	)
	if err = row.Scan(&name, &height, &climbedDate, &weather, &number, &description, &imageURL); err != nil {
		return climbedMountain, err
	}
	climbedMountain.ID = id
	climbedMountain.Name = name
	climbedMountain.Height = height
	climbedMountain.ClimbedDate = climbedDate
	climbedMountain.Weather = weather
	climbedMountain.Number = number
	climbedMountain.Description = description
	climbedMountain.ImageURL = imageURL

	return climbedMountain, nil
}

func (repo *ClimbedMountainRepository) FindAll() ([]domain.ClimbedMountain, error) {
	ｍountains := []domain.ClimbedMountain{}
	getsStatement := fmt.Sprintf("%s id, %s %s %s", selectC, fields, from, tableName)
	rows, err := repo.Query(getsStatement)
	if err != nil {
		return ｍountains, err
	}
	defer rows.Close()

	for rows.Next() {
		mountain := domain.ClimbedMountain{}
		if err = rows.Scan(&mountain.ID, &mountain.Name, &mountain.Height, &mountain.ClimbedDate,
			&mountain.Weather, &mountain.Number, &mountain.Description, &mountain.ImageURL); err != nil {
			return ｍountains, err
		}
		ｍountains = append(ｍountains, mountain)
	}
	return ｍountains, nil
}

func (repo *ClimbedMountainRepository) Delete(id int) error {
	var err error
	err = repo.TxBegin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = repo.TxRollback()
		}
	}()

	deleteStatement := fmt.Sprintf("%s %s %s %s id=%d", deleteC, from, tableName, where, id)
	_, err = repo.Execute(deleteStatement)
	if err != nil {
		return err
	}

	if err = repo.TxCommit(); err != nil {
		return err
	}

	return nil
}
