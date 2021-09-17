provider "vault" {
  path = "./dist/example"
  private_key = file("private_key.pem")
}