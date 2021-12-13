using System;
using ComputeOptimalLocationService.Core.Configuration;
using ComputeOptimalLocationService.Core.Interfaces.Repositories;
using ComputeOptimalLocationService.Core.Services;
using ComputeOptimalLocationService.Repository.Repositories;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Newtonsoft.Json.Serialization;

namespace ComputeOptimalLocationService.Web
{
    public class Startup
    {
        private readonly BaseConfiguration _configuration;
        
        public Startup(BaseConfiguration configuration)
        {
            _configuration = configuration;
            Started = DateTime.UtcNow;
        }

        public static DateTime Started { get; private set; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers()
                .AddNewtonsoftJson(options =>
            {
                // Camel case property names: https://google.github.io/styleguide/jsoncstyleguide.xml#Property_Name_Guidelines
                options.SerializerSettings.ContractResolver = new CamelCasePropertyNamesContractResolver();
                options.SerializerSettings.ReferenceLoopHandling = Newtonsoft.Json.ReferenceLoopHandling.Ignore;
            });
            // swagger (not in prod)
            if (Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT") != "prod")
            {
                services.AddSwaggerGen();
            }
            
            services.AddRouting(options =>
            {
                options.LowercaseUrls = true;
            });
            
            services.AddMvc(options =>
            {
                options.EnableEndpointRouting = false;
            });

            services.AddHttpClient();
            
            // Services
            services.AddTransient<Core.Services.ComputeOptimalLocationService>();
            services.AddTransient<EmailNotificationService>();
            services.AddTransient<IStatusService, StatusService>();

            //Repositories
            services.AddTransient<INotificationsRepository, SnsRepository>();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            
            if (System.Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT") != "prod")
            {
                app.UseSwagger();
                app.UseSwaggerUI(c => { c.SwaggerEndpoint("/swagger/v1/swagger.json", "Compute Optimal Location Service"); });
            }
            
            
            app.UseExceptionHandler("/error");

            app.UseRouting();
            app.UseMvc();
        }
    }
}