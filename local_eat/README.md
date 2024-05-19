# Localeat

website Local-Eat : [Local-Eat](https://www.localeat.ephec-ti.be/)  

Il est possible de se créer un compte Utilisateur, dans l'améliorer en compte producteur et, en tant que producteur créer et se connecter à une compagnie. 

Si connecté à une compagnie le producteur pourra ajouter des produits au catalogue de cette compagnie.

## Environment setup

You need to have [Go](https://golang.org/),
[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows)
installed on your computer.

Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

## Start in development mode

In the project directory run the command (you might
need to prepend it with `sudo` depending on your setup):
```sh
docker compose up
```

This starts a local MySQL database on `localhost:3306`.
The database will be populated with test records from
the [init-db.sql](init-db.sql) file.

Navigate to the `server` folder and start the back end:

```sh
cd api
go run main.go
```
The back end will serve on http://localhost:8080.

Navigate to the `webapp` folder, install dependencies,
and start the front end development server by running:

```sh
cd web
npm install
npm start
```
The application will be available on http://localhost:3000.
 
## Start in production mode

Perform:
```sh
docker compose -f docker-compose-deployment.yml up
```
This will build 4 containers:
- db: the database container which is connected to the api
- api: the api container connected to the db and web container, the image needs to be rebuilt each time a new version of the api is done
- web: the web server connected to the api, it only needs to be rebuilt when the servers config are changed but if you rebuild the angular app you do need to restart the container for it to work correctly
- certbot: the certbot container generates the ssl certificates to enable https on the web server
