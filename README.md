
# GoDBusNotifier

Una aplicación de línea de comandos para enviar notificaciones de escritorio
usando D-Bus en Go.

## Uso

```bash
go run main.go [flags] <título> <cuerpo>
```

### Argumentos

- `título`: El título de la notificación.
- `cuerpo`: El cuerpo de la notificación.

### Flags

- `-urgency`: El nivel de urgencia de la notificación (`low`, `normal`,
  `critical`). Por defecto es `normal`.
- `-icon`: El directorio a un icono para mostrar en la notificación.

## Ejemplos

### Notificación básica

```bash
go run main.go "Hola" "Este es un mensaje."
```

### Notificación crítica

```bash
go run main.go -urgency=critical "¡Alerta!" "Este es un mensaje crítico."
```

### Notificación con un icono

```bash
go run main.go -icon=/usr/share/icons/Adwaita/32x32/actions/go-home-symbolic.symbolic.png "Icono" "Esta notificación tiene un icono."
```
