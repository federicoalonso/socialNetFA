
## *Universidad ORT Uruguay*

## *Facultad de Ingeniería*


# Fundamentos de Ingenier�a de Software
# Letra del Obligatorio 1


*Cristian Palma

Federico Alonso

Valeria Vera*


10 de octubre de 2020



# 1. INVESTIGACION

### 1.1 Actividades realizadas

Para comenzar, posterior a recibir la orden de trabajo y previo al contacto con el cliente, se comenz� a indagar sobre c�mo se deb�a abordar la investigaci�n. En ese momento se contaba con la informaci�n de que era un sistema para que los ni�os aprendan a tocar guitarra online, siendo nuestro cliente el profesor a cargo de impartir las clases. El cliente impartia clases presenciales a ni�os antes de comenzar la pandemia,  su hijo (quien nos contact�) es quien realizaba toda la administraci�n, y por motivos de la pandemia tuvieron que dejar de dar clases, perdiendo de esta forma una fuente de ingreso extra.

Resultante de la informaci�n recibida, se visualiza de forma r�pida contenido relacionado con la ense�anza de guitarra (a efectos de informarse sobre los t�rminos utilizados en las mismas: acordes, partituras, videos explicativos, afinaci�n, tipos de guitarra, etc.), y se prepara una entrevista primaria a los principales involucrados (el profesor de guitarra y su hijo), abordando con uno los requerimientos del �rea de educaci�n y con el otro los de la administraci�n del negocio.

Se decide realizar entrevistas debido a que el cliente est� bien definido, son personas en particular que desean que desarrollemos dicho sistema. Pudiendo de esta forma enfocarnos en los problemas que ellos desean que podamos resolver. Las mismas aparecen en los documentos [Entrevista 1](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-Entrevista1-2020) y [Entrevista 2](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob2-Entrevista2-2020)

Del primer contacto con el cliente, pudimos obtener datos de los adultos a cargo de algunos de los alumnos, para entrevistar a distintos tipos de personas con el fin de poder vislualizar los diferentes puntos de vista que tengan en cuanto a la creaci�n del sistema e incluirlos en el mismo. Se decide entrevistar a un padre de alumno [Entrevista 4](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-Entrevista4-2020), una madre [Entrevista 3](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-Entrevista3-2020) y un tutor legal, el abuelo de uno de los alumnos [Entrevista 5](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-Entrevista5-2020), debido a que fueron los que accedieron a realizar la entrevista. 
Adem�s como dato resultante de la entrevista con el hijo del cliente, se obtuvo una p�gina web de su inter�s, por lo que se aplic� el m�todo de ingenier�a inversa [Ingenier�a inversa](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-Ingenier�aInversa-2020) para as� poder obtener un mejor detalle de las expectativas del cliente. Extrayendo de la misma diferentes roles de usuarios e informaci�n que deber�a aparecer en los diferentes apartados del sistema.

### 1.2 Resultados obtenidos

Luego de realizar las tareas de investigaci�n se determina que el sistema de mayor conveniencia es el desarrollo de una aplicaci�n web, en la cual se carguen las lecciones del profesor y sus alumnos puedan ver las mismas desde su hogar. 
Para lo mismo se establecen tres user persona:
- Los ni�os: Acceder�n al sistema, y recibir�n las clases grabadas con anterioridad.
- Los padres: Ayudar�n a los ni�os a acceder al sistema y controlar su avance en las clases.
- El administrador: Se encargara de subir y organizar las clases y validar el acceso de los usuarios al sistema.

Se podr� mantener un contacto con el profesor mediante el medio que �l desee y en un horario establecido, pero el mismo ser� por fuera del sistema a desarrollar.
Los accesos al sistema ser�n gestionados por un administrador que recibir� los pagos por un medio externo a la aplicaci�n, y controlar� los mismos de forma manual, a efectos de tener flexibilidad con los distintos clientes. El sistema debe estar adaptado a dispositivos m�viles ya que los ni�os tienen un amplio dominio y uso de los mismos.

# 2. OBJETIVOS DEL SISTEMA

### 2.1 Introducci�n

El nombre del sistema ser� K.I.D.S Lessons, debido a que es el usado por el cliente como publicidad de las clases que �ste imparte.

### 2.2 Objetivos del producto

