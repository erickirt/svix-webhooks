// this file is @generated
using System.Text;
using Newtonsoft.Json;

namespace Svix.Models
{
    public class ListResponseIngestEndpointOut
    {
        [JsonProperty("data", Required = Required.Always)]
        public required List<IngestEndpointOut> Data { get; set; }

        [JsonProperty("done", Required = Required.Always)]
        public required bool Done { get; set; }

        [JsonProperty("iterator")]
        public string? Iterator { get; set; } = null;

        [JsonProperty("prevIterator")]
        public string? PrevIterator { get; set; } = null;

        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();

            sb.Append("class ListResponseIngestEndpointOut {\n");
            sb.Append("  Data: ").Append(Data).Append('\n');
            sb.Append("  Done: ").Append(Done).Append('\n');
            sb.Append("  Iterator: ").Append(Iterator).Append('\n');
            sb.Append("  PrevIterator: ").Append(PrevIterator).Append('\n');
            sb.Append("}\n");
            return sb.ToString();
        }
    }
}
