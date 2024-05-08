// Code generated by smithy-go-codegen DO NOT EDIT.

// Package ebs provides the API client, operations, and parameter types for Amazon
// Elastic Block Store.
//
// You can use the Amazon Elastic Block Store (Amazon EBS) direct APIs to create
// Amazon EBS snapshots, write data directly to your snapshots, read data on your
// snapshots, and identify the differences or changes between two snapshots. If
// you’re an independent software vendor (ISV) who offers backup services for
// Amazon EBS, the EBS direct APIs make it more efficient and cost-effective to
// track incremental changes on your Amazon EBS volumes through snapshots. This can
// be done without having to create new volumes from snapshots, and then use Amazon
// Elastic Compute Cloud (Amazon EC2) instances to compare the differences.
//
// You can create incremental snapshots directly from data on-premises into
// volumes and the cloud to use for quick disaster recovery. With the ability to
// write and read snapshots, you can write your on-premises data to an snapshot
// during a disaster. Then after recovery, you can restore it back to Amazon Web
// Services or on-premises from the snapshot. You no longer need to build and
// maintain complex mechanisms to copy data to and from Amazon EBS.
//
// This API reference provides detailed information about the actions, data types,
// parameters, and errors of the EBS direct APIs. For more information about the
// elements that make up the EBS direct APIs, and examples of how to use them
// effectively, see [Accessing the Contents of an Amazon EBS Snapshot]in the Amazon Elastic Compute Cloud User Guide. For more
// information about the supported Amazon Web Services Regions, endpoints, and
// service quotas for the EBS direct APIs, see [Amazon Elastic Block Store Endpoints and Quotas]in the Amazon Web Services General
// Reference.
//
// [Accessing the Contents of an Amazon EBS Snapshot]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html
// [Amazon Elastic Block Store Endpoints and Quotas]: https://docs.aws.amazon.com/general/latest/gr/ebs-service.html
package ebs
