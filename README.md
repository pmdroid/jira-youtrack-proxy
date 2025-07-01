# Jira-YouTrack Proxy

A lightweight HTTP proxy server that translates Jira API requests to YouTrack API calls, enabling seamless integration between systems that expect Jira's REST API format.

**Note**: This proxy implements only a subset of the Jira REST API, specifically designed to enable Fleet MDM to create vulnerability tickets in YouTrack.

## Features

- **Issue Creation**: Proxies Jira issue creation requests to YouTrack
- **Project Details**: Fetches project information from YouTrack
- **Field Mapping**: Configurable mapping between Jira and YouTrack field IDs
- **Health Check**: Built-in health endpoint for monitoring

## Configuration

Configure the proxy using environment variables:

- `YOUTRACK_URL`: YouTrack instance URL (default: `https://example.youtrack.cloud`)
- `PORT`: Server port (default: `8080`)
- `FIELD_MAPPING_FILE_PATH`: Path to field mappings JSON file (default: `field_mappings.json`)

## Field Mappings

Edit `field_mappings.json` to map YouTrack project IDs to YouTrack field IDs:

```json
{
  "type_field_map": {
    "youtrack-project-id": "youtrack-field-id"
  },
  "priority_field_map": {
    "youtrack-project-id": "youtrack-field-id"
  },
  "assignee_field_map": {
    "youtrack-project-id": "youtrack-field-id"
  }
}
```

## Usage

```bash
go run main.go
```

## API Endpoints

- `POST /rest/api/2/issue` - Create issue
- `GET /rest/api/2/project/:id` - Get project details
- `GET /health` - Health check