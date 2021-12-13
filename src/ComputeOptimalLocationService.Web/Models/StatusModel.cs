using System.Collections.Generic;

namespace ComputeOptimalLocationService.Web.Models
{
    public class StatusModel
    {
        public StatusModel()
        {
            Name = string.Empty;
            Environment = string.Empty;
            StartedTime = string.Empty;
            Language = string.Empty;
            Services = new List<ServiceModel>();
        }

        public string Name { get; set; }
        public string Environment { get; set; }
        public string StartedTime { get; set; }
        public string Language { get; set; }
        public List<ServiceModel> Services { get; set; }
    }
}