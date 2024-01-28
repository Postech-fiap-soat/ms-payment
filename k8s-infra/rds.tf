resource "aws_security_group" "sg-db" {
    vpc_id = aws_vpc.new-vpc.id
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        prefix_list_ids = []
    }
    ingress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        prefix_list_ids = []
    }
    tags = {
        Name = "${var.prefix}-sg-db"
    }
}

resource "aws_db_instance" "soat-db-instance-2" {
    allocated_storage = 10
    engine = "mysql"
    engine_version = "5.7"
    instance_class = "db.t2.micro"
    username = "soatuser"
    password = "soatpassword"
    skip_final_snapshot = true
    db_name = "soatdb"
    db_subnet_group_name = aws_db_subnet_group.db-subnet.id
    vpc_security_group_ids = [aws_security_group.sg-db.id]
}

resource "aws_db_subnet_group" "db-subnet" {
    name = "dbsubnet"
    subnet_ids = aws_subnet.new-subnets[*].id
}