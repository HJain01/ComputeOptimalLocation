variable "deploy_to_account" {
  description = "The account to deploy into. Passed in from concourse."
}

variable "ecr_uri" {
  description = "URI to the repo for the image to pull and deploy. Passed into the container definition."
}

variable "ecr_tag" {
  description = "Tag of the container to pull. Passed in from concourse."
}

variable "ecs_env_blob" {
  default     = "[]"
  description = "JSON blob name, value pairs to be passed to ECS environment block. This is created in the pipeline."
}