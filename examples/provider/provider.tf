terraform {
  required_providers {
    oneshot = {
      source  = "winebarrel/cronplan"
      version = ">= 3"
    }
  }
}

provider "cronplan" {
}
