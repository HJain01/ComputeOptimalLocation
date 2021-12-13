using System.Threading.Tasks;
using ComputeOptimalLocationService.Core.Services;
using Microsoft.AspNetCore.Mvc;

namespace ComputeOptimalLocationService.Web.Controllers
{
    [Route("[controller]")]
    public class UnsubscribeController : Controller
    {
        private readonly EmailNotificationService _emailNotificationService;

        public UnsubscribeController(EmailNotificationService emailNotificationService)
        {
            _emailNotificationService = emailNotificationService;
        }

        [HttpGet("{arn}")]
        public async Task<IActionResult> Unsubscribe(string arn)
        {
            await _emailNotificationService.Unsubscribe(arn);

            return Ok("You have been unsubscribed!");
        }
    }
}