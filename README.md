## GIN MGO API

Ejemplo de una API REST utilizando [GO] 1.10 con [gin] como router y [mgo] como driver para mongoDB

####Endpoins disponibles

* `GET /products`
* `GET /products/:id`
* `POST /products`
* `PUT /products/:id`
* `DELETE /products/:id`

####Ejemplo de JSON
```json
{
  "name": "Fanta",
  "brand": "coca-cola",
  "variants": [
    {
      "description": "1,5 lts",
      "sku": "123456",
      "price": 1200,
      "stock": 5.0
    },
    {
      "description": "3 lts",
      "sku": "987654",
      "price": 2000,
      "stock": 7.0
    }
  ]
}
```
* Validaciones en progreso.
* Estructura de archivos en progreso.

[GO]:https://golang.org/
[gin]:https://gin-gonic.github.io/gin/
[mgo]:https://labix.org/mgo