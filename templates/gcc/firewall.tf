resource "google_compute_firewall" "mesos-firwall" {
    name = "test"
    network = "${var.deployment_id}-net"

    allow {
        protocol = "icmp"
    }

    allow {
        protocol = "tcp"
        ports = ["80", "8080", "1000-2000"]
    }

    source_ranges = ["0.0.0.0/0"]

    source_tags = ["web"]
}