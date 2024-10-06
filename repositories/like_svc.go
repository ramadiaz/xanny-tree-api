package repositories

func (r *compRepository) AddLike() (*int64, error) {
	var total int64

	err := r.DB.QueryRow("UPDATE like_count SET total = total + 1 WHERE name = 'profile' RETURNING total;").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &total, nil
}
