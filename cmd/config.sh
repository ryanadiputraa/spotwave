#!/bin/bash
# Extract command line arguments
server_port="$1"
server_web_url="$2"
google_api_key="$3"
spotify_client_id="$4"
spotify_client_secret="$5"
spotify_state="$6"
spotify_redirect_uri="$7"
rapid_api_key="$8"

# Define the YAML content with placeholders replaced by command line arguments
YAML_CONTENT="server:
  port: ':$server_port'
  webUrl: '$server_web_url'

google:
  apiKey: '$google_api_key'

spotify:
  clientId: '$spotify_client_id'
  clientSecret: '$spotify_client_secret'
  state: '$spotify_state'
  redirectUri: '$spotify_redirect_uri'

rapidApi:
  apiKey: '$rapid_api_key'"

# Write the YAML content to the file
echo "$YAML_CONTENT" > config/config.yml

echo "YAML content with provided arguments has been written to config/config.yml"