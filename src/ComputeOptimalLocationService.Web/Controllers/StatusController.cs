using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using ComputeOptimalLocationService.Core.Services;

namespace ComputeOptimalLocationService.Web.Controllers
{
    [ApiController]
    [Route("[controller]")]
    [AllowAnonymous]
    public class StatusController : ControllerBase
    {
        private readonly IStatusService _statusService;

        public StatusController(IStatusService statusService)
        {
            _statusService = statusService;
        }

        [HttpGet]
        [ProducesResponseType(StatusCodes.Status200OK)]
        [ProducesResponseType(StatusCodes.Status503ServiceUnavailable)]
        public async Task<IActionResult> GetStatus()
        {
            return await _statusService.IsHealthy() 
                ? Ok()
                : StatusCode(StatusCodes.Status503ServiceUnavailable);
        }
    }
}