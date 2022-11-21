
resource "aws_vpc" "IaCPresentationVpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  instance_tenancy     = "default"
}

resource "aws_subnet" "IacPresentationPublicSubnetOne" {
  vpc_id               = aws_vpc.IaCPresentationVpc.id
  cidr_block           = "10.0.0.0/24"
  availability_zone_id = "use1-az1"
}

resource "aws_subnet" "IacPresentationPrivateSubnetOne" {
  vpc_id               = aws_vpc.IaCPresentationVpc.id
  cidr_block           = "10.0.1.0/24"
  availability_zone_id = "use1-az1"
}

resource "aws_subnet" "IacPresentationPublicSubnetTwo" {
  vpc_id               = aws_vpc.IaCPresentationVpc.id
  cidr_block           = "10.0.2.0/24"
  availability_zone_id = "use1-az1"
}

resource "aws_subnet" "IacPresentationPrivateSubnetTwo" {
  vpc_id               = aws_vpc.IaCPresentationVpc.id
  cidr_block           = "10.0.3.0/24"
  availability_zone_id = "use1-az1"
}

resource "aws_internet_gateway" "IaCPresentationInternetGateway" {
}

resource "aws_internet_gateway_attachment" "InternetGatewayAttachment" {
  internet_gateway_id = aws_internet_gateway.IaCPresentationInternetGateway.id
  vpc_id              = aws_vpc.IaCPresentationVpc.id
}

resource "aws_eip" "IaCPresentationElasticIp" {
  vpc = true
}

resource "aws_nat_gateway" "IaCPresentationNatGateway" {
  subnet_id = aws_subnet.IacPresentationPublicSubnetOne.id
  allocation_id = aws_eip.IaCPresentationElasticIp.id
}

resource "aws_route_table" "IaCPresentationRouteTablePublic" {
  vpc_id = aws_vpc.IaCPresentationVpc.id
}

resource "aws_route_table" "IaCPresentationRouteTablePrivate" {
  vpc_id = aws_vpc.IaCPresentationVpc.id
}

resource "aws_route_table_association" "PublicSubnetOneRouteTableAssociation" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePublic.id
  subnet_id      = aws_subnet.IacPresentationPublicSubnetOne.id
}

resource "aws_route_table_association" "PublicSubnetTwoRouteTableAssociation" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePublic.id
  subnet_id      = aws_subnet.IacPresentationPublicSubnetTwo.id
}

resource "aws_route_table_association" "PrivateSubnetOneRouteTableAssociation" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePublic.id
  subnet_id      = aws_subnet.IacPresentationPrivateSubnetOne.id
}

resource "aws_route_table_association" "PrivateSubnetTwoRouteTableAssociation" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePublic.id
  subnet_id      = aws_subnet.IacPresentationPrivateSubnetTwo.id
}

resource "aws_route" "IaCPresentationRoutePublic" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePublic.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.IaCPresentationInternetGateway.id
}

resource "aws_route" "IaCPresentationRoutePrivate" {
  route_table_id = aws_route_table.IaCPresentationRouteTablePrivate.id
  nat_gateway_id = aws_nat_gateway.IaCPresentationNatGateway.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_security_group" "IaCPresentationFunctionSecurityGroup" {
  vpc_id = aws_vpc.IaCPresentationVpc.id
  name = "IaC Presentation Function Security Group"
  description = "IaC Presentation Function Security Group"
  ingress {
    from_port = 0
    protocol  = "tcp"
    to_port   = 65535
    cidr_blocks = [aws_vpc.IaCPresentationVpc.cidr_block]
  }
}