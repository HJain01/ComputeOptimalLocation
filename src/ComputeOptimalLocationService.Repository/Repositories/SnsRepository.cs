using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Amazon;
using Amazon.SimpleNotificationService;
using ComputeOptimalLocationService.Core.Configuration;
using ComputeOptimalLocationService.Core.Interfaces.Repositories;

namespace ComputeOptimalLocationService.Repository.Repositories
{
    public class SnsRepository : INotificationsRepository
    {
        private readonly BaseConfiguration _configuration;

        public SnsRepository(BaseConfiguration configuration)
        {
            _configuration = configuration;
        }
        
        public async Task<List<(string,string)>?> GetSubscriptions()
        {
            using (var snsClient = CreateSnsClient())
            {
                var subscriptions = await snsClient.ListSubscriptionsByTopicAsync(_configuration.Emails.SnsTopicArn);
                return subscriptions.Subscriptions.Select(p => (p.Endpoint, p.SubscriptionArn)).ToList();
            }
        }

        public async Task Unsubscribe(string arn)
        {
            using (var snsClient = CreateSnsClient())
            {
                await snsClient.UnsubscribeAsync(arn);
            }
        }

        private AmazonSimpleNotificationServiceClient CreateSnsClient()
        {
            return new AmazonSimpleNotificationServiceClient(RegionEndpoint.GetBySystemName(_configuration.Emails.Region));
        }
    }
}