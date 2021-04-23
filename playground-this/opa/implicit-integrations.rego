package implicit_integrations

same_volume_different_deployments[names] {
    deployment1 := input.items[_]
    pvc_name1 := deployment1.spec.template.spec.volumes[_].persistentVolumeClaim.claimName
    deployment2 := input.items[_]
    deployment2.metadata.uid != deployment1.metadata.uid
    pvc_name2 := deployment2.spec.template.spec.volumes[_].persistentVolumeClaim.claimName
    names := { "deployment1": deployment1.metadata.name, "deployment2": deployment2.metadata.name, "pvc_name": pvc_name2 }
    pvc_name1 == pvc_name2
}
