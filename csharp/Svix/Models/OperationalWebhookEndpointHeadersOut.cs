// this file is @generated
using System.Text;
using Newtonsoft.Json;

namespace Svix.Models
{
    public class OperationalWebhookEndpointHeadersOut
    {
        [JsonProperty("headers", Required = Required.Always)]
        public required Dictionary<string, string> Headers { get; set; }

        [JsonProperty("sensitive", Required = Required.Always)]
        public required List<string> Sensitive { get; set; }

        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();

            sb.Append("class OperationalWebhookEndpointHeadersOut {\n");
            sb.Append("  Headers: ").Append(Headers).Append('\n');
            sb.Append("  Sensitive: ").Append(Sensitive).Append('\n');
            sb.Append("}\n");
            return sb.ToString();
        }
    }
}
