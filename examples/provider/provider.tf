terraform {
  required_providers {
    cronplan = {
      source  = "winebarrel/cronplan"
      version = ">= 3"
    }
  }
}

provider "cronplan" {
}
