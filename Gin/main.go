package main

import "net/http"
import "github.com/gin-gonic/gin"

// album represents data about a record album.
// Note : Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON.
// Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums",postAlbums)
	
    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum.
	if err := context.BindJSON(&newAlbum); err != nil {
		// if err object is non nil i.e if it contains an error then we return from the function. If not, then code continues execution
		return
	}

	// We will add new album to the slice
	albums = append(albums, newAlbum)
    //                  Add a 201 status code
    context.IndentedJSON(http.StatusCreated,newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}