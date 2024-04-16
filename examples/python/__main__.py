import pulumi
import pulumi_flyio as flyio

my_random_resource = flyio.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
