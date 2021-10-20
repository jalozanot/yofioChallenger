# YoFioChallenger
Este es un proyecto para la empresa YoFio

la aplicacion esta desplegada en la plataforma de heroku para su respectiva pruebas.

Curl para el servicio post /credit-assignment

```
curl --location --request POST 'https://challenger-yofio.herokuapp.com/credit-assignment' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw ' {
"investment":1500
}'
```
-------------------------------------------------------------------------------------

Curl para el servicio post /statistics

```
curl --location --request POST 'https://challenger-yofio.herokuapp.com/statistics' \
--header 'Accept: application/json' \
--data-raw ''
```



## Install

```sh
git clone https://github.com/jalozanot/yofioChallenger.git
```

subir el proyecto a cualquier IDE como visual studio code o intellij
ejecutar el comando para descargar dependencias.

go mod tidy 

crear una base de datos posgrest

### PostgreSQL

- host := "localhost"
- port := "5432"
- dbname := "test"
- user := "yofio"
- password := "yofio"

**Create a user**
```sh
    sudo -u postgres psql -p 5432 -c "CREATE USER yofio WITH PASSWORD 'yofio';"
```

**Create a database**
```sh
    sudo -u postgres psql -p 5432 -c "CREATE DATABASE test WITH OWNER yofio;"
```


## End Points

### POST /create-assignment

Crea una asignación basada en la inversión proporcionada. si la inversion no es correcta se guarda la peticion y responde bad request.


### Request

**Headers**

| Name         | Value            |
| ------------ | ---------------- |
| Content-Type | application/json |

**Body**

```json
{
    "investment": 1500
}
```


### Response

**Headers**

| Name           | Value            |
| -------------- | ---------------- |
| Content-Type   | application/json |


**Body**

```json
// cuando la inversion es correcta
{
{
  "credit_type_300": 1,
  "credit_type_500": 1,
  "credit_type_700": 1
}
}

// cuando la inversion es incorrecta
{
"bad request"
}

```


### POST /statistics

Obtiene las estadísticas de las asignaciones.

### Response

**Headers**

| Name           | Value            |
| -------------- | ---------------- |
| Content-Type   | application/json |


**Body**

```json
// cuando la consulta es exitosa
{
  "total_asignments_done": 5,
  "total_asignments_success": 4,
  "total_asignments_fail": 1,
  "avg_investment_success": 1500,
  "avg_investment_fail": 8749
}


```
