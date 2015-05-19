
resource "google_compute_instance" "mesos-master" {
    count = "${var.nodes}"
    name = "${var.deployment_id}-node-${count.index}"
    machine_type = "n1-standard-1"
    zone = "europe-west1-d"
    tags = ["http","https","ssh","vpn"]

    disk {
        image = "ubuntu-os-cloud/ubuntu-1404-trusty-v20150128"
        type = "pd-ssd"
    }

    # declare metadata for configuration of the node
    metadata {
      nodecount = "${var.nodes}"
      myid = "${count.index}"
    }

    # network interface
    network_interface {
      network = "${google_compute_network.mesos-net.name}"
      access_config {
        // ephemeral address
      }
    }

    # define default connection for remote provisioners
    connection {
      user = "ubuntu"
      key_file = "~/.ssh/triforkse_rsa.pub"
    }
}
