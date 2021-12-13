using Microsoft.Extensions.Logging;

namespace ComputeOptimalLocationService.Core.Services
{
    public class ComputeOptimalLocationService
    {
       
        private readonly ILogger<ComputeOptimalLocationService> _logger;

        public ComputeOptimalLocationService(ILogger<ComputeOptimalLocationService> logger)
        {
           _logger = logger;
        }
    }
}