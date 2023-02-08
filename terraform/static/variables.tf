variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
  default     = "Pinger Static Bot"
}
variable "ami" {
  description = "The EC2 Instance AMI ID"
  type        = string
  default     = "ami-08e2d37b6a0129927"
}
variable "instance_type" {
  description = "The EC2 Instance Type"
  type        = string
  default     = "t2.micro"
}
variable "region" {
  description = "The EC2 Instance Region"
  type        = string
  default     = "us-west-2"
}
variable "key_name" {
  description = "The EC2 Key Pair to allow SSH access to the instance"
  type        = string
  default     = "AWS"
}
variable "security_group" {
  description = "The EC2 Security Group to assign to the instance"
  type        = string
  default     = "AllowPingerBotTraffic"
}
variable "config_files" {
  description = "Config Files path"
  type        = list(string)
}
variable "credentials_files" {
  description = "Credentials Files path"
  type        = list(string)
}
variable "profile" {
  description = "AWS Profile"
  type        = string
}