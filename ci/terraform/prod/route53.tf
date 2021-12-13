data "aws_route53_zone" "root_zone" {
  name = "aveanna.io"
}

resource "aws_route53_record" "staging" {
  type    = "A"
  name    = "compute-optimal-location.aveanna.io"
  zone_id = data.aws_route53_zone.root_zone.zone_id

  alias {
    name                   = module.app_task.lb_hostname
    zone_id                = module.app_task.lb_zone_id
    evaluate_target_health = false
  }
}