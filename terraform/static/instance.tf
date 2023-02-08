resource "aws_instance" "app_server" {
  ami                    = var.ami
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.ec2_trafic.id]
  key_name               = var.key_name

  tags = {
    Name = var.instance_name
  }
  user_data = data.template_file.init.rendered
}

data "template_file" "init" {
    template = file("${path.module}/init.tpl")
}