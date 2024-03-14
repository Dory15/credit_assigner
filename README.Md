Credit Assigner
-----

Para correr el sistema de la manera más sencilla.
```shell
docker compose up 
# O 
docker-compose up
```

Las claves para conectarse a la base de datos estan puestas para evitar
problemas al correrlo.

Despues acceder a la base de datos desde un cliente para SQL Server
con el nombre `credit_assigner`.

Y por último correr el contenido del archivo `table.sql` seguido por el de `sp.sql` ya que utiliza procedimientos almacenados.

Las pruebas unitarias se encuentran en la carpeta de core, y las pruebas de integración se encuentran en Postman en la colección adjuntada