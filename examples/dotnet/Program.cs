using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Flyio = Pulumi.Flyio;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Flyio.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

