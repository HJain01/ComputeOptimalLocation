using System.Collections.Generic;
using System.Threading.Tasks;

namespace ComputeOptimalLocationService.Core.Interfaces.Repositories
{
    public interface INotificationsRepository
    {
        Task<List<(string,string)>?> GetSubscriptions();
        Task Unsubscribe(string arn);
    }
}