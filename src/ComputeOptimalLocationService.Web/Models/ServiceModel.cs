using System.Net;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace ComputeOptimalLocationService.Web.Models
{
    public class ServiceModel
    {
        public ServiceModel(string name, Task<HttpStatusCode?> task)
        {
            Name = name;
            Task = task;
        }

        public string Name { get; set; }
        public int StatusCode { get; set; }
        public string? Status { get; set; }

        [JsonIgnore]
        public Task<HttpStatusCode?> Task { get; set; }
    }
}