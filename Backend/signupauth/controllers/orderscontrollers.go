package controllers

import (
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/models"
	"database/sql"
	"fmt"
)
func GetOrders() ([]models.Orderstable, error) {

	rows, err := database.DB.Query("SELECT Orid,Status,Description,Disputeeligibility from models.Orderstable")
	// Rows are selected to display the contents of Orders table
	if err != nil {
		return nil, err
	}
	//Makes sure to close the rows after the action is done
	defer rows.Close()

	Orders := make([]models.Orderstable, 0,9)

	for rows.Next() {
		eachOrder := models.Orderstable{}
		err = rows.Scan(&eachOrder.Orid, &eachOrder.Status, &eachOrder.Description, &eachOrder.Disputeeligibility)

		if err != nil {
			return nil, err
		}
		Orders = append(Orders, eachOrder)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return Orders, err
}

func GetOrderById(Orid uint) (models.Orderstable, error) {

	stmt, err := database.DB.InstanceGet("SELECT Orid,Status,Description,Disputeeligibility from models.Orderstable WHERE Orid=?")

	if err != nil {
		return models.Orderstable{}, err
	}

	eachOrderById := models.Orderstable{}

	sqlErr := stmt.QueryRow(Orid).Scan(&eachOrderById.Orid, &eachOrderById.Status, &eachOrderById.Description, &eachOrderById.Disputeeligibility)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Orderstable{}, nil
		}
		return models.Orderstable{}, sqlErr
	}
	return eachOrderById, nil
}

func UpdateTransitStatus(transitStatus models.Orderstable, Orid uint) (bool, error) {

	ctx, err := database.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := ctx.Prepare("UPDATE models.Orderstable SET status = ? WHERE Orid = ?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(transitStatus.models.Orderstable, transitStatus.Orid)

	if err != nil {
		return false, err
	}

	ctx.Commit()

	return true, nil
}

func CancelLuggageOrder(Orid uint) (bool, error) {

	ctx, err := database.DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := database.DB.Prepare("DELETE from models.Orderstable where Orid = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(Orid)

	if err != nil {
		return false, err
	}

	ctx.Commit()

	return true, nil
}
