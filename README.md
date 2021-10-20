# YoFioChallenger
Este es un proyecto para la empresa YoFio

The cloud version is in [Heroku](https://challenger-yofio.herokuapp.com/)

## Install

```sh
git clone https://github.com/jalozanot/yofioChallenger.git
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
| Content-Length | (Body length)    |

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
