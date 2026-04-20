import json

with open("postman/postman_collection.json", "r") as f:
    data = json.load(f)

new_folder = {
    "name": "ApplicationResult Management",
    "item": [
        {
            "name": "Get Admin Application Result",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/applications/1/result",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "applications", "1", "result"]
                }
            }
        },
        {
            "name": "Update Application Result (Manual)",
            "request": {
                "method": "PUT",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"priority_score\": 1.0,\n    \"additional_score\": 0.5,\n    \"result_status\": \"Passed\",\n    \"notes\": \"Hồ sơ hoàn hảo\"\n}"
                },
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/applications/1/result",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "applications", "1", "result"]
                }
            }
        },
        {
            "name": "Recalculate Application Result",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/applications/1/result/recalculate",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "applications", "1", "result", "recalculate"]
                }
            }
        },
        {
            "name": "Recalculate Ranking By Admission Period",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{admin_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/admin/admission-periods/1/results/recalculate-ranking",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "admin", "admission-periods", "1", "results", "recalculate-ranking"]
                }
            }
        },
        {
            "name": "Get User Application Result",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{user_access_token}}"
                    }
                ],
                "url": {
                    "raw": "{{base_url}}/api/v1/user/me/applications/1/result",
                    "host": ["{{base_url}}"],
                    "path": ["api", "v1", "user", "me", "applications", "1", "result"]
                }
            }
        }
    ]
}

exists = False
for item in data['item']:
    if item['name'] == 'ApplicationResult Management':
        item['item'] = new_folder['item']
        exists = True
        break

if not exists:
    data['item'].append(new_folder)

with open("postman/postman_collection.json", "w") as f:
    json.dump(data, f, indent=4, ensure_ascii=False)
