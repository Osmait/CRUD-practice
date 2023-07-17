#!/bin/bash

# Obtener la lista de imágenes Docker sin nombre de repositorio (<none>)
images=$(docker images --format "{{.ID}} {{.Repository}}:{{.Tag}}")

# Recorrer cada línea de la lista de imágenes
while read -r image; do
  # Obtener el ID y el nombre de la imagen
  image_id=$(echo "$image" | awk '{print $1}')
  repository=$(echo "$image" | awk '{print $2}')

  # Verificar si el nombre de repositorio es "<none>"
  if [[ "$repository" == "<none>" ]]; then
    # Eliminar la imagen
    echo "Eliminando imagen: $image_id"
    docker rmi "$image_id" -f
  fi
done <<< "$images"
