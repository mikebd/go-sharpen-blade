package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"go-sharpen-blade/internal/command"
	"log"
)

func Register() {
	command.Register("aws", aws)
}

func aws() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	log.Println("AWS Region:", cfg.Region)

	// Display charged resources in the region

	autoScalingClient := autoscaling.NewFromConfig(cfg)
	autoScalingGroups, err := autoScalingClient.DescribeAutoScalingGroups(context.TODO(), &autoscaling.DescribeAutoScalingGroupsInput{})
	if err != nil {
		return err
	}
	log.Println("Auto Scaling Groups:", autoScalingGroups.AutoScalingGroups)

	elbClient := elasticloadbalancingv2.NewFromConfig(cfg)
	loadBalancers, err := elbClient.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{})
	if err != nil {
		return err
	}
	log.Println("Load Balancers:", loadBalancers.LoadBalancers)

	// Delete charged resources in the region

	forceDelete := true

	for _, autoScalingGroup := range autoScalingGroups.AutoScalingGroups {
		result, err := autoScalingClient.DeleteAutoScalingGroup(context.TODO(), &autoscaling.DeleteAutoScalingGroupInput{
			AutoScalingGroupName: autoScalingGroup.AutoScalingGroupName,
			ForceDelete:          &forceDelete,
		})
		if err != nil {
			return err
		}
		log.Println("Delete Auto Scaling Group:", result)
	}

	for _, loadBalancer := range loadBalancers.LoadBalancers {
		result, err := elbClient.DeleteLoadBalancer(context.TODO(), &elasticloadbalancingv2.DeleteLoadBalancerInput{
			LoadBalancerArn: loadBalancer.LoadBalancerArn,
		})
		if err != nil {
			return err
		}
		log.Println("Delete Load Balancer:", result)
	}

	return nil
}
