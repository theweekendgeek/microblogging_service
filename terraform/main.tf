provider "google" {
  project     = "twitter-service-365615"
  region      = "europe-west1"
  zone        = "europe-west1-b"
  credentials = "./twitter-service-365615-c5315f98529c.json"
}

resource "google_cloud_scheduler_job" "job" {
	name             = "test-job"
	description      = "test http job"
	schedule         = "*/30 */4 * * *"
	time_zone        = "America/New_York"
	attempt_deadline = "320s"

	retry_config {
		retry_count = 1
	}

	http_target {
		http_method = "GET"
		uri         = "https://us-central1-twitter-service-365615.cloudfunctions.net/twitter-service"
	}
}