El objetivo del producto es permitir a los usuarios reproducir clases previamente grabadas de un profesor de guitarra para aprender a tocar el instrumento de forma intuitiva y amigable, haciendo posible la continuaci�n del aprendizaje en el marco de la pandemia COVID-19.

### 2.3 Alcance del producto

El software a desarrollar permite la baja, alta y visualizaci�n de videos, la alta, baja y asignaci�n de cursos a diferentes videos, contactarse con el docente para realizarle todo tipo de preguntas y un hist�rico de reproducciones.
Existen dos tipos de usuarios que utilizan el sistema para diferentes tareas, usuario del tipo alumno y usuario del tipo administrador.

### 2.4 Funciones principales

- Gesti�n de usuarios
    - Alta de usuarios del tipo alumno
    - Baja de usuarios del tipo alumno
    - Suspensi�n temporal de usuario
- Autenticaci�n de usuarios
    - Autenticar usuario de tipo administrador
    - Autenticar usuario de tipo alumno
- Gesti�n de cursos
    - Alta de un curso
    - Baja de un curso
    - Modificaci�n de curso
- Gesti�n de videos
    - Alta de un video
    - Baja de un baja
- Visualizaci�n del contenido
    - Consultar cursos disponibles
    - Seleccionar curso
    - Ver video
- Hist�rico de reproducciones
    - Visualizar hist�rico

### 2.5 Restricciones

- El sistema debe correr sobre Windows server 2019 64 bits.
- El motor de base de datos a utilizar ser� SQL server 2019.
- La aplicaci�n deber� correr en cualquier navegador que soporte HTML5.

# 3. DESCRIPCION DE USUARIOS

### 3.1 Usuarios de internet sin autenticar

Los usuarios sin autenticar son personas que a trav�s del marketing llegaron a la URL del sistema.
Tienen acceso a la p�gina de inicio, nosotros y contacto. 
En la p�gina de inicio se informa sobre la modalidad del curso.
En la p�gina de nosotros, se informa sobre la empresa y el profesor.
En la p�gina de contacto, se muestra un apartado de preguntas frecuentes y un formulario para enviar consultas administrativas las cuales ser�n recib�das por el usuario administrativo. Este las responde por fuera de la plataforma v�a email, por lo tanto antes de enviar la consulta debe escribir un mail a donde ser� respondida.

### 3.2 Alumnos autenticados en el sistema

![Im�gen usuario alumno](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/UserPersonaAlumno.jpg)

Los alumnos, espec�ficamente ni�os entre 8 a 12 a�os, se van a autenticar en el sistema con un usuario y contrase�a previamente suministrada por el administrador.
Una vez autenticados, van a poder navegar en la p�gina, seleccionar cursos de guitarra y mirar las clases online. 
Tendr�n una secci�n, donde podr�n enviar dudas al profesor. 
El sistema guardar� las clases que fueron reproducidas por el alumno, a efectos de mostrar un resumen del progreso.

### 3.3 Administrador

![Im�gen usuario administrador](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/UserPersonaAdministrador.jpg)

El administrador del sistema tiene las siguientes tareas:
- Crear cursos de guitarra.
- Subir videos a los cursos creados.
- Crear, Modificar y Eliminar usuarios con el rol de Alumno.
- Actualizar la informaci�n del pago de cuota de los usuarios Alumno.
- Obtener un reporte de los ingresos en un periodo seleccionado.

# 4. MODELO DE AN�LISIS

### 4.1 Diccionario

Se especif�ca la terminolog�a utilizada en el documento:

- *Usuario*: Persona que utiliza el sistema.
- *Usuario final*: Persona o personas que van a manipular de manera directa un producto de software
- *Usuario administrador del sitio*: Usuario encargado de la configuraci�n del sistema
- *Usuario autenticado*: Usuario del sistema que tiene una sesi�n activa.
- *Responsive*: Dise�o de p�gina web que se adapta de forma autom�tica a la resoluci�n de la pantalla donde est� siendo visualizado.

### 4.2 Modelo conceptual

##### 4.2.1 Diagrama de casos de uso

![Diagrama de casos de uso](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/DiagramaCasosDeUso.jpg)

##### 4.2.2 UML del dominio

![UML del dominio](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/UML.jpg)

# 5. Casos de uso

### 5.1 User Story 01 Alta Usuario

