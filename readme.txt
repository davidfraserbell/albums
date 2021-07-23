To run unit tests

assuming go is installed type the following command...
go test -v ./...

you should then see the following output

=== RUN   Test_AddNewAlbum_Error
--- PASS: Test_AddNewAlbum_Error (0.00s)
=== RUN   Test_GetAlbumById_Error
--- PASS: Test_GetAlbumById_Error (0.00s)
=== RUN   Test_DeleteAlbum
--- PASS: Test_DeleteAlbum (0.00s)
=== RUN   Test_DeleteAlbum_Error
--- PASS: Test_DeleteAlbum_Error (0.00s)
=== RUN   Example_AddNewAlbum
--- PASS: Example_AddNewAlbum (0.00s)
=== RUN   Example_GetAlbumById
--- PASS: Example_GetAlbumById (0.00s)
=== RUN   Example_GetAlbumsByArtist
--- PASS: Example_GetAlbumsByArtist (0.00s)
=== RUN   Example_UpdateAlbum_Melorama
--- PASS: Example_UpdateAlbum_Melorama (0.00s)
=== RUN   Example_UpdateAlbum_Fake
--- PASS: Example_UpdateAlbum_Fake (0.00s)
PASS
ok      records/Albums  0.076s

