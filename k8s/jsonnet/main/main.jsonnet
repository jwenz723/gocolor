local k = (import "1.18/main.libsonnet");
function (
  port=8080,
  portName="http-service",
  image="jwenz723/gocolor:sha-cac4dd6",
  name="gocolor",
  account="unknown",
  color="yello",
  region="unknown",
)
[
  k.apps.v1.deployment.new(
    name=name,
    containers=[
      k.core.v1.container.new(name=name, image=image) +
      k.core.v1.container.withArgs(["-addr=:" + port]) +
      k.core.v1.container.withEnvMap({
        "GOCOLOR_ACCOUNT": account,
        "GOCOLOR_COLOR": color,
        "GOCOLOR_REGION": region,
      }) +
      k.core.v1.container.withPorts([k.core.v1.containerPort.newNamed(port, portName)])
    ],
    replicas=1,
  ) + k.apps.v1.deployment.metadata.withLabels({
    "app.kubernetes.io/name": name,
  }) + k.apps.v1.deployment.spec.selector.withMatchLabels({
    "app.kubernetes.io/name": name,
  }) + k.apps.v1.deployment.spec.template.metadata.withLabels({
    "app.kubernetes.io/name": name,
  }),
  k.core.v1.service.new(
    name=name,
    selector={
      "app.kubernetes.io/name": name,
    },
    ports=[k.core.v1.servicePort.newNamed(
      name=portName,
      port=port,
      targetPort=portName,
    )]
  ) + k.core.v1.service.metadata.withLabels({
    "app.kubernetes.io/name": name,
  })
]