**ID:** #1	

**T�tulo:** Alta Alumno

**Narrativa:**
- Como administrador
- Quiero poder registrar a un nuevo alumno para que el mismo tenga acceso a utilizar las funcionalidades de la aplicaci�n.

**Criterios de aceptaci�n:**
- El email debe ser �nico, y v�lido.
- La contrase�a debe ser alfanum�rica de m�nimo 8 caracteres.
- Nombre completo y direcci�n deben de tener al menos 5 caracteres.
- Tel�fono de contacto debe ser un campo num�rico y por lo menos 8 cifras, 

### 5.2 Caso de Uso 01 Alta Usuario

![ ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/CasoDeUso01.jpg)

### 5.3 User Story 12 Ver Video

**ID:** #12

**T�tulo:** Ver Video

**Narrativa:**
- Como alumno
- Quiero poder continuar con la visualizaci�n de videos de un curso para poder seguir avanzando en el curso seleccionado.

Criterios de aceptaci�n:
- El alumno debe estar autenticado con usuario y contrase�a.
- El alumno debe estar habilitado.

### 5.4 Caso de Uso 12 Ver Video

![ ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/CasoDeUso12.jpg)

### 5.5 User Story 05 Alta Curso

**ID:** #5

**T�tulo:** Alta Curso

**Narrativa:**
- Como administrador
- Quiero poder dar de alta un nuevo curso para posteriormente cargarle videos y poder present�rselo a los alumnos.

**Criterios de aceptaci�n:**
- El nombre del curso debe ser �nico y poseer al menos 5 caracteres.
- El curso debe ser creado con el estado deshabilitado por defecto, ya que al momento de creaci�n no posee videos asociados. 

### 5.6 Caso de Uso 05 Alta Curso

![ ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/CasoDeUso05.jpg)

# 6. REQUERIMIENTOS FUNCIONALES

### 6.1 Gesti�n de usuarios

##### **RF01 - Alta de usuario**

**Rol:** Administrador

**Descripci�n:** El sistema deber� permitir crear un nuevo usuario del tipo alumno, utilizando un formulario de registro. Se deber� proporcionar nombre completo, direcci�n de correo electr�nico, domicilio, tel�fono de contacto y contrase�a. 

**Prioridad:** ALTA

**Ranking:** #2

**Especificaci�n:** CU-01


##### **RF02 - Baja de usuario**

**Rol:** Administrador

**Descripci�n:** El sistema deber� permitir la eliminaci�n permanente de usuarios del tipo alumno en el sistema. 

**Prioridad:** BAJA

**Ranking:** #13

**Especificaci�n:** CU-02


##### **RF03 - Suspensi�n temporal de usuario**

**Rol:** Administrador

**Descripci�n:** El sistema deber� permitir la suspensi�n de un usuario temporalmente, impidiendo el acceso al sistema, sin eliminar sus datos e hist�rico de reproducciones.

**Prioridad:** MEDIA

**Ranking:** #8

**Especificaci�n:** CU-03


### 6.2 Autenticaci�n de usuarios

##### **RF04 - Login**

**Rol:** Usuario no Autenticado

**Descripci�n:** El sistema deber� detectar autom�ticamente solo con el nombre de usuario y contrase�a que el usuario es del tipo administrador o tipo alumno. Al ingresar habilitar� al usuario a acceder a las secciones correspondientes a su rol.

**Prioridad:** ALTA

**Ranking:** #5

**Especificaci�n:** CU-04


### 6.3 Gesti�n de cursos

##### **RF05 - Alta Curso**

**Rol:** Administrador.

**Descripci�n:** El sistema deber� contar con una funcionalidad que permita dar alta de nuevos cursos, ingresando nombre del curso (el cual debe ser �nico), descripci�n y nivel (el nivel se selecciona de un combo box que brinda la opci�n: Principiante, intermedio y avanzado).

**Prioridad:** ALTA

**Ranking:** #1

**Especificaci�n:** CU-05


##### **RF06 � Baja Curso**

**Rol:** Administrador.

**Descripci�n:** El sistema deber� contar con una funcionalidad que permita eliminar los cursos existentes. Se debe seleccionar el curso a eliminar y luego presionar Eliminar, se solicitar� confirmar la operaci�n.

**Comentarios:** No es posible eliminar un curso si el mismo tiene videos asociados.

