# bootstrap_brokers_tls = "b-1.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:9094,b-2.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:9094,b-3.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:9094"
# zookeeper_connect_string = "z-1.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:2181,z-2.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:2181,z-3.example.uy30p7.c8.kafka.us-east-1.amazonaws.com:2181"

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region  = "us-east-1"
  profile = "tzookb"
}

resource "aws_secretsmanager_secret" "examplekeyyyy" {
  name = "AmazonMSK_tfkey"
}
# data "aws_secretsmanager_secret_version" "creds" {
#   # Fill in the name you gave to your secret
#   secret_id = aws_secretsmanager_secret.examplekeyyyy.id
# }

# resource "aws_vpc" "vpc" {
#   cidr_block = "192.168.0.0/22"
#   tags = {
#     Name = "zooooo"
#   }
# }

# data "aws_availability_zones" "azs" {
#   state = "available"
# }

# resource "aws_subnet" "subnet_az1" {
#   availability_zone = data.aws_availability_zones.azs.names[0]
#   cidr_block        = "192.168.0.0/24"
#   vpc_id            = aws_vpc.vpc.id
#   tags = {
#     Name = "z1-msk"
#   }
# }

# resource "aws_subnet" "subnet_az2" {
#   availability_zone = data.aws_availability_zones.azs.names[1]
#   cidr_block        = "192.168.1.0/24"
#   vpc_id            = aws_vpc.vpc.id
#   tags = {
#     Name = "z2-msk"
#   }
# }

# resource "aws_subnet" "subnet_az3" {
#   availability_zone = data.aws_availability_zones.azs.names[2]
#   cidr_block        = "192.168.2.0/24"
#   vpc_id            = aws_vpc.vpc.id
#   tags = {
#     Name = "z3-msk"
#   }
# }

# resource "aws_security_group" "sg" {
#   vpc_id = aws_vpc.vpc.id
# }

# resource "aws_kms_key" "kms" {
#   description = "example"
# }

# resource "aws_kms_alias" "kms" {
#   name          = "alias/msk-key"
#   target_key_id = aws_kms_key.kms.key_id
# }

# resource "aws_msk_cluster" "example" {
#   cluster_name           = "trf-msk-cluster"
#   kafka_version          = "3.2.0"
#   number_of_broker_nodes = 2

#   broker_node_group_info {
#     instance_type = "kafka.t3.small"
#     client_subnets = [
#       aws_subnet.subnet_az1.id,
#       aws_subnet.subnet_az2.id,
#       # aws_subnet.subnet_az3.id,
#     ]
#     storage_info {
#       ebs_storage_info {
#         volume_size = 100
#       }
#     }
#     security_groups = [aws_security_group.sg.id]
#   }

#   encryption_info {
#     encryption_at_rest_kms_key_arn = aws_kms_key.kms.arn
#   }

#   client_authentication {
#     sasl {
#       scram = 
#     }
#   }
# }

# output "zookeeper_connect_string" {
#   value = aws_msk_cluster.example.zookeeper_connect_string
# }

# output "bootstrap_brokers_tls" {
#   description = "TLS connection host:port pairs"
#   value       = aws_msk_cluster.example.bootstrap_brokers_tls
# }
