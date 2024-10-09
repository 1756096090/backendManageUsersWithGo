## Configuración de la Base de Datos

En el archivo `config/config.go`, se encuentra toda la lógica necesaria para conectar nuestra aplicación a una base de datos MongoDB remota. A continuación, se describen los componentes clave:

- **Conexión a MongoDB**: La aplicación se conecta a la base de datos utilizando las credenciales y la cadena de conexión proporcionadas. Se utiliza el paquete `mongo-driver` para gestionar la conexión y realizar operaciones con la base de datos.

- **Modelo**: La capa de modelo define las estructuras de datos que representan los documentos en la base de datos. Aquí se especifican las propiedades y tipos de datos que se utilizarán en las operaciones de la aplicación.

- **Estructura de Rutas**: La aplicación cuenta con un sistema de rutas que dirigen las solicitudes a los controladores correspondientes. Estos controladores son responsables de manejar las peticiones y de interactuar con los servicios.

- **Servicios**: Los servicios actúan como intermediarios entre los controladores y la lógica del negocio. Aquí se implementan las operaciones específicas que se desean realizar con los datos.

- **Repositorio**: La capa de repositorio se encarga de la lógica relacionada con la base de datos, facilitando las interacciones necesarias para acceder y manipular los datos.

Con esta estructura, la aplicación puede gestionar de manera eficiente las operaciones sobre la base de datos, garantizando una separación clara de responsabilidades.
