package Albums_test

import (
	"fmt"
	"records/Albums"
	"testing"
)

var AllAlbums = Albums.AllAlbums


//----------------------------------------------------------------------------------------------------------------------
// Add new Album -------------------------------------------------------------------------------------------------------

func Example_AddNewAlbum() {

	album := Albums.Album{
		Uuid:          "65f59662-6d69-42b1-b549-f97e2ed872e2",
		Album_Title:   "The Bends",
		Artist:        "Radiohead ",
		Year_Released: "1995",
	}
	err := AllAlbums.AddAlbum(album)
	if  err == nil {
		record, err := AllAlbums.GetAlbumById(album.Uuid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(record)
	}
	// Output: {65f59662-6d69-42b1-b549-f97e2ed872e2 The Bends Radiohead  1995}
}


// Test_AddNewAlbum_Error checks that the correct error is returned when trying to add an album with an id that is
// already taken
func Test_AddNewAlbum_Error(t *testing.T) {

	album := Albums.Album{
		Uuid:          "f337fd51-7bf5-44bf-9553-5826162bc83a",
		Album_Title:   "Pink Noise",
		Artist:        "Laura Mvula",
		Year_Released: "2021",
	}

	err := AllAlbums.AddAlbum(album)

	expectedError := "an album with chosen id already exists"
	if err.Error() != expectedError {
		t.Errorf("got: %s, want: %s", err.Error(), expectedError)
	}

}
//----------------------------------------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------------------------------------
// GetAlbumById ---------------------------------------------------------------------------------------

const missingId = "d9ca9e60-093c-422a-841c-1b5ef48a522f"
const idKnownToExist = "f337fd51-7bf5-44bf-9553-5826162bc83a"
// tests getting an album we know exists
func Example_GetAlbumById() {
	id := "f337fd51-7bf5-44bf-9553-5826162bc83a"
	album, err := AllAlbums.GetAlbumById(id)

	if err != nil {
		panic(err)
	}


	fmt.Print(album)

    // Output: {f337fd51-7bf5-44bf-9553-5826162bc83a Pink Noise Laura Mvula 2021}
}


// test for an attempt at getting an album using an id that does not exist
func Test_GetAlbumById_Error(t *testing.T) {

	_, err := AllAlbums.GetAlbumById(missingId)

	expectedError := "error: album with provided id does not exist"
	if err.Error() != expectedError {
		t.Errorf("got: %s, want: %s", err.Error(), expectedError)
	}

}
//----------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------
// Test_DeleteAlbum ----------------------------------------------------------------------------------------------------
func Test_DeleteAlbum(t *testing.T) {
	const id = "17cd04a4-ef0a-468f-9f47-5d9dbb1c0dbd"
	err := AllAlbums.DeleteAlbum(id)

	if err != nil {
		t.Errorf("error even thought we know the album does exist and can be deleted")
	}

	if AllAlbums.AlbumIdExists(id) {
		t.Errorf("album still exists even after an attempt to remove")
	}

}

// test the attempted deletion of an album we know currently does not exist in collection
// returns the expected error
// ---------------------------------------------------------------------------------------------------------------------
func Test_DeleteAlbum_Error(t *testing.T) {
	const missingId = "d9ca9e60-093c-422a-841c-1b5ef48a522f"
	err := AllAlbums.DeleteAlbum(missingId)

	expectedError := "error deleting album: album with provided id does not exist"
	if err.Error() != expectedError {
		t.Errorf("got: %s, want: %s", err.Error(), expectedError)
	}

}

// ---------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------
// Test_GetAlbumByArtist -----------------------------------------------------------------------------------------------
func Example_GetAlbumsByArtist() {
	albums := AllAlbums.GetAlbumsByArtist("Gorillaz")
	for _, album := range albums {
		fmt.Println(album)
	}

	// Output:
	// {19fe1358-70b0-419a-b58b-c5da5990a75e Demon Days Gorillaz 2005}
	// {77b0b2a6-1fb3-4f7f-adcd-df2c1be11416 Plastic Beach Gorillaz 2010}
	// {b0f4e28a-f41b-4b79-96a4-2776bfb73c7b Demon Days Gorillaz 2005}
}

func Example_UpdateAlbum_Melorama() {

	albumToUpdate := Albums.Album{
		Uuid:          "c2263b8c-6718-4494-8218-1e739cf04e0a",
		Album_Title:   "Melodrama",
		Artist:        "Lorde",
		Year_Released: "2017"}

	album := AllAlbums.Update(albumToUpdate)

	fmt.Println(album)

	// Output:
	// {c2263b8c-6718-4494-8218-1e739cf04e0a Melodrama Lorde 2017}
}

// Example_UpdateAlbum_Fake ensures all fields are updated
func Example_UpdateAlbum_Fake() {

	albumToUpdate := Albums.Album{
		Uuid:          "c2263b8c-6718-4494-8218-1e739cf04e0a",
		Album_Title:   "updated_title",
		Artist:        "updated artist",
		Year_Released: "updated year released"}

	album := AllAlbums.Update(albumToUpdate)

	fmt.Println(album)

	// Output:
	// {c2263b8c-6718-4494-8218-1e739cf04e0a updated_title updated artist updated year released}
}



