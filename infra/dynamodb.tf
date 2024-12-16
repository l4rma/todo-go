# dynamodb.tf
resource "aws_dynamodb_table" "dynamodb-task-table" {
  name           = "tasks"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name        = "dynamodb-tasks-table"
    Environment = "dev"
  }
}
