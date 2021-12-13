locals {
  app_name       = "compute-optimal-location-service"
  app_port       = 4010
  vpc_id         = "vpc-03d76a285267c60b2"
  public_subnets = ["subnet-0a5efa8e3e982ea70", "subnet-0bc3dd4044093e68f", "subnet-0223c3e359d5294da"]
}

module "app_task" {
  source = "github.com/7factor/terraform-ecs-http-task"

  # Where we want to deploy the thing
  vpc_id       = local.vpc_id
  cluster_name = data.aws_ecs_cluster.cluster.cluster_name

  # Information about what we're deploying
  app_name         = local.app_name
  app_port         = local.app_port
  service_role_arn = "arn:aws:iam::${var.deploy_to_account}:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS"
  task_role_arn    = aws_iam_role.sns_publish.arn

  # Load balancers and health checking
  health_check_path    = "/status"
  health_check_matcher = "200,301,302"
  health_check_timeout = 15
  lb_cert_arn          = "arn:aws:acm:us-east-1:455764919342:certificate/648a6516-e8fd-469f-bb53-5febeef79f15"
  lb_public_subnets    = local.public_subnets
  cluster_lb_sg_id     = data.aws_security_group.access_sg.id
  desired_task_count   = 1
  is_lb_internal       = true
  idle_timeout         = 3600
  launch_type          = null
  ordered_placement_strategies = [
    {
      type  = "binpack", 
      field = "cpu" 
    }
  ]

  cpu    = 512
  memory = 512

  # Let's start with a hard coded container definition. A zero as host port means
  # we want an ephemeral range of ports.
  container_definition = <<EOF
[
  {
    "image": "${var.ecr_uri}:${var.ecr_tag}",
    "name": "${local.app_name}",
    "portMappings": [
      {
        "containerPort": ${local.app_port},
        "hostPort": 0
      }
    ],
    "environment": ${var.ecs_env_blob},
    "logConfiguration": {
      "logDriver": "awslogs",
      "options": {
          "awslogs-group": "ecs-logs",
          "awslogs-region": "${data.aws_region.current.name}",
          "awslogs-stream-prefix": "${local.app_name}",
          "awslogs-create-group": "true"
      }
    }
  }
]
EOF
}