# TODO

* [ ] Instance
    * [ ] DescribeInstanceAttribute
    * [ ] DescribeInstances
    * [ ] ModifyInstanceAttribute
    * [ ] RebootInstances
    * [ ] RunInstances
    * [ ] StartInstances
    * [ ] StopInstances
    * [ ] TerminateInstances
    * [ ] CopyInstances
    * [ ] CancelCopyInstances
    * [ ] ImportInstance
    * [ ] NiftyCreateInstanceSnapshot
    * [ ] NiftyModifyInstanceSnapshotAttribute
    * [ ] NiftyDescribeInstanceSnapshots
    * [ ] NiftyDeleteInstanceSnapshot
    * [ ] NiftyRestoreInstanceSnapshot
    * [ ] NiftyRetryImportInstance
* [ ] Volume
    * [ ] AttachVolume
    * [ ] CreateVolume
    * [ ] DeleteVolume
    * [ ] DescribeVolumes
    * [ ] DetachVolume
    * [ ] ModifyVolumeAttribute
* [ ] KeyPair
    * [ ] CreateKeyPair
    * [ ] DeleteKeyPair
    * [ ] DescribeKeyPairs
    * [ ] NiftyModifyKeyPairAttribute
    * [ ] ImportKeyPair
* [ ] Image
    * [ ] DescribeImages
    * [ ] CreateImage
    * [ ] DeleteImage
    * [ ] ModifyImageAttribute
    * [ ] NiftyAssociateImage
* [ ] LoadBalancer
    * [ ] ConfigureHealthCheck
    * [ ] CreateLoadBalancer
    * [ ] DeleteLoadBalancer
    * [ ] DeregisterInstancesFromLoadBalancer
    * [ ] DescribeInstanceHealth
    * [ ] DescribeLoadBalancers
    * [ ] RegisterInstancesWithLoadBalancer
    * [ ] RegisterPortWithLoadBalancer
    * [ ] SetFilterForLoadBalancer
    * [ ] UpdateLoadBalancer
    * [ ] UpdateLoadBalancerOption
    * [ ] SetLoadBalancerListenerSSLCertificate
    * [ ] UnsetLoadBalancerListenerSSLCertificate
    * [ ] ClearLoadBalancerSession
* [ ] SecurityGroup
    * [ ] AuthorizeSecurityGroupIngress
    * [ ] CreateSecurityGroup
    * [ ] DeleteSecurityGroup
    * [ ] DeregisterInstancesFromSecurityGroup
    * [ ] DescribeSecurityActivities
    * [ ] DescribeSecurityGroups
    * [ ] RegisterInstancesWithSecurityGroup
    * [ ] RevokeSecurityGroupIngress
    * [ ] UpdateSecurityGroup
    * [ ] UpdateSecurityGroupOption
    * [ ] DescribeSecurityGroupOption
    * [ ] NiftyRegisterRoutersWithSecurityGroup
    * [ ] NiftyDeregisterRoutersFromSecurityGroup
    * [ ] NiftyRegisterVpnGatewaysWithSecurityGroup
    * [ ] NiftyDeregisterVpnGatewaysFromSecurityGroup
* [ ] SslCertificate
    * [ ] CreateSslCertificate
    * [ ] DeleteSslCertificate
    * [ ] DescribeSslCertificates
    * [ ] DescribeSslCertificateAttribute
    * [ ] DownloadSslCertificate
    * [ ] ModifySslCertificateAttribute
    * [ ] RegisterCorporateInfoForCertificate
    * [ ] UploadSslCertificate
    * [ ] NiftyDescribeCorporateInfoForCertificate
* [ ] Usage
    * [ ] DescribeResources
    * [ ] DescribeServiceStatus
    * [ ] DescribeUsage
    * [ ] DescribeUserActivities
* [ ] Upload
    * [ ] DescribeUploads
    * [ ] CancelUpload
* [ ] IPAddress
    * [ ] AllocateAddress
    * [ ] AssociateAddress
    * [ ] NiftyModifyAddressAttribute
    * [ ] DescribeAddresses
    * [ ] DisassociateAddress
    * [ ] ReleaseAddress
* [ ] AutoScaling
    * [ ] NiftyCreateAutoScalingGroup
    * [ ] NiftyDescribeAutoScalingGroups
    * [ ] NiftyDescribeScalingActivities
    * [ ] NiftyUpdateAutoScalingGroup
    * [ ] NiftyDeleteAutoScalingGroup
