output "lb_hostname" {
  value = module.app_task.lb_hostname
}

output "health_check_path" {
  value = module.app_task.health_check_path
}

output "deployed_container_tag" {
  value = var.ecr_tag
}