function (
  port=8080,
  portName="http-service",
  image="jwenz723/gocolor:sha-cac4dd6",
  name="gocolor"
)
{
  local k = (import "vendor/1.18/main.libsonnet"),
  gocolor: {
      deployment: k.apps.v1.deployment.new(
        name=name,
        containers=[
          k.core.v1.container.new(name=name, image=image) +
          k.core.v1.container.withArgs(["-addr=:" + port]) +
          k.core.v1.container.withPorts([k.core.v1.containerPort.newNamed(port, portName)])
        ],
        replicas=1,
      ) + k.apps.v1.deployment.metadata.withLabels({
        name: name,
      }),
      service: k.core.v1.service.new(
        name=name,
        selector={
          name: name,
        },
        ports=[k.core.v1.servicePort.newNamed(
          name=portName,
          port=port,
          targetPort=portName,
        )]
      )
  }
}
