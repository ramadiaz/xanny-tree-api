package repositories

import "xanny-tree-api/dto"

func (r *compRepository) RegisterUrl(data dto.Tree) error {
	_, err := r.DB.Exec("INSERT INTO tree_url (name, short_url, original_url) VALUES($1, $2, $3)", data.Name, data.ShortUrl, data.OriginalUrl)
	if err != nil {
		return err
	}

	return nil
}

func (r *compRepository) GetUrl(url string) (*string, error) {
	var origin string

	stmt, err := r.DB.Prepare("SELECT original_url FROM tree_url WHERE short_url = $1 ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(url).Scan(&origin)
	if err != nil {
		return nil, err
	}

	return &origin, nil
}
