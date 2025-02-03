terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

# Pulls the image
resource "docker_image" "nginx" {
  name = "rmnobarra/nginx:green"
}

# Create a container
resource "docker_container" "nginx_container" {
  image = docker_image.nginx.image_id
  name = "nginx_container-${terraform.workspace}"

  ports {
    internal = 80
    external = 80
  }
}