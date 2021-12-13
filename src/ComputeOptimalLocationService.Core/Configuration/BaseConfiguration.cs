namespace ComputeOptimalLocationService.Core.Configuration
{
    public class BaseConfiguration
    {
        public BaseConfiguration()
        {
            BaseUrl = string.Empty;
        }
        
        public int Port { get; set; }
        
        public string BaseUrl { get; set; }
        
        public EmailsConfiguration Emails { get; set; }
        
    }
}