* [ ] Alerm
    * [ ] NiftyCreateAlarm
    * [ ] NiftyDescribeAlarms
    * [ ] NiftyUpdateAlarm
    * [ ] NiftyDeleteAlarm
    * [ ] NiftyDescribeAlarmsPartitions
    * [ ] NiftyDescribePerformanceChart
    * [ ] NiftyDescribeAlarmHistory
    * [ ] NiftyDescribeAlarmRulesActivities
* [ ] PrivateLAN
    * [ ] NiftyCreatePrivateLan
    * [ ] NiftyDeletePrivateLan
    * [ ] NiftyDescribePrivateLans
    * [ ] NiftyModifyPrivateLanAttribute
* [ ] Router
    * [ ] NiftyCreateRouter
    * [ ] NiftyDeleteRouter
    * [ ] NiftyDescribeRouters
    * [ ] NiftyRebootRouters
    * [ ] NiftyModifyRouterAttribute
    * [ ] NiftyUpdateRouterNetworkInterfaces
    * [ ] NiftyReplaceRouterLatestVersion
    * [ ] NiftyRestoreRouterPreviousVersion
    * [ ] NiftyReleaseRouterBackupState
    * [ ] CreateRouteTable
    * [ ] DeleteRouteTable
    * [ ] DescribeRouteTables
    * [ ] AssociateRouteTable
    * [ ] DisassociateRouteTable
    * [ ] ReplaceRouteTableAssociation
    * [ ] CreateRoute
    * [ ] DeleteRoute
    * [ ] ReplaceRoute
    * [ ] NiftyCreateNatTable
    * [ ] NiftyDeleteNatTable
    * [ ] NiftyDescribeNatTables
    * [ ] NiftyAssociateNatTable
    * [ ] NiftyDisassociateNatTable
    * [ ] NiftyReplaceNatTableAssociation
    * [ ] NiftyCreateNatRule
    * [ ] NiftyDeleteNatRule
    * [ ] NiftyReplaceNatRule
    * [ ] NiftyEnableDhcp
    * [ ] NiftyDisableDhcp
    * [ ] CreateDhcpOptions
    * [ ] DeleteDhcpOptions
    * [ ] DescribeDhcpOptions
    * [ ] NiftyReplaceDhcpOption
    * [ ] NiftyReplaceDhcpConfig
    * [ ] NiftyDescribeDhcpStatus
    * [ ] NiftyCreateDhcpConfig
    * [ ] NiftyDeleteDhcpConfig
    * [ ] NiftyDescribeDhcpConfigs
    * [ ] NiftyCreateDhcpStaticMapping
    * [ ] NiftyDeleteDhcpStaticMapping
    * [ ] NiftyCreateDhcpIpAddressPool
    * [ ] NiftyDeleteDhcpIpAddressPool
    * [ ] NiftyCreateWebProxy
    * [ ] NiftyDeleteWebProxy
    * [ ] NiftyDescribeWebProxies
    * [ ] NiftyModifyWebProxyAttribute
* [ ] VpnGateway
    * [ ] CreateVpnGateway
    * [ ] DeleteVpnGateway
    * [ ] DescribeVpnGateways
    * [ ] NiftyDescribeVpnGatewayActivities
    * [ ] NiftyModifyVpnGatewayAttribute
    * [ ] NiftyRebootVpnGateways
    * [ ] NiftyAssociateRouteTableWithVpnGateway
    * [ ] NiftyDisassociateRouteTableFromVpnGateway
    * [ ] NiftyReplaceRouteTableAssociationWithVpnGateway
    * [ ] NiftyReplaceVpnGatewayLatestVersion
    * [ ] NiftyRestoreVpnGatewayPreviousVersion
    * [ ] NiftyReleaseVpnGatewayBackupState
    * [ ] NiftyUpdateVpnGatewayNetworkInterfaces
    * [ ] CreateCustomerGateway
    * [ ] DeleteCustomerGateway
    * [ ] DescribeCustomerGateways
    * [ ] NiftyModifyCustomerGatewayAttribute
    * [ ] CreateVpnConnection
    * [ ] DeleteVpnConnection
    * [ ] DescribeVpnConnections
* [ ] Others
    * [x] DescribeAvailabilityZones
    * [ ] DescribeRegions
    * [ ] AssociateUsers
    * [ ] DissociateUsers
    * [ ] DescribeAssociatedUsers
