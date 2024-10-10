package repositories

import "time"

func (r *compRepository) AddLike() (*int64, error) {
	var total int64

	err := r.DB.QueryRow("UPDATE like_count SET total = total + 1, updated_at = $1 WHERE name = 'profile' RETURNING total;", time.Now()).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &total, nil
}

func (r *compRepository) GetLike() (*int64, error) {
	var total int64

	err := r.DB.QueryRow("SELECT total FROM like_count WHERE name = 'profile';").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &total, nil
}
