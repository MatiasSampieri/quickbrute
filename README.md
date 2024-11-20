# Quick Distributed Brute Force

Una herramienta simple para realizar ataques de fuerza bruta distribuidos.

## Flags
`-b <cantidad>` (500) Cantidad de peticiones en paralelo (reducir si el programa crashea)

`-o <nombre archivo>` (out.log) Archivo de salida de peticiones que cumplen con el criterio 

`-p <puerto>` (7575) Puerto para la comunicacion entre helpers e instancia principal

## Uso 

`qdbf [<flags>] [<config.yml>]`

Si no se especifica archivo de configuracion se inicia en modo "helper" donde se espera a que una insancia principal se comunique y le envie lo que tiene que hacer.

## Helpers

La herramienta se puede usar de forma 'standalone' o distribuida. En el archivo de configuracion se pueden especificar de 0 a muchas direcciones IP de maquinas que ayudaran en el ataque. Estas maquinas seran llamadas "helpers".

## Estructura de archivo de configuracion (config.yml)

```yaml
request:
  method: <GET | POST | PUT | ...>
  url: http://test/
  params: # parametros query en formato clave: valor
    username: Test
    password: bee
    id: "$id$" # las variables se definen entre '$'
  headers:
    cookie: "sessionid=ab3235bac5"
  body: '{"user":"test", "tst":55}'

params: # valores que asumiran las variables
  id: # nombre de variable
    type: RANGE # tipos: RANGE, DICT, FILE 
    from: 1
    to: 10000

criteria: 
  type: <STOP | LOG> # frenar o guardar y seguir
  response:
    status: 200
    headers:
      cookie: "sesion=sa4d65541a"
    body: 'sjdhaksjdh' # soporta REGEX

helpers: # otras instancia que ayudaran
  - 192.168.1.1
  - 123.45.67.8
```

## Request

Se especifica todo lo necesario para realizar la peticion: URL, metodo, headers, body y parametros de query (en formato `<nombre param>: valor`). Los parametros query, el body y los headers son opcionales.

```yaml
request:
  method: <GET | POST | PUT | ...>
  url: <url>
  params:
    nombre_param: <valor>
  headers:
    nombre_header: <valor>
  body: <contenido del body>
```

## Variables/Params

Se definen entre '$' (ej. `$nombre$`) y pueden estar en el body, algun header, el path de la URL, o algun parametro de query. En la seccion de `params` se define que valores tendra la variable en cada peticion. No es necesario que se utilice la variable, si quieres realizar 100 peticiones pero no quieres que cambie nada entre ellas simplemente crea una variable RANGE de 1 al 100 y no la uses en ninguna parte de la peticio.

Las variables RANGE se incrementan de uno en uno en cada peticion desde `from:` hasta `to:`.

```yaml
nombre_var:
    type: RANGE
    from: <inicio>
    to: <fin>
```

Las variables DICT toman sus valores de una lista de valores especificada en el mismo archivo de configuracion con `dict:`

```yaml
nombre_var:
    type: DICT
    dict: ['a', 'b', 'c']
```

Las variables FILE toman sus valores de una lista de valores que se encuentra en un archivo, cada linea de archivo se toma como un valor de la lista. Se especifica el nombre del archivo con `file:`

```yaml
nombre_var:
    type: FILE
    file: <nombre archivo>
```

## Criterio de corte [opcional]

Indica que debe contener una respuesta para ser aceptada. Se puede especificar contenido de encabecado con `headers:` (igual formato que la peticion), codigo de status con `status:` y contenido del body con `body:` especificando una expresion regex. 

```yaml
criteria:
    type: <STOP | LOG>
    response:
        status: <codigo status>
        headers:
            nombre_header: <valor>
        body: <regex del contenido esperado>
```

No es necesario especificar todos estos campos, por ejemplo si quiero frenar el ataque cuando el codigo de status es 200 seria:

```yaml
criteria:
    type: STOP
    response:
        status: 200
```

Se puede especificar si cuando una respuesta cumpla con el cirterio se frene el ataque `type: STOP` o solamente se la guarde y el ataque continua `type: LOG`. En ambos casos los resultados seran guardados en el archivo de salida especificado con el flag `-o` (out.log por efecto), si se estan usando helpers este archivo estara solo en la instancia principal pero logueara los resultados de todas las demas instancias.