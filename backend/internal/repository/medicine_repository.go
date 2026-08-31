package repository

import (
	"context"
	"go-pharma/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineRepository struct {
	db *pgxpool.Pool
}

func NewMedicineRepository(db *pgxpool.Pool) *MedicineRepository {
	return &MedicineRepository{db: db}
}

func (r *MedicineRepository) GetAll(ctx context.Context, search string) ([]model.Medicine, error) {
	query := `SELECT id, name, category, stock, unit, price FROM medicines`
	args := []interface{}{}

	if search != "" {
		query += ` WHERE name ILIKE $1 OR category ILIKE $1`
		args = append(args, "%"+search+"%")
	}
	query += ` ORDER BY id`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []model.Medicine
	for rows.Next() {
		var m model.Medicine
		if err := rows.Scan(&m.Id, &m.Name, &m.Category, &m.Stock, &m.Unit, &m.Price); err != nil {
			return nil, err
		}
		medicines = append(medicines, m)
	}
	return medicines, nil
}

func (r *MedicineRepository) Create(ctx context.Context, m *model.Medicine) error {
	query := `INSERT INTO medicines (name, category, stock, unit, price)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRow(ctx, query, m.Name, m.Category, m.Stock, m.Unit, m.Price).Scan(&m.Id)
}

func (r *MedicineRepository) Update(ctx context.Context, id int, m *model.Medicine) error {
	query := `UPDATE medicines SET name=$1, category=$2, stock=$3, unit=$4, price=$5 WHERE id=$6`
	cmdTag, err := r.db.Exec(ctx, query, m.Name, m.Category, m.Stock, m.Unit, m.Price, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *MedicineRepository) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.db.Exec(ctx, `DELETE FROM medicines WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
