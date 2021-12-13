namespace ComputeOptimalLocationService.Core.Configuration
{
    public class EmailsConfiguration
    {
        public EmailsConfiguration()
        {
            Region = string.Empty;
            SnsTopicArn = string.Empty;
            FromEmail = string.Empty;
        }

        public string? Region { get; set; }
        public string? SnsTopicArn { get; set; }
        
        public string? FromEmail { get; set; }
        public string? Bucket { get; set; }
    }
}