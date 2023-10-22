# yconnect backen 

Yconnect es una app de ios para conectar organizaciones con usuarios. Este repo contiene el backend hecho en golang con el framework __fiber__ para todas las llamadas web

## Estructura
- certs: los certificados (autosigned) de ssl
- config: configuraciones de apis (en este caso la base de datos de mongodb)
- models: los modelos para las bases de datos y para tipado estatico en el backend
- controllers: las funciones a las que se va a llamar en el api
- rotues: el lugar donde se asigna que rutas van a llamar a que funcion


## Routes
Aqui les dejo una lista con los paths disponibles hasta ahora y sus datos

### organizations

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
  "tags": [] // los tags que quieras buscar
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
}
```
- Descripcion: Manera de buscar una organizacion por id

```
/organizations/ModifyOrg
```

- Tipo: POST
- body: la organizacion entera
- Descripcion: una manera de modificar una organizacion

```
/organizations/SendMail
```

- Tipo: POST
- body: La organizacion entera
- Descripcion: le manda un mail al administrador de que una empresa se quiere registrar


### Users

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

```
/users/addTags
```
- Tipo: POST
- body:
```json
{
    "userid":"id del usuario"
    "tags":["tags que quieras agregar"]
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
        "userid":"id del usuario"
    }
```
- Descripcion: borrar un usuario

```
/users/Update
```
- Tipo: POST
- Body:el usuario entero
- Descripcion: actualizar un usuario


```
/users/Login
```
- Tipo: POST
- Body:
```json
{
    "telephone":"el numero de telefono"

}
```
- Descripcion: iniciar sesion de un usuario