**Prioridad:** MEDIA

**Ranking:** #12

**Especificaci�n:** CU-06


##### **RF07 � Modificar Curso**

**Rol:** Administrador.

**Descripci�n:** El sistema deber� contar con una funcionalidad que permita modificar los cursos. El usuario administrador selecciona un curso. El sistema ofrece modificar el estado del curso, el nombre, la descripci�n y el nivel del curso.

**Comentarios:** No es posible habilitar un curso que no tenga videos asociados. No es posible modificar el nombre de un curso si el nombre que se desea ya existe. 

**Prioridad:** ALTA

**Ranking:** #6

**Especificaci�n:** CU-07


### 6.4 Gesti�n de videos

##### **RF08 � Alta Video**

**Rol:** Administrador.

**Descripci�n:** El sistema ser� capaz de almacenar videos. El usuario hace click en el bot�n subir video y el sistema abre un cuadro de dialogo en donde el usuario selecciona un video de su ordenador. 

**Comentario:** El bot�n de subir video est� disponible en la ventana de Modificar Curso del curso seleccionado. Cuando el video termina de subir, aparece visible en la ventana modificar curso.

**Prioridad:** ALTA

**Ranking:** #4

**Especificaci�n:** CU-08


##### **RF09 � Baja Video**

**Rol:** Administrador.

**Descripci�n:** El sistema ser� capaz de permitir al usuario administrador eliminar videos asociados a un curso. El usuario selecciona el bot�n eliminar video. 

**Comentario:** El bot�n de eliminar video est� disponible en la ventana de Modificar Curso del curso seleccionado. El usuario debe confirmar la eliminaci�n.

**Prioridad:** MEDIA

**Ranking:** #11

**Especificaci�n:** CU-09


### 6.5 Visualizaci�n de contenido

##### **RF10 � Consulta de cursos**

**Rol:** Alumno.

**Descripci�n:** El sistema deber� permitir la consulta de los cursos habilitados en el mismo.

**Prioridad:** ALTA

**Ranking:** #10

**Especificaci�n:** CU-10


##### **RF11 � Seleccionar curso**

**Rol:** Alumno.

**Descripci�n:** El sistema deber� permitir seleccionar un curso. Luego de que el usuario selecciona el mimo, se mostrara en otra ventana los videos asociados al curso seleccionado.

**Prioridad:** ALTA

**Ranking:** #7

**Especificaci�n:** CU-11


##### **RF12 � Ver video**

**Rol:** Alumno.

**Descripci�n:** El sistema deber� permitir reproducir, pausar, adelantar, atrasar, subir y bajar volumen de videos subidos en la plataforma.

**Comentario:** Para seleccionar un video a ser reproducido, se debe estar dentro de un curso seleccionado.

**Prioridad:** ALTA

**Ranking:** #3

**Especificaci�n:** CU-12


### 6.6 Hist�rico de reproducciones

##### **RF13 � Ver -Historial**

**Rol:** Alumno.

**Descripci�n:** El sistema deber� permitir reproducir, pausar, adelantar, atrasar, subir y bajar volumen de videos subidos en la plataforma.

**Comentario:** Para seleccionar un video a ser reproducido, se debe estar dentro de un curso seleccionado.

**Prioridad:** ALTA

**Ranking:** #9

**Especificaci�n:** CU-13


# 7. REQUERIMIENTOS NO FUNCIONALES


##### **RNF01 - Dise�o de Interfaz de Usuario**

**Descripci�n:** La interfaz de usuario debe ser responsive (se debe adaptar a dispositivos web y m�viles).

**Categor�a:** Usabilidad.

**Prioridad:** Alta 

**Ranking:** #1


##### **RNF02 - Informe de Errores**

**Descripci�n:** El sistema deber� informar cada vez que se cometa un error en el flujo normal del sistema a trav�s de mensajes informativos al usuario final. A fallas distintas debe dar mensajes de errores espec�ficos a efecto de ser reconocido por los usuarios.

**Categor�a:** Usabilidad.

**Prioridad:** Alta

**Ranking:** #2


##### **RNF03 - Servicio de Hosting**

**Descripci�n:** Se utilizar� el servicio de hosting de GoDaddy.

**Categor�a:** Software

**Prioridad:** Baja 

