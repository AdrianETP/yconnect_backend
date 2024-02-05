# yconnect backend

Yconnect es una app de ios para conectar organizaciones con usuarios. Este repo contiene el backend hecho en golang con el framework __fiber__ para todas las llamadas web

## Estructura
- config: configuraciones de apis (en este caso la base de datos de mongodb)
- models: los modelos para las bases de datos y para tipado estatico en el backend
- controllers: las funciones a las que se va a llamar en el api
- rotues: el lugar donde se asigna que rutas van a llamar a que funcion


## Routes
Aqui les dejo una lista con los paths disponibles hasta ahora y sus datos

### Organizations

```
/organizations/GetAll
```

- tipo: Post
- body: 
```json
{
    "token": "token del usuario"
}
``` 
- descripcion: obtiene todas las organizaciones



```
/organizations
```

- tipo: POST
- body: 
```json
{
    
    "organization": objeto de la organizacion
    "token": token del usuario

}

}`
- descripcion: agrega una organizacion

```
/organizations/searchByTag
```

- tipo: POST
- body:

```json
{
    "tags": [] // los tags que quieras buscar
    "token": "token del usuario"
}
```

- descripcion: obtiene organizaciones que tienen los tags que buscas

```
/organizations/Favorites
```

- tipo: POST
- body:

```json
{
  "userId": "id del usuario"
    "token": "token del usuario"
}
```

- descripcion: obtiene todas las organizaciones favoritas del usuario

```
/organizations/Favorites
```
- Tipo: POST
- body: 

```json
{
    "name":"el nombre de la organizacion"
"token": "token del usuario"
}
```
- descripcion: una manera de buscar organizaciones por nombre


```
/organizations/Delete
```
- Tipo: POST
- body:

```json
{
    "orgid":"id de la organizacion"
"token": "token del usuario"
}
```
- Descripcion: borrar una organizacion

```
/organizations/SearchById
```
- Tipo: POST
- body:

```json
{
    "orgid":"id de la organizacion"
"token": "token del usuario
}
```
- Descripcion: Manera de buscar una organizacion por id

```
/organizations/ModifyOrg
```

- Tipo: POST
- body: 
```json
{
    "organization": objeto de la organizacion
    "token": token del usuario
}
```
- Descripcion: una manera de modificar una organizacion

```
/organizations/SendMail
```

- Tipo: POST
- body: 
```json
{
    "organization": objeto de la organizacion
    "token": token del usuario
}
```
- Descripcion: le manda un mail al administrador de que una empresa se quiere registrar


### Users

```
/users/getAll
```

- metodo: POST
- body: 
```json 
{
    "token": "token del usuario"
}
```
- descripcion: obtiene todos los usuarios

```
/users
```

- metodo: POST
- body: 
```json
{
    "user": objeto del usuario
    "token": token del usuario
}
```
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
        "token": "token del usuario"
    }
```

- description: agrega una organizacion a los favoritos de un usuario

```
/users/addTags
```
- Tipo: POST
- body:
```json
{
    "userid":"id del usuario"
    "tags":["tags que quieras agregar"]
    "token": "token del usuario"
}
```
- Descripcion: agregar tags a un usuario

```
/users/Delete
```
- Tipo: POST
- Body:
```json
    {
        "userid":"id del usuario",
        "token": "token del usuario"
    }
```
- Descripcion: borrar un usuario

```
/users/Update
```
- Tipo: POST
- Body:
```json
    {
    "user":objeto del usuario
    "token": "token del usuario"
    }
```
- Descripcion: actualizar un usuario


```
/users/Login
```
- Tipo: POST
- Body:
```json
{
    "email":"email del usuario",
    "password":"password del usuario


}
```
- Descripcion: iniciar sesion de un usuario

### Posts

```
/posts
```

- tipo: GET
- body: un objeto json con la organizacion
- descripcion: obtiene todas los posts de una organizacion

```json
{
    "orgId": "id de la organizacion"
}
```

```
/posts
```

- tipo: POST
- body: nueva publicacion
- descripcion: crea una nueva publicacion para una organizacion

```json
{
    "orgId": "id de la organizacion",
    "content": "text de la publicacion",
    "media": [] // links de las imágenes o videos
}
```


```
/posts/addLike
```

- tipo: POST
- body: id de la publicacion y el usuario
- descripcion: agrega a usuario a la lista de likes de la publicacon

```json
{
    "userId": "id del usuario",
    "postId": "id de la publicacion"
}
```

```
/posts/addComment
```

- tipo: POST
- body: id de la publicacion, el usuario, y el texto del comentario
- descripcion: agrega comentario a una publicacion

```json
{
    "userId": "id del usuario",
    "postId": "id de la publicacion",
    "content": "Hola mundo"
}
```

