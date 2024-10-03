package repositories

import "xanny-tree-api/dto"

func (r *compRepository) RegisterUrl(data dto.Tree) error {
	_, err := r.DB.Exec("INSERT INTO tree_url (name, short_url, original_url) VALUES($1, $2, $3)", data.Name, data.ShortUrl, data.OriginalUrl)
	if err != nil {
		return err
	}

	return nil
}