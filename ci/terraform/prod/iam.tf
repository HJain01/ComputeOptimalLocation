data "aws_iam_policy_document" "ecs_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "sns_publish" {
  name                = "dev-compute-optimal-location-service"
  assume_role_policy  = data.aws_iam_policy_document.ecs_assume_role_policy.json
  managed_policy_arns = [aws_iam_policy.sns_publish.arn]
}

resource "aws_iam_policy" "sns_publish" {
  name = "dev-compute-optimal-location-service"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "sns:*",
          "s3:*",
          "ses:*"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}