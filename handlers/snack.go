package handlers

import (
	"net/http"
	"snacks-backend/db"
	"snacks-backend/models"

	"github.com/labstack/echo/v4"
)

func GetSnacks(c echo.Context) error {
	rows, err := db.DB.Query("SELECT id, name, category, price, image_url, health_benefits, description, discount, ingredients FROM snacks")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "DB query failed"})
	}
	defer rows.Close()

	var snacks []models.Snack
	for rows.Next() {
		var s models.Snack
		if err := rows.Scan(&s.ID, &s.Name, &s.Category, &s.Price, &s.ImageURL, &s.HealthBenefits, &s.Description, &s.Discount, &s.Ingredients); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Row scan failed"})
		}
		snacks = append(snacks, s)
	}

	return c.JSON(http.StatusOK, snacks)
}

func CreateSnack(c echo.Context) error {
	var s models.Snack
	if err := c.Bind(&s); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	_, err := db.DB.Exec("INSERT INTO snacks (name, category, price, image_url, health_benefits, description, discount, ingredients) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		s.Name, s.Category, s.Price, s.ImageURL, s.HealthBenefits, s.Description, s.Discount, s.Ingredients,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Insert failed"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Snack created"})
}

func DeleteSnack(c echo.Context) error {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM snacks WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Delete failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Snack deleted"})
}

func UpdateSnack(c echo.Context) error {
	id := c.Param("id")
	var s models.Snack
	if err := c.Bind(&s); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	_, err := db.DB.Exec("UPDATE snacks SET name=?, category=?, price=?, image_url=?, health_benefits=?, description=?, discount=?, ingredients=? WHERE id=?",
		s.Name, s.Category, s.Price, s.ImageURL, s.HealthBenefits, s.Description, s.Discount, s.Ingredients, id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Update failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Snack updated"})
}
