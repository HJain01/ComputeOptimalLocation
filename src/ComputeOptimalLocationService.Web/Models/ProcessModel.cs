namespace ComputeOptimalLocationService.Web.Models
{
    public class ProcessModel
    {
        public bool? SkipValidate { get; set; }
        public bool? SaveToDb  { get; set; }
        public bool? DryRun  { get; set; }
        public bool? Submit { get; set; }
        public bool? AddOnly { get; set; }
        public bool? AutoComplete { get; set; }
    }
}