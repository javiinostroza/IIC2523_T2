# Scraping [MyAnimeList](https://myanimelist.net/topanime.php)

### Este algoritmo realiza un scraping al Top Anime de [MyAnimeList](https://myanimelist.net/topanime.php) para obtener los 200 animes con mejor ranking según los usuarios de la plataforma. Los datos que se recolectan son almacenados en un archivo CSV con nombre ```data.csv```.


### Ejecutar scraper:
- Usar el comando ```go run scraper.go```.

### Consideraciones:

- Fue necesario ejecutar ```go mod init oxylabs.io/web-scraping-with-go``` para crear un archivo ```go.mod``` y ```go get github.com/gocolly/colly``` para instalar Colly y sus dependencias. Este último comando además modifica el archivo ```go.mod``` y crea un archivo ```go.sum```.

- Los datos que se recolectan en ```data.csv``` corresponden a:
     - **Rank:** número en el ranking que tiene el anime
     - **Title:** título del anime
     - **Score:** Puntaje del anime (otorgado por los usuarios)
     - **Type:** Tipo de anime (TV, película, etc.)
     - **Members:** Cantidad de usuarios que agregaron el anime a su lista.