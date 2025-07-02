# DevRev Webhooks Test Server

This is a simple server to test DevRev webhooks. It handles webhook verification
and processes webhook events for work created, updated, and deleted.
Detailed documentation on webhooks can be found [here](https://developer.devrev.ai/public/guides/webhooks).

## Setup

1. Start ngrok:
```bash
ngrok http 3000
```

2. Register your webhook with DevRev using the ngrok URL:
```bash
curl --request POST 'https://api.devrev.ai/webhooks.create' \
  --header "Authorization: Bearer $DEVREV_TOKEN" \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "event_types": [
      "work_created",
      "work_updated", 
      "work_deleted"
     ],
    "url": "your-ngrok-url"
  }'
```

3. Start the webhook server:

The `WEBHOOK_SECRET` environment variable is required. It will be available as
`secret` in the response of the above `webhooks.create` API call.

```bash
WEBHOOK_SECRET=your_webhook_secret go run main.go
```

## Example webhook payload

The following is an example of a `work_updated` webhook payload. Note that
`work_updated` contains `old_work` and `work` fields.

```
{
  "id": "don:integration:dvrv-us-1:devo/11FVC3ScK:webhook/1jGzMh1N:webhook_event/U7gbBDxjiWM",
  "timestamp": "2025-07-02T13:07:40.937074Z",
  "webhook_id": "don:integration:dvrv-us-1:devo/11FVC3ScK:webhook/1jGzMh1N",
  "work_updated": {
    "old_work": {
      "type": "issue",
      "applies_to_part": {
        "type": "product",
        "display_id": "PROD-2",
        "id": "don:core:dvrv-us-1:devo/11FVC3ScK:product/2",
        "id_v1": "don:DEV-11FVC3ScK:product:2",
        "name": "prod-foo",
        "owned_by": [
          {
            "type": "dev_user",
            "display_handle": "shivansh-rai",
            "display_id": "DEVU-1",
            "display_name": "shivansh-rai",
            "email": "shivansh.rai@devrev.ai",
            "full_name": "Shivansh Rai",
            "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
            "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
            "state": "active",
            "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
          }
        ],
        "stage": {
          "name": ""
        }
      },
      "body": "\u003cdon:core:dvrv-us-1:devo/11FVC3ScK:issue/98\u003e",
      "created_by": {
        "type": "dev_user",
        "display_handle": "shivansh-rai",
        "display_id": "DEVU-1",
        "display_name": "shivansh-rai",
        "email": "shivansh.rai@devrev.ai",
        "full_name": "Shivansh Rai",
        "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
        "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
        "state": "active",
        "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
      },
      "created_date": "2025-06-14T13:20:46.963Z",
      "custom_fields": {
        "ctype__color": "Blue",
        "tnt__a_rich_text_field": "A new issue \u003cdon:core:dvrv-us-1:devo/11FVC3ScK:issue/91\u003e",
        "tnt__a_text_field": ""
      },
      "custom_schema_fragments": [
        "don:core:dvrv-us-1:devo/11FVC3ScK:tenant_fragment/503",
        "don:core:dvrv-us-1:devo/11FVC3ScK:custom_type_fragment/480"
      ],
      "display_id": "ISS-98",
      "id": "don:core:dvrv-us-1:devo/11FVC3ScK:issue/98",
      "id_v1": "don:DEV-11FVC3ScK:issue:98",
      "modified_by": {
        "type": "dev_user",
        "display_handle": "shivansh-rai",
        "display_id": "DEVU-1",
        "display_name": "shivansh-rai",
        "email": "shivansh.rai@devrev.ai",
        "full_name": "Shivansh Rai",
        "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
        "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
        "state": "active",
        "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
      },
      "modified_date": "2025-07-02T13:04:12.1Z",
      "owned_by": [
        {
          "type": "dev_user",
          "display_handle": "gowtham-tg",
          "display_id": "DEVU-10",
          "display_name": "gowtham-tg",
          "email": "gowtham.tg@devrev.ai",
          "full_name": "Gowtham Gopinath",
          "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/10",
          "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-10",
          "state": "active",
          "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Gowtham%20Gopinath.png"
        }
      ],
      "priority": "p1",
      "priority_v2": {
        "id": 2,
        "label": "P1",
        "ordinal": 2
      },
      "references": [
        {
          "type": "issue",
          "display_id": "ISS-98",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:issue/98",
          "id_v1": "don:DEV-11FVC3ScK:issue:98",
          "owned_by": [
            {
              "type": "dev_user",
              "display_handle": "gowtham-tg",
              "display_id": "DEVU-10",
              "display_name": "gowtham-tg",
              "email": "gowtham.tg@devrev.ai",
              "full_name": "Gowtham Gopinath",
              "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/10",
              "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-10",
              "state": "active",
              "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Gowtham%20Gopinath.png"
            }
          ],
          "priority": "p2",
          "priority_v2": {
            "id": 3,
            "label": "P2",
            "ordinal": 3
          },
          "stage": {
            "name": "prioritized",
            "stage": {
              "display_name": "Prioritized",
              "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_stage/27",
              "name": "prioritized"
            }
          },
          "title": "issue with a subtype"
        }
      ],
      "stage": {
        "display_name": "Prioritized",
        "name": "prioritized",
        "notes": "",
        "ordinal": 3600,
        "stage": {
          "display_name": "Prioritized",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_stage/27",
          "name": "prioritized"
        },
        "state": {
          "display_name": "Open",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_state/1",
          "is_final": false,
          "name": "open"
        }
      },
      "state": "open",
      "stock_schema_fragment": "don:core:dvrv-us-1:stock_sf/5774789",
      "subtype": "test",
      "title": "issue with a subtype"
    },
    "work": {
      "type": "issue",
      "applies_to_part": {
        "type": "product",
        "display_id": "PROD-2",
        "id": "don:core:dvrv-us-1:devo/11FVC3ScK:product/2",
        "id_v1": "don:DEV-11FVC3ScK:product:2",
        "name": "prod-foo",
        "owned_by": [
          {
            "type": "dev_user",
            "display_handle": "shivansh-rai",
            "display_id": "DEVU-1",
            "display_name": "shivansh-rai",
            "email": "shivansh.rai@devrev.ai",
            "full_name": "Shivansh Rai",
            "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
            "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
            "state": "active",
            "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
          }
        ],
        "stage": {
          "name": ""
        }
      },
      "body": "\u003cdon:core:dvrv-us-1:devo/11FVC3ScK:issue/98\u003e",
      "created_by": {
        "type": "dev_user",
        "display_handle": "shivansh-rai",
        "display_id": "DEVU-1",
        "display_name": "shivansh-rai",
        "email": "shivansh.rai@devrev.ai",
        "full_name": "Shivansh Rai",
        "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
        "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
        "state": "active",
        "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
      },
      "created_date": "2025-06-14T13:20:46.963Z",
      "custom_fields": {
        "ctype__color": "Blue",
        "tnt__a_rich_text_field": "A new issue \u003cdon:core:dvrv-us-1:devo/11FVC3ScK:issue/91\u003e",
        "tnt__a_text_field": ""
      },
      "custom_schema_fragments": [
        "don:core:dvrv-us-1:devo/11FVC3ScK:tenant_fragment/503",
        "don:core:dvrv-us-1:devo/11FVC3ScK:custom_type_fragment/480"
      ],
      "display_id": "ISS-98",
      "id": "don:core:dvrv-us-1:devo/11FVC3ScK:issue/98",
      "id_v1": "don:DEV-11FVC3ScK:issue:98",
      "modified_by": {
        "type": "dev_user",
        "display_handle": "shivansh-rai",
        "display_id": "DEVU-1",
        "display_name": "shivansh-rai",
        "email": "shivansh.rai@devrev.ai",
        "full_name": "Shivansh Rai",
        "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/1",
        "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-1",
        "state": "active",
        "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Shivansh%20Rai.png"
      },
      "modified_date": "2025-07-02T13:07:40.704Z",
      "owned_by": [
        {
          "type": "dev_user",
          "display_handle": "gowtham-tg",
          "display_id": "DEVU-10",
          "display_name": "gowtham-tg",
          "email": "gowtham.tg@devrev.ai",
          "full_name": "Gowtham Gopinath",
          "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/10",
          "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-10",
          "state": "active",
          "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Gowtham%20Gopinath.png"
        }
      ],
      "priority": "p2",
      "priority_v2": {
        "id": 3,
        "label": "P2",
        "ordinal": 3
      },
      "references": [
        {
          "type": "issue",
          "display_id": "ISS-98",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:issue/98",
          "id_v1": "don:DEV-11FVC3ScK:issue:98",
          "owned_by": [
            {
              "type": "dev_user",
              "display_handle": "gowtham-tg",
              "display_id": "DEVU-10",
              "display_name": "gowtham-tg",
              "email": "gowtham.tg@devrev.ai",
              "full_name": "Gowtham Gopinath",
              "id": "don:identity:dvrv-us-1:devo/11FVC3ScK:devu/10",
              "id_v1": "don:DEV-11FVC3ScK:dev_user:DEVU-10",
              "state": "active",
              "thumbnail": "https://api.dev.devrev-eng.ai/internal/display-picture/Gowtham%20Gopinath.png"
            }
          ],
          "priority": "p2",
          "priority_v2": {
            "id": 3,
            "label": "P2",
            "ordinal": 3
          },
          "stage": {
            "name": "prioritized",
            "stage": {
              "display_name": "Prioritized",
              "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_stage/27",
              "name": "prioritized"
            }
          },
          "title": "issue with a subtype"
        }
      ],
      "stage": {
        "display_name": "Prioritized",
        "name": "prioritized",
        "notes": "",
        "ordinal": 3600,
        "stage": {
          "display_name": "Prioritized",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_stage/27",
          "name": "prioritized"
        },
        "state": {
          "display_name": "Open",
          "id": "don:core:dvrv-us-1:devo/11FVC3ScK:custom_state/1",
          "is_final": false,
          "name": "open"
        }
      },
      "state": "open",
      "stock_schema_fragment": "don:core:dvrv-us-1:stock_sf/5774789",
      "subtype": "test",
      "title": "issue with a subtype"
    }
  },
  "type": "work_updated"
}
```
