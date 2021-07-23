package Albums

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

// Data structure
// Chosen slice of struct as to facilitate fast searching with any field

type Album struct {
	Uuid          string `json:"uuid"`
	Album_Title   string `json:"album_title"`
	Artist        string `json:"artist"`
	Year_Released string `json:"year_released"`
}

type Albums []Album

var AllAlbums = allAlbums()

//------------------------------------------------------------------------------------------------------
// for use adding to the collection from csv file
func newCsvAlbum(uuid string, album_Title string, artist string, year_Released string) *Album {
	return &Album{Uuid: uuid, Album_Title: album_Title, Artist: artist, Year_Released: year_Released}
}

// reads in csv file
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
//gets all albums from a csv file into memory
func allAlbums() Albums {
	records := readCsvFile("../rms_albums.csv")
	var albums []Album
	for i, record := range records {
		if i != 0 {
			csvRecord := newCsvAlbum(record[0], record[1], record[2], record[3])
			albums = append(albums, *csvRecord)
		}
	}

	return albums
}

// adds the passed Album to the collection
func (a *Albums) AddAlbum(album Album) error {

	// ckeck to make sure an album with this id does not already exist
	if !a.AlbumIdExists(album.Uuid) {
		// add the passed album to the Albums collection
		*a = append(*a, album)
		return nil
	}

	return errors.New("an album with chosen id already exists")

}

func (a Albums) AlbumIdExists(id string) bool {
	for _, record := range a {
		if record.Uuid == id {
			return true
		}
	}
	return false
}

// Gets an individual record by Id
// As the Record Id's are unique this will only ever return one record
func (a Albums) GetAlbumById(id string) (Album, error) {

	for _, album := range a {
		if album.Uuid == id {
			return album, nil
		}
	}
	return Album{}, errors.New("error: album with provided id does not exist")
}

// returns all the Albums for a given Artist name
func (a Albums) GetAlbumsByArtist(artist string) []Album {
	var albums []Album
	for _, album := range a {
		if album.Artist == artist {
			albums = append(albums, album)
		}
	}

	return albums
}

// updates the Title, Artist, and Year_Released of Album to match the Album passed in
func (a Albums) Update(updatedalbum Album) Album {
	for _, album := range a {
		// if album found
		if album.Uuid == updatedalbum.Uuid {
			//update fields
			album.Album_Title = updatedalbum.Album_Title
			album.Artist = updatedalbum.Artist
			album.Year_Released = updatedalbum.Year_Released
			return album
		}

	}
	return Album{}
}

// this will not retain the order of the data
func (a Albums) DeleteAlbum(id string) error {

	if !a.AlbumIdExists(id) {
		return fmt.Errorf("error deleting album: album with provided id does not exist") //errors.New("album does not exist")
	}
	for i, album := range a {
		if album.Uuid == id {
			a[i] = a[len(a)-1]
			a = a[:len(a)-1]

		}
	}
	return nil
}


