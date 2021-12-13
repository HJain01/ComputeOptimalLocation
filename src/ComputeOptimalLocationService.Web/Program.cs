using System;
using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Hosting;
using CustomEnvironmentConfig;
using ComputeOptimalLocationService.Core.Configuration;
using Serilog;
using Serilog.Debugging;
using Serilog.Events;

namespace ComputeOptimalLocationService.Web
{
    public class Program
    {
        public static void Main(string[] args)
        {
            SelfLog.Enable(Console.WriteLine);

            var configuration = ConfigurationParser.Parse<BaseConfiguration>();

            Log.Logger = new LoggerConfiguration()
                        .MinimumLevel.Debug()
                        .MinimumLevel.Override("Microsoft", LogEventLevel.Warning)
                        .MinimumLevel.Override("System", LogEventLevel.Warning)
                         // The Authentication handler automatically outputs debug information. These messages are
                         // irrelevant for how we authenticate and introduce noise in the logs. Below we raise the
                         // minimum logging level to Information so these messages do not end up in the final output.
                         // https://github.com/aspnet/Security/blob/dev/src/Microsoft.AspNetCore.Authentication/LoggingExtensions.cs#L44
                         // https://github.com/aspnet/Security/blob/dev/src/Microsoft.AspNetCore.Authentication/AuthenticationHandler.cs#L143
                        .Enrich.FromLogContext()
                        .Enrich.WithProperty("source", "PhTS")
                        .Enrich.WithProperty("env", System.Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT"))
                        .WriteTo.Console()
                        .CreateLogger();
            try
            {
                Log.Information("Starting PhTS Service");
                BuildWebHost(args, configuration).Run();
            }
            catch (Exception ex)
            {
                Log.Fatal(ex, "Host terminated unexpectedly");
            }
            finally
            {
                Log.CloseAndFlush();
            }
        }

        private static IWebHost BuildWebHost(string[] args, BaseConfiguration config) =>
            WebHost.CreateDefaultBuilder(args)
                   .UseEnvironmentConfiguration(config)
                   .UseStartup<Startup>()
                   .UseSerilog()
                   .UseKestrel(options =>
                    {
                        options.ListenAnyIP(config.Port);
                    })
                   .Build();
    }
}