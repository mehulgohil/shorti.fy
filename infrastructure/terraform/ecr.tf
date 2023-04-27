resource "aws_ecr_repository" "shortify-writer" {
  name                 = "shortify-writer"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_repository" "shortify-redirect" {
  name                 = "shortify-redirect"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_repository" "shortify-frontend" {
  name                 = "shortify-frontend"
  image_tag_mutability = "MUTABLE"
}
