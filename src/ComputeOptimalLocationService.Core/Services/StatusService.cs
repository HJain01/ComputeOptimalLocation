using System;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;

namespace ComputeOptimalLocationService.Core.Services
{
    public interface IStatusService
    {
        Task<bool> IsHealthy();
    }
    public class StatusService : IStatusService
    {
        private readonly ILogger<StatusService> _logger;
        
        public StatusService(ILogger<StatusService> logger)
        {
            _logger = logger;
        }

        public async Task<bool> IsHealthy()
        {
           return true;
        }
    }
}