**Ranking:** #6


##### **RNF04 - Tiempos de Respuesta**

**Descripci�n:** Los tiempos promedios de respuesta de las operaciones que realicen los usuarios contra el servidor no deber�n exceder los 3 segundos.

**Categor�a:** Calidad

**Prioridad:** Media 

**Ranking:** #3


##### **RNF05 - Idioma**

**Descripci�n:** El idioma del sistema ser� espa�ol (Latinoamericano).

**Categor�a:** Usabilidad

**Prioridad:** Baja

**Ranking:** #5


##### **RNF6 - Sistema**

**Descripci�n:** El sistema debe correr sobre Windows 10 o superior.

**Categor�a:** Funcionalidad

**Prioridad:** Media

**Ranking:** #4

# 8. VERIFICACION DE REQUERIMIENTOS

Luego de especificar la totalidad de los requerimientos funcionales y no funcionales del sistema, se procedi� a verificar la totalidad de los mismos chequeando que cada uno de ellos cumpla con las caracter�sticas que a continuaci�n se detallan seg�n el est�ndar SRS (Software Requirements Specification).

# 9. VALIDACION DE REQUERIMIENTOS 

Para la validaci�n de requerimientos se utiliz� la t�cnica de prototipado, la cual fue realizada mediante 3 reuniones de forma presencial con el cliente.

### 9.1 Primera reuni�n

Se le presenta al cliente dos alternativas de interfaz de usuario, se pueden los prototipos completos en el documento: [Documento Primera Reuni�n ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PrototipadoPrimeraReunion-2020) 

![Ejemplo prototipo 1](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/EjemploPrototipo1.jpg)

![Ejemplo prototipo 2](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/EjemploPrototipo2.jpg)

##### Retroalimentaci�n de la primera reuni�n:

El cliente se inclin� por el prototipo 2 y realiz� para cada imagen del documento APV-ob1-PrototipadoPrimeraReunion-2020 las siguientes observaciones:

- Figura 16: Se cambia la columna habilitado,para que muestre en cada fila un bot�n que al presionarlo cambia el estado del alumno entre habilitado y deshabilitado.
- Figura 18: A la descripci�n del video, se le agrega el curso al cual pertenece y el nivel de dicho curso.
- Figura 19: Se agrega el curso a la cual pertenece la clase.
- El cliente, adem�s de las p�ginas mostradas en el prototipo solicit� que se agregue una nueva p�gina en la cual se muestren todos los cursos de la plataforma.

### 9.2 Segunda reuni�n

Se agregan, adem�s de los cambios solicitados, los modelos de formularios para simular la implementaci�n de los requerimientos. El estudio detallado se encuentra en el archivo: [Documento Segunda Reuni�n ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PrototipadoSegundaReunion-2020)

![Ejemplos formularios ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/Formularios.jpg) 

![Selecci�n de cursos ](https://github.com/ORT-FIS-202008/ob1-alonso-palma-vera-1/blob/master/APV-ob1-PNGs/SeleccionDeCursos.jpg) 

##### Retroalimentaci�n de la segunda reuni�n:

- En los formularios cambiar el texto del bot�n Close por Cerrar.
- En el formulario 2, cambiar el dise�o del video en miniatura junto al bot�n de eliminar. Al cliente no le pareci� est�tico.

Se obtuvo un feedback suficiente para dise�ar un prototipo evolutivo en HTML, el cual ser� utilizado en una tercera reuni�n.

### 9.3 Tercera reuni�n

Se le presenta al cliente un prototipo evolutivo en HTML, que le permiti� ver al cliente como funcionan los requisitos solicitados al implementarlos, obteniendo de esta manera una validaci�n m�s precisa.
El cliente experiment� con el prototipo (c�digo fuente en carpeta: APV-ob1-Prototipos-HTML-2020) y se obtuvieron las observaciones que a continuaci�n se detallan:

- Experimentando la navegaci�n como alumno, se solicita que el alumno pueda acceder de forma m�s r�pida al curso que viene siguiendo.

Se corrigi� la observaci�n, agregando un bot�n de accedo directo a la �ltima lecci�n, en la p�gina de inicio del alumno.

Tras la aceptaci�n del cliente de esta ultima modificaci�n, se dio por finalizada la etapa de validaci�n.


### 10.	Reflexi�n



