variable "region" { }
variable "project" { }

provider "google" {
    account_file = "account.json"
    project = "${var.project}"
    region = "${var.region}"
}

resource "google_compute_instance" "default" {
    name = "micro-master"
    machine_type = "n1-standard-1"
    zone = "europe-west1-d"

    disk {
        image = "debian-7-wheezy-v20140814"
    }
    
    network_interface {
        network = "default"
        access_config {
            // Ephemeral IP
        }
    }

    service_account {
        scopes = ["userinfo-email", "compute-ro", "storage-ro"]
    }
}
