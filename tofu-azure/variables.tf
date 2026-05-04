variable "resource_group_name" {
  description = "Name der Resource Group"
  type        = string
  default     = "student-aks-rg"
}

variable "location" {
  description = "Azure Region"
  type        = string
  default     = "germanywestcentral"
}

variable "cluster_name" {
  description = "Name des AKS Clusters"
  type        = string
  default     = "student-aks-cluster"
}

variable "dns_prefix" {
  description = "DNS Prefix für den AKS Cluster"
  type        = string
  default     = "student-aks"
}

variable "node_count" {
  description = "Anzahl der AKS Nodes"
  type        = number
  default     = 1
}

variable "vm_size" {
  description = "VM-Größe für den AKS Node Pool"
  type        = string
  default     = "Standard_DC2s_v3"
}