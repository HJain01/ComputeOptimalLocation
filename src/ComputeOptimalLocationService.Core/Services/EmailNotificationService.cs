using System;
using System.Text;
using System.Threading.Tasks;
using ComputeOptimalLocationService.Core.Configuration;
using ComputeOptimalLocationService.Core.Interfaces.Repositories;

namespace ComputeOptimalLocationService.Core.Services
{
    public class EmailNotificationService
    {
        private readonly INotificationsRepository _notificationsRepository;
        private readonly BaseConfiguration _configuration;
        
        public EmailNotificationService(
            INotificationsRepository notificationsRepository, 
            BaseConfiguration configuration)
        {
            _notificationsRepository = notificationsRepository;
            _configuration = configuration;
        }

        public async Task Unsubscribe(string arn)
        {
            arn = Encoding.ASCII.GetString(Convert.FromBase64String(arn));
            await _notificationsRepository.Unsubscribe(arn);
        }
    }
}