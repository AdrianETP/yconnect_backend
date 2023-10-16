# yconnect backend

Aqui les dejo una lista con los paths disponibles hasta ahora y sus datos

## organizations
```
/organizations
```
- tipo: GET
- body: nada
- descripcion: obtiene todas las organizaciones

```
/organizations
```

- tipo: POST
- body: un objeto json cos la organizacion
- descripcion: agrega una organizacion


```
/organizations/searchByTag
```

- tipo: POST
- body: 
```json
{
        "tags":[] // los tags que quieras buscar
    }
```
- descripcion: obtiene organizaciones que tienen los tags que buscas



## Posts (redes sociales)

### Instagram
```
/posts/ig/GetFromTag
```

- tipo: POST
- body: 
```json
{
        "tags":[] // las tags que quieras buscar
    }
```
- descripcion:parecido al `organizations/searghByTag` pero en vez de una organizacion te da los posts de instagram de dicha organizacion


## Users

```
/users
```
- metodo: GET
- body: nada
- descripcion: obtiene todos los usuarios


```
/users
```
- metodo: POST
- body: el usuario que quieras agregar en JSON
- descripcion: agcega un usuario a la base de datos

```
/users/addFavorites
```
- metodo: POST
- body:
```json
{
        "user":"el id del usuario"
        "organization":"el id de la organizacion"
    }
```
- description: agrega una organizacion a los favoritos de un usuario
