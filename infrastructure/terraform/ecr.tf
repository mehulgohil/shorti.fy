resource "aws_ecr_repository" "shortify-writer" {
  name                 = "shortify-writer"
  image_tag_mutability = "IMMUTABLE"
}

resource "aws_ecr_repository" "shortify-redirect" {
  name                 = "shortify-redirect"
  image_tag_mutability = "IMMUTABLE"
}
