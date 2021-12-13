provider "aws" {
  version = "~> 3.32"

  assume_role {
    role_arn = "arn:aws:iam::${var.deploy_to_account}:role/AveannaContinuousDelivery"
  }
}

data "aws_region" "current" {}

# Use data to safeguard against a non-existing cluster during
# the CI step.
data "aws_ecs_cluster" "cluster" {
  cluster_name = "rc-cluster"
}

# We name these things in a repeatable way, take advantage of this.
data "aws_security_group" "access_sg" {
  name = "${data.aws_ecs_cluster.cluster.cluster_name}-ecs-access-sg"
}

data "aws_security_group" "ecs_sg" {
  name = "${data.aws_ecs_cluster.cluster.cluster_name}-ecs-sg"
}