package disposability

no_liveness_probe[names] {
    deployment := input.items[_]
    container := deployment.spec.template.spec.containers[_]
    names := { "deployment": deployment.metadata.name, "container": container.name }
    not container.livenessProbe
}

no_readiness_probe[names] {
    deployment := input.items[_]
    container := deployment.spec.template.spec.containers[_]
    names := { "deployment": deployment.metadata.name, "container": container.name }
    not container.readinessProbe
}

high_resource_request_cpu[names] {
    deployment := input.items[_]
    container := deployment.spec.template.spec.containers[_]
    names := { "deployment": deployment.metadata.name, "container": container.name }
    # hmm, not so easy. How do we compare the size values which are given as strings with units

}
