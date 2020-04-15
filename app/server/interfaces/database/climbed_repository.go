package database

import (
	"fmt"

	"github.com/shiki-tak/shiki-web/app/server/domain"
)

type ClimbedMountainRepository struct {
	SqlHandler
}

const (
	fields    = "name, height, climbed_date,weather, number, description, image_url"
	values    = "VALUES "
	v7        = values + "(?, ?, ?, ?, ?, ?, ?)"
	tableName = "climbed_mountains"
	insertC   = "INSERT INTO "
	selectC   = "SELECT "
	from      = "FROM "
)

func (repo *ClimbedMountainRepository) Store(m domain.ClimbedMountain) error {
	// TODO: climbed_mountain -> change to const value
	insertStatement := fmt.Sprintf("%s %s (%s) %s", insertC, tableName, fields, v7)
	result, err := repo.Execute(insertStatement, m.Name, m.Height, m.ClimbedDate, m.Weather, m.Number, m.Description, m.ImageURL)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (repo *ClimbedMountainRepository) FindById(key int) (domain.ClimbedMountain, error) {
	var climbedMountain domain.ClimbedMountain
	getStatement := fmt.Sprintf("%s %s (%s) %s", selectC, tableName, fields, v7)
	row, err := repo.Query(getStatement, key)
	defer row.Close()
	if err != nil {
		return climbedMountain, err
	}

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
	climbedMountain.ID = key
	climbedMountain.Name = name
	climbedMountain.Height = height
	climbedMountain.ClimbedDate = climbedDate
	climbedMountain.Weather = weather
	climbedMountain.Number = number
	climbedMountain.Description = description
	climbedMountain.ImageURL = imageURL

	return climbedMountain, nil
}

func (repo *ClimbedMountainRepository) Gets() ([]domain.ClimbedMountain, error) {
	ｍountains := []domain.ClimbedMountain{}
	getsStatement := fmt.Sprintf("%s id %s %s %s", selectC, fields, from, tableName)
	rows, err := repo.Query(getsStatement)
	defer rows.Close()
	if err != nil {
		return ｍountains, err
	}
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
