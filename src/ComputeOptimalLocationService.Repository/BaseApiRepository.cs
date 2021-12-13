using System;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;

namespace ComputeOptimalLocationService.Repository
{
    public class BaseApiRepository
    {
        private readonly IHttpClientFactory _httpClient;
        private readonly ILogger<BaseApiRepository> _logger;
        protected virtual string BaseUrl { get; set; }

        public BaseApiRepository(IHttpClientFactory httpClient, 
                                 ILogger<BaseApiRepository> logger)
        {
            _httpClient = httpClient;
            _logger = logger;

            BaseUrl = string.Empty;
        }
        
        public async Task<T?> Get<T>(string relativeUrl, int timeout = 60000)
        {
            var response = await Get(BaseUrl, relativeUrl, timeout);
            var output = await response.Content.ReadAsStringAsync();
            if (!response.IsSuccessStatusCode)
            {
                throw new HttpRequestException($"Status Code: {response.StatusCode}, Content: {output}");
            }
            
            return JsonConvert.DeserializeObject<T>(output, new JsonSerializerSettings
            {
                NullValueHandling = NullValueHandling.Ignore
            });
        }
        
        public async Task<HttpResponseMessage?> Post<T>(string relativeUrl, T body, int timeout = 60000)
        {
            return await Post(BaseUrl, relativeUrl, body, timeout);
        }
        
        public async Task<T1?> Post<T,T1>(string relativeUrl, T body, int timeout = 60000)
        {
            var response = await Post(BaseUrl, relativeUrl, body, timeout);
            var output = await response.Content.ReadAsStringAsync();

            return JsonConvert.DeserializeObject<T1>(output, new JsonSerializerSettings
            {
                NullValueHandling = NullValueHandling.Ignore
            });
        }

        protected async Task<HttpResponseMessage> Get(string absoluteUrl, string relativeUrl, int timeout = 60000)
        {
            var client = _httpClient.CreateClient();
            client.Timeout = TimeSpan.FromMilliseconds(timeout);
            var url = new Uri(new Uri(absoluteUrl), relativeUrl);
            return await client.GetAsync(url);
        }
        
        protected async Task<HttpResponseMessage> Post<T>(string absoluteUrl, string relativeUrl, T body, int timeout = 60000)
        {
            var client = _httpClient.CreateClient();
            client.BaseAddress = new Uri(absoluteUrl);
            client.Timeout = TimeSpan.FromMilliseconds(timeout);
            var jsonEncoded = JsonConvert.SerializeObject(body, new JsonSerializerSettings
            {
                NullValueHandling = NullValueHandling.Ignore,
                ContractResolver = new DefaultContractResolver
                {
                    NamingStrategy = new CamelCaseNamingStrategy()
                }
            });
            
            var request = new HttpRequestMessage(HttpMethod.Post, relativeUrl);
            request.Content = new StringContent(jsonEncoded, Encoding.UTF8, "application/json");
            
            return await client.SendAsync(request);
        }
        
        public async Task<HttpStatusCode?> Status(int timeout = 2000)
        {
            try
            {
                var response = await Get(BaseUrl, "/status", timeout);
                return response.StatusCode;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Error while contacting {BaseUrl}");
                return null;
            }
        }
    }
}