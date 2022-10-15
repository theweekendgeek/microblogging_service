provider "google" {
  project = "twitter-service-365615"
  region = "europe-west1"
  zone = "europe-west1-b"
}

resource "google_storage_bucket" "source" {
  name = "twitter-service-source"
  location = "EU"
  force_destroy = true

  uniform_bucket_level_access = true
}