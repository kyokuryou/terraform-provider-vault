resource "vault_secret_resource" "example" {
  name = "example"
  vars = {
    authc_type = "ssh"
    image_type = "linux"
    ip_address = "10.10.10.1"
    username = "exampleuser"
  }
}

resource "vault_secret_resource" "example1" {
  name = "example1"
  vars = {
    authc_type = "userpass"
    image_type = "windows"
    ip_address = "10.10.10.2"
    username = "exampleuser"
    password = "qwert@123456"
  }
}