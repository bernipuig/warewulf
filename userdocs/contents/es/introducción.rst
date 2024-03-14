============
Introducción
============

La visión de Warewulf
=====================

Warewulf ha tenido varias iteraciones en los últimos 20 años, pero 
sus principios de diseño siempre han sido los mismos: un sistema sencillo, escalable,
sin estado (aunque algunas versiones podían aprovisionar con estado), y 
muy flexible para todo tipo de clústeres.
Esta es una visión general de cómo funciona Warewulf.

Acerca de Warewulf
==================

Warewulf es una plataforma de aprovisionamiento de sistemas operativos para Linux
diseñada para producir implementaciones de clúster seguras, escalables y listas para
usar, que mantienen la flexibilidad y la simplicidad. 

Desde su lanzamiento inicial en 2001, Warewulf se ha convertido en el sistema de 
aprovisionamiento de código abierto e independiente del proveedor más popular dentro 
de la comunidad global de HPC. Warewulf es conocido por su enorme escalabilidad y su 
sencilla gestión del aprovisionamiento sin estado (disco opcional). 

Warewulf aprovecha un modelo administrativo sencillo que centraliza la administración 
en torno a imágenes de nodos virtuales que se utilizan para el aprovisionamiento de 
los nodos del clúster. Esto significa que puede tener cientos o miles de nodos de clúster 
arrancando y funcionando con la misma imagen idéntica de sistema de archivos de nodo virtual. 
A partir de Warewulf v4, la imagen de nodo virtual es una imagen de contenedor estándar, lo que
significa que todos los recursos informáticos de un clúster se pueden gestionar utilizando 
cualquier herramienta de contenedor existente y/o canalizaciones de CI que se estén utilizando. 
Esto puede ser tan simple como DockerHub o su propia infraestructura privada de GitLab CI.

Con esta arquitectura, Warewulf combina lo mejor de la computación de alto rendimiento (HPC), 
la nube, la Hyperscale y los principios de despliegue empresarial para crear y mantener grandes
clústeres escalables sin estado.

Aunque las raíces de Warewulf se encuentran en HPC, se ha utilizado para muchos tipos diferentes 
de tareas y casos de uso. Todo, desde servidores web en clúster, a granjas de renderizado, e incluso
Kubernetes y despliegues en la nube, Warewulf aporta beneficios en la experiencia de la gestión general del sistema
operativo a escala.

Características 
===============

* **Ligero**: Warewulf aprovisiona imágenes de sistemas operativos sin estado y luego se quita de en medio.
  No debe haber dependencias del sistema subyacente ni cambios en los sistemas operativos de los nodos de 
  clúster aprovisionados. 

* **Sencillo**: Warewulf es utilizado por aficionados, investigadores, científicos, ingenieros y administradores
  de sistemas porque es fácil, ligero y sencillo. 

* **Flexible**: Warewulf es muy flexible y puede responder a las necesidades de cualquier entorno: desde un
  laboratorio informático con estaciones de trabajo gráficas hasta clústeres de sobremesa, pasando por centros de
  supercomputación masivos que proporcionan capacidades HPC tradicionales a miles de usuarios. 

* **Agnóstico**: Desde la distribución de Linux elegida hasta el hardware subyacente, Warewulf es agnóstico y
  compatible con los estándares. Desde ARM a x86, Atos a Dell, Debian, SUSE, Rocky, CentOS y RHEL, 
  Warewulf puede con todo. 

* **Seguro**: Warewulf es el único sistema de aprovisionamiento sin estado compatible con SELinux. 
  Sólo tiene que habilitar su nodo contenedor de sistema operativo para que soporte SELinux y Warewulf hará el resto. 

* **Código abierto**: Durante los últimos 20 años, Warewulf ha seguido siendo de código abierto y continúa siendo
  el estándar de oro para el aprovisionamiento de clústeres.
