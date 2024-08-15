# crear un proyecto
ng new nombre-app --skip-git

<!-- iniciar el server -->
ng serve

<!-- Crear componentes -->
ng generate component register






<!-- Crear proyectos -->
ng new <name-project> [options] / ng n <name-project> [options]
ng new curso-angular17 --style=scss

<!-- Ejecutar nuestro proyecto localmente: -->
ng serve <name-project> [options] / ng s <name-project> [options]

<!-- Para ejecutar el Proyecto debes estar dentro de la carpeta: -->
ng s

<!-- Generar Elementos de Angular -->
ng generate <schematic> / ng g <schematic>
ng generate component <name> [options] / ng g component <name> [options]
ng generate service <name> [options] / ng g s <name> [options]
ng generate interceptor <name> [options] / ng g i <name> [options]
ng generate directive <name> [options] / ng g d <name> [options]
<!-- Ejemplos: -->
ng generate component login-page --standalone / ng g component login-page –standalone
ng generate service login --skip-tests/ ng g component login-page --skip-tests

<!-- Compilar / Transpilar nuestro Proyecto -->
ng build <project> [options] / ng b <project> [options]
ng build / ng b
<!-- Compilar / Transpilar usando un environment -->
ng build --configuration <name-environment> / ng b --configuration <name-environment>
ng build --configuration develop / ng b --configuration develop