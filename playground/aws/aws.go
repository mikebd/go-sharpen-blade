package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"go-sharpen-blade/command"
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

	return nil
}
