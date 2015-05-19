 account_file="cs-cisco-ca214d29dc66.json"
 gce_ssh_user="554985525398-p9se88l5e3fupvj1v8t6tujq5qsumh1q.apps.googleusercontent.com"
 gce_ssh_private_key_file="pkey"
 region ="europe-west1"
 zone="europe-west1-d"
 project="cs-cisco"
 image="ubuntu-os-cloud/ubuntu-1404-trusty-v20150128"
 master_machine_type="n1-standard-2"
 slave_machine_type= "n1-standard-4"
 network= "10.20.30.0/24"
 localaddress="92.111.228.8/32"
 name="mymesoscluster"
 masters= "1"



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
      user = "${var.gce_ssh_user}"
      key_file = "${var.gce_ssh_private_key_file}"
    }
}
