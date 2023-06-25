package services

import "fmt"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Albums struct {
	Items []Album
}

func (Albums *Albums) FindIndex(ID string) (int, error) {
	for index, album := range Albums.Items {
		if album.ID == ID {
			return index, nil
		}
	}

	return -1, fmt.Errorf("not found index of %s", ID)
}

func (Albums *Albums) DeleteByIndex(index int) {
	Albums.Items = append(Albums.Items[:index], Albums.Items[index+1:]...)
}

func (Albums *Albums) Create(album Album) {
	Albums.Items = append(Albums.Items, album)
}
