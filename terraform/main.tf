variable "uri" {
	type = string
}

variable "credentials" {
	type = string
}

provider "google" {
  project     = "twitter-service-365615"
  region      = "europe-west1"
  zone        = "europe-west1-b"
  credentials = var.credentials
}

resource "google_cloud_scheduler_job" "job" {
  name             = "test-job"
  description      = "test http job"
  schedule         = "0 */1 * * *"
  time_zone        = "Europe/Berlin"
  attempt_deadline = "320s"

  retry_config {
    retry_count = 1
  }



  http_target {
    http_method = "GET"
    uri         = var.uri
  }
}
