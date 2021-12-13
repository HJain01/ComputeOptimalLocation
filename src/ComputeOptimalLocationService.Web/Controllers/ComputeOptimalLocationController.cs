using Microsoft.AspNetCore.Mvc;

namespace ComputeOptimalLocationService.Web.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ComputeOptimalLocationController : Controller
    {
        private readonly Core.Services.ComputeOptimalLocationService _computeOptimalLocationService;

        public ComputeOptimalLocationController(Core.Services.ComputeOptimalLocationService computeOptimalLocationService)
        {
            _computeOptimalLocationService = computeOptimalLocationService;
        }
        
    }
}