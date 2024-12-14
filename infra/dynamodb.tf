# dynamodb.tf
resource "aws_dynamodb_table" "dynamodb-task-table" {
  name           = "tasks"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "id"
  range_key      = "title"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "title"
    type = "S"
  }

  tags = {
    Name        = "dynamodb-tasks-table"
    Environment = "dev"
  }
}
