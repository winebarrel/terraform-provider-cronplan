terraform {
  required_providers {
    oneshot = {
      source  = "winebarrel/cronplan"
      version = ">= 0.3.0"
    }
  }
}

provider "cronplan" {
